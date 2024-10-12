package impl

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	_ "embed"
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
func ShowHelpInfo() {
	fmt.Println("gd4spx version", version)
	fmt.Println(`
Usage:

    gdspx <command> [path]      

The commands are:

    - init            # Create a gdspx project in the current directory
    - run             # Run the current project
    - editor          # Open the current project in editor mode
    - build           # Build the dynamic library
    - export          # Export the PC package (macOS, Windows, Linux) (TODO)
    - buildweb        # Build for WebAssembly (WASM)
    - exportweb       # Export the web package

 eg:

    gdspx init                      # create a project in current path
    gdspx init ./test/demo01        # create a project at path ./test/demo01 
	`)
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
	dir := path.Join(project, "build/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdspx.wasm")
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	RunGolang(envVars, "build", "-o", filePath)
	os.Chdir(rawdir)
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
