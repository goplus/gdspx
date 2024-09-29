package main

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
	//go:embed template/project/project.godot
	project_gd4spx string

	//go:embed template/project/gdspx.gdextension
	library_gdextension string

	//go:embed template/project/main.tscn
	main_tscn string

	//go:embed template/project/export_presets.cfg
	export_presets_cfg string

	//go:embed template/project/extension_list.cfg
	extension_list_cfg string

	//go:embed template/go.mod.txt
	go_mode_txt string

	//go:embed template/main.go
	main_go string

	targetDir string
)

func printHelp() {
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
func main() {
	checkPythonInstalled()
	targetDir = "."
	if len(os.Args) > 2 {
		targetDir = os.Args[2]
	}
	if len(os.Args) <= 1 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "help", "version":
		printHelp()
		return
	case "init":
		if err := setupFile(false, targetDir+"/go.mod", go_mode_txt); err != nil {
			panic(err)
		}
		if err := setupFile(false, targetDir+"/main.go", main_go); err != nil {
			panic(err)
		}
		rawDir, err := os.Getwd()
		os.Chdir(targetDir)
		err = runGolang(nil, "mod", "tidy")
		os.Chdir(rawDir)
		if err != nil {
			println("gdspx project create failed ", targetDir)
			panic(err)
		} else {
			println("gdspx project create succ ", targetDir)
		}
	}
	if err := wrap(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func wrap() error {
	var GOOS, GOARCH = runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" {
		return errors.New("gd4spx requires an amd64, or an arm64 system")
	}
	gd4spxPath, err := checkAndGetBinPath()

	if err != nil {
		return fmt.Errorf("gd4spx requires Godot v%s to be installed as a binary at $GOPATH/bin/gd4spx-%s: %w", version, err)
	}
	wd := targetDir
	// look for a go.mod file
	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return fmt.Errorf("gd4spx requires your project to have a go.mod file")
		}
		_, err := os.Stat(wd + "/go.mod")
		if err == nil {
			break
		} else if os.IsNotExist(err) {
			continue
		} else {
			return err
		}
	}

	curDir := targetDir
	project := path.Join(curDir, "project")
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
	libPath := path.Join(project, "lib", libraryName)
	if err := setup(gd4spxPath, wd, project, libPath); err != nil {
		return err
	}

	switch os.Args[1] {
	case "init":
		return nil
	case "run", "editor", "export", "build":
		buildDll(project, libPath)
	case "buildweb", "exportweb":
		buildWasm(project)
	}

	switch os.Args[1] {
	case "run":
		return runGd4spx(gd4spxPath, project, "")
	case "editor":
		return runGd4spx(gd4spxPath, project, "-e")
	}

	return nil
}

func buildDll(projectPath, outputPath string) {
	rawdir, _ := os.Getwd()
	os.Chdir(path.Join(projectPath, "../"))
	runGolang(nil, "build", "-tags", "platform_pc", "-o", outputPath, "-buildmode=c-shared")
	os.Chdir(rawdir)
}

func buildWasm(project string) {
	dir := path.Join(project, "../build/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdspx.wasm")
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	runGolang(envVars, "build", "-tags", "platform_web", "-o", filePath)
}

func setup(gd4spxPath string, wd, project, libPath string) error {
	_, err := os.Stat(project + "/.godot")
	hasInited := !os.IsNotExist(err)
	if err := os.MkdirAll(project+"/.godot", 0755); err != nil {
		return err
	}
	if err := setupFile(false, project+"/main.tscn", main_tscn); err != nil {
		return err
	}
	if err := setupFile(false, project+"/project.godot", project_gd4spx); err != nil {
		return err
	}
	if err := setupFile(true, project+"/gdspx.gdextension", library_gdextension); err != nil {
		return err
	}
	if err := setupFile(false, project+"/.godot/extension_list.cfg", extension_list_cfg); err != nil {
		return err
	}
	if !hasInited {
		buildDll(project, libPath)
		gd4spx := exec.Command(gd4spxPath, "--import", "--headless")
		gd4spx.Dir = project
		gd4spx.Stderr = os.Stderr
		gd4spx.Stdout = os.Stdout
		gd4spx.Stdin = os.Stdin
		return gd4spx.Run()
	}
	return nil
}

func runGolang(envVars []string, args ...string) error {
	golang := exec.Command("go", args...)

	if envVars != nil {
		golang.Env = append(os.Environ(), envVars...)
	}
	golang.Stderr = os.Stderr
	golang.Stdout = os.Stdout
	golang.Stdin = os.Stdin
	return golang.Run()
}
func runGd4spx(gd4spxPath string, project string, args string) error {
	gd4spx := exec.Command(gd4spxPath, args)
	gd4spx.Dir = project
	gd4spx.Stderr = os.Stderr
	gd4spx.Stdout = os.Stdout
	gd4spx.Stdin = os.Stdin
	return gd4spx.Run()
}

func setupFile(force bool, name, embed string, args ...any) error {
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
