package impl

import (
	"crypto/tls"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const version = "4.2.2"

var (
	//go:embed template/project.godot
	project_gd4spx string

	//go:embed template/gdspx.gdextension
	library_gdextension string

	//go:embed template/main.tscn
	main_tscn string

	//go:embed template/export_presets.cfg
	export_presets_cfg string

	//go:embed template/gdspx_web_server.py
	gdspx_web_server_py string

	//go:embed template/extension_list.cfg
	extension_list_cfg string

	//go:embed template/go.mod.txt
	go_mode_txt string

	//go:embed template/main.go
	main_go string

	//go:embed template/gitignore.txt
	gitignore string

	TargetDir string
)

func ReplaceTemplate(mod string, main string, ignore string) {
	main_go = main
	go_mode_txt = mod
	gitignore = ignore
}

func PrepareGoEnv() {
	os.MkdirAll(TargetDir, 0755)
	if err := SetupFile(false, TargetDir+"/go.mod", go_mode_txt); err != nil {
		panic(err)
	}
	if err := SetupFile(false, TargetDir+"/main.go", main_go); err != nil {
		panic(err)
	}
	rawDir, err := os.Getwd()
	os.Chdir(TargetDir)
	err = RunGolang(nil, "mod", "tidy")
	os.Chdir(rawDir)
	if err != nil {
		println("gdspx project create failed ", TargetDir)
		panic(err)
	} else {
		println("gdspx project create succ ", TargetDir, "\n=====you can type  'gdspx run "+TargetDir+"'  to run the project======")
	}

}

func ShowHelpInfo(cmdName string) {
	msg := `
Usage:

    #CMDNAME <command> [path]      

The commands are:

    - init            # Create a #CMDNAME project in the current directory
    - run             # Run the current project
    - editor          # Open the current project in editor mode
    - build           # Build the dynamic library
    - export          # Export the PC package (macOS, Windows, Linux) (TODO)
    - runweb          # Launch the web server
    - buildweb        # Build for WebAssembly (WASM)
    - exportweb       # Export the web package
    - clear           # Clear the project 

 eg:

    #CMDNAME init                      # create a project in current path
    #CMDNAME init ./test/demo01        # create a project at path ./test/demo01 
	`
	fmt.Println(strings.ReplaceAll(msg, "#CMDNAME", cmdName))
}
func CopyEmbed(embedDir embed.FS, srcDir, dstDir string) error {
	if _, err := os.Stat(dstDir); !os.IsNotExist(err) {
		return nil
	}

	fsys, err := fs.Sub(embedDir, srcDir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dstDir, path)
		if d.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		} else {
			srcFile, err := fsys.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			return err
		}
	})
}

func BuildDll(project, outputPath string) {
	rawdir, _ := os.Getwd()
	os.Chdir(project)
	envVars := []string{"CGO_ENABLED=1"}
	RunGolang(envVars, "build", "-o", outputPath, "-buildmode=c-shared")
	os.Chdir(rawdir)
}

func BuildWasm(project string) {
	rawdir, _ := os.Getwd()
	os.Chdir(project)
	dir := path.Join(project, ".builds/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdspx.wasm")
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	RunGolang(envVars, "build", "-o", filePath)
	os.Chdir(rawdir)
}

func ExportBuild(gd4spxPath string, project string, platform string) error {
	println("start export: platform =", platform, " project =", project)
	os.MkdirAll(filepath.Join(project, ".builds", strings.ToLower(platform)), os.ModePerm)
	cmd := exec.Command(gd4spxPath, "--headless", "--quit", "--path", project, "--export-debug", platform)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error exporting to web:", err)
	}
	return err
}
func setup(gd4spxPath string, wd, project, libPath string) error {
	_, err := os.Stat(project + "/.godot")
	hasInited := !os.IsNotExist(err)
	if err := os.MkdirAll(project+"/.godot", 0755); err != nil {
		return err
	}
	if err := SetupFile(false, project+"/main.tscn", main_tscn); err != nil {
		return err
	}
	if err := SetupFile(false, project+"/project.godot", project_gd4spx); err != nil {
		return err
	}
	if err := SetupFile(false, project+"/.gitignore", gitignore); err != nil {
		return err
	}
	if err := SetupFile(true, project+"/gdspx.gdextension", library_gdextension); err != nil {
		return err
	}
	if err := SetupFile(false, project+"/.godot/extension_list.cfg", extension_list_cfg); err != nil {
		return err
	}

	if !hasInited {
		BuildDll(project, libPath)
		return ImportProj(project, libPath, gd4spxPath)
	}
	return nil
}

func ImportProj(project string, libPath string, gd4spxPath string) error {
	cmd := exec.Command(gd4spxPath, "--import", "--headless")
	cmd.Dir = project
	err := cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("ImportProj finished with error: %v\n", err)
	} else {
		fmt.Println("ImportProj successfully")
	}
	return nil
}

func RunGoplus(envVars []string, args ...string) error {
	golang := exec.Command("gop", args...)

	if envVars != nil {
		golang.Env = append(os.Environ(), envVars...)
	}
	golang.Stderr = os.Stderr
	golang.Stdout = os.Stdout
	golang.Stdin = os.Stdin
	return golang.Run()
}

func RunGolang(envVars []string, args ...string) error {
	golang := exec.Command("go", args...)

	if envVars != nil {
		golang.Env = append(os.Environ(), envVars...)
	}
	golang.Stderr = os.Stderr
	golang.Stdout = os.Stdout
	golang.Stdin = os.Stdin
	return golang.Run()
}
func RunGdspx(gd4spxPath string, project string, args string) error {
	SetupFile(false, filepath.Join(project, ".godot", "extension_list.cfg"), extension_list_cfg)
	SetupFile(false, filepath.Join(project, "gdspx.gdextension"), library_gdextension)

	println("run: ", gd4spxPath, project, args)
	gd4spx := exec.Command(gd4spxPath, args)
	gd4spx.Dir = project
	gd4spx.Stderr = os.Stderr
	gd4spx.Stdout = os.Stdout
	gd4spx.Stdin = os.Stdin
	return gd4spx.Run()
}
func StopWebServer() {
	if runtime.GOOS == "windows" {
		content := "taskkill /F /IM python.exe\r\ntaskkill /F /IM pythonw.exe\r\n"
		tempFileName := "temp_kill.bat"
		ioutil.WriteFile(tempFileName, []byte(content), 0644)
		cmd := exec.Command("cmd.exe", "/C", tempFileName)
		cmd.Run()
		os.Remove(tempFileName)
	} else {
		cmd := exec.Command("pkill", "-f", "gdspx_web_server.py")
		cmd.Run()
	}
}
func RunWebServer(gd4spxPath string, projPath string, libPath string, port int) error {
	if !IsFileExist(filepath.Join(projPath, ".builds", "web")) {
		ExportWeb(gd4spxPath, projPath, libPath)
	}
	if port == 0 {
		port = 8005
	}
	StopWebServer()
	scriptPath := filepath.Join(projPath, ".godot", "gdspx_web_server.py")
	executeDir := filepath.Join(projPath, "../", ".builds/web")
	SetupFile(false, scriptPath, gdspx_web_server_py)
	println("web server running at http://localhost:" + fmt.Sprint(port))
	cmd := exec.Command("python", scriptPath, "-r", executeDir, "-p", fmt.Sprint(port))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}
	return nil
}
func Export(gd4spxPath string, project string) error {
	platform := "Win"
	args := "--headless --quit --export-debug " + platform
	println("run: ", gd4spxPath, project, args)
	gd4spx := exec.Command(gd4spxPath, args)
	gd4spx.Dir = project
	gd4spx.Stderr = os.Stderr
	gd4spx.Stdout = os.Stdout
	gd4spx.Stdin = os.Stdin
	return gd4spx.Run()
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
func SetupFile(force bool, name, embed string, args ...any) error {
	if _, err := os.Stat(name); force || os.IsNotExist(err) {
		if len(args) > 0 {
			embed = fmt.Sprintf(embed, args...)
		}
		if err := os.WriteFile(name, []byte(embed), 0644); err != nil {
			return err
		}
	}
	return nil
}

func ExecCmds(buildDllFunc func(project, outputPath string)) error {
	gd4spxPath, project, libPath, err := SetupEnv()
	if err != nil {
		return err
	}
	switch os.Args[1] {
	case "init":
		return nil
	case "run", "editor", "export", "build":
		buildDllFunc(project, libPath)
	case "runweb", "buildweb":
		buildDllFunc(project, libPath)
		if !IsFileExist(path.Join(TargetDir, ".builds/web/index.html")) {
			ExportWeb(gd4spxPath, project, libPath)
		}
		BuildWasm(project)
	}
	switch os.Args[1] {
	case "run":
		return RunGdspx(gd4spxPath, project, "")
	case "editor":
		return RunGdspx(gd4spxPath, project, "-e")
	case "runweb":
		return RunWebServer(gd4spxPath, project, libPath, 8005)
	case "exportweb":
		return ExportWeb(gd4spxPath, project, libPath)
	case "export":
		return Export(gd4spxPath, project)
	}
	return nil
}
func downloadFile(url string, dest string) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func SetupEnv() (string, string, string, error) {
	var GOOS, GOARCH = runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" {
		return "", "", "", errors.New("gd4spx requires an amd64, or an arm64 system")
	}
	gd4spxPath, err := checkAndGetBinPath()
	if err != nil {
		return "", "", "", fmt.Errorf("gd4spx requires Godot v%s to be installed as a binary at $GOPATH/bin/gd4spx-%s: %w", version, version, err)
	}
	wd := TargetDir

	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return "", "", "", fmt.Errorf("gd4spx requires your project to have a go.mod file")
		}
		_, err := os.Stat(wd + "/go.mod")
		if err == nil {
			break
		} else if os.IsNotExist(err) {
			continue
		} else {
			return "", "", "", err
		}
	}

	curDir := TargetDir
	project := curDir
	if GOOS == "android" {
		project = "/sdcard/gd4spx/" + filepath.Base(wd)
	}

	var libraryName = fmt.Sprintf("gdspx-%v-%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryName += ".dll"
	case "darwin":
		libraryName += ".dylib"
	default:
		libraryName += ".so"
	}
	libPath := path.Join("lib", libraryName)
	if err := setup(gd4spxPath, wd, project, libPath); err != nil {
		return "", "", "", err
	}
	return gd4spxPath, project, libPath, nil
}

func ExportWeb(gd4spxPath string, projectPath string, libPath string) error {
	BuildDll(projectPath, libPath)
	// Delete gdextension
	os.RemoveAll(filepath.Join(projectPath, "lib"))
	os.Remove(filepath.Join(projectPath, ".godot", "extension_list.cfg"))
	os.Remove(filepath.Join(projectPath, "gdspx.gdextension"))
	// Copy template files
	SetupFile(false, filepath.Join(projectPath, "export_presets.cfg"), export_presets_cfg)

	BuildWasm(projectPath)
	err := ExportBuild(gd4spxPath, projectPath, "Web")
	return err
}
func ClearGdspx(projectPath string) {
	os.RemoveAll(filepath.Join(projectPath, "lib"))
	os.RemoveAll(filepath.Join(projectPath, ".godot"))
	os.RemoveAll(filepath.Join(projectPath, ".build"))
}
