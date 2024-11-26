package cmdtool

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/impl"
	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/util"
)

func CheckCmd(ext ...string) bool {
	cmds := []string{
		"help", "version", "editor",
		"init", "clear", "clearbuild",
		"build", "run", "export",
		"runweb", "buildweb", "exportweb", "stopweb",
	}
	cmds = append(cmds, ext...)

	cmdName := os.Args[1]
	for _, b := range cmds {
		if b == cmdName {
			return true
		}
	}
	return false
}

func ShowHelpInfo(cmdName, version string) {
	msg := `
Usage:

    #CMDNAME <command> [path]      

The commands are:

    - init            # Create a #CMDNAME project in the current directory
    - editor          # Open the current project in editor mode

    - build           # Build the dynamic library
    - run             # Run the current project
    - export          # Export the PC package (macOS, Windows, Linux) 

    - buildweb        # Build for WebAssembly (WASM)
    - runweb          # Launch the web server
    - exportweb       # Export the web package

    - clear           # Clear the project 

 eg:

    #CMDNAME init                      # create a project in current path
    #CMDNAME init ./test/demo01        # create a project at path ./test/demo01 
	`
	fmt.Println((cmdName + " version = " + version + "\n") + strings.ReplaceAll(msg, "#CMDNAME", cmdName))
}

func PrepareEnv(fsRelDir, dstDir string) {
	util.CopyDir(proejctFS, fsRelDir, dstDir, false)
	rawDir, _ := os.Getwd()
	os.Chdir(goDir)
	util.RunGolang(nil, "mod", "tidy")
	os.Chdir(rawDir)
}

func SetupEnv() error {
	var GOOS, GOARCH = runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" {
		return errors.New("gdx requires an amd64, or an arm64 system")
	}
	var err error
	binPostfix, cmdPath, err = impl.CheckAndGetAppPath(appName, version)
	if err != nil {
		return fmt.Errorf(appName+"requires engine to be installed as a binary at $GOPATH/bin/: %w", err)
	}
	projectDir, _ = filepath.Abs(targetDir + "/project")
	goDir, _ = filepath.Abs(projectDir + "/go")

	wd := goDir

	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return fmt.Errorf(appName + " requires your project to have a go.mod file")
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

	var libraryName = fmt.Sprintf(appName+"-%v-%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryName += ".dll"
	case "darwin":
		libraryName += ".dylib"
	default:
		libraryName += ".so"
	}
	libPath, _ = filepath.Abs(path.Join(projectDir, "lib", libraryName))

	ImportProj()
	return nil
}
func ImportProj() error {
	_, err := os.Stat(projectDir + "/.godot")
	hasInited := !os.IsNotExist(err)
	os.MkdirAll(projectDir+"/.godot", 0755)
	if !hasInited {
		curCmd.BuildDll()
	}

	cmd := exec.Command(cmdPath, "--import", "--headless")
	cmd.Dir = projectDir
	cmd.Start()
	cmd.Wait()
	fmt.Println("Import done")
	return nil
}

func ExportWebEditor() error {
	gopath := build.Default.GOPATH
	editorZipPath := path.Join(gopath, "bin", appName+version+"_web.zip")
	dstPath := path.Join(projectDir, ".builds/web")
	os.MkdirAll(dstPath, os.ModePerm)
	if util.IsFileExist(editorZipPath) {
		util.Unzip(editorZipPath, dstPath)
	} else {
		return errors.New("editor zip file not found: " + editorZipPath)
	}
	os.Rename(path.Join(dstPath, "godot.editor.html"), path.Join(dstPath, "index.html"))
	return nil
}
func CheckExportWeb() error {
	if !util.IsFileExist(path.Join(projectDir, ".builds/web")) {
		return ExportWeb()
	}
	return nil
}

func Run(args string) error {
	return util.RunCommandInDir(projectDir, cmdPath, args)
}

func ExportWeb() error {
	curCmd.BuildDll()
	// Delete gdextension
	os.RemoveAll(filepath.Join(projectDir, "lib"))
	os.Remove(filepath.Join(projectDir, ".godot", "extension_list.cfg"))
	os.Remove(filepath.Join(projectDir, "gdx.gdextension"))
	curCmd.BuildWasm()
	err := ExportBuild("Web")
	return err
}

func Clear() {
	os.RemoveAll(filepath.Join(projectDir, "lib"))
	os.RemoveAll(filepath.Join(projectDir, ".godot"))
	os.RemoveAll(filepath.Join(projectDir, ".build"))
}

func StopWeb() {
	if runtime.GOOS == "windows" {
		content := "taskkill /F /IM python.exe\r\ntaskkill /F /IM pythonw.exe\r\n"
		tempFileName := "temp_kill.bat"
		os.WriteFile(tempFileName, []byte(content), 0644)
		cmd := exec.Command("cmd.exe", "/C", tempFileName)
		cmd.Run()
		os.Remove(tempFileName)
	} else {
		cmd := exec.Command("pkill", "-f", "gdx_web_server.py")
		cmd.Run()
	}
}
func RunWeb() error {
	port := serverPort
	curCmd.StopWeb()
	scriptPath := filepath.Join(projectDir, ".godot", "gdx_web_server.py")
	executeDir := filepath.Join(projectDir, "../", ".builds/web")
	println("web server running at http://localhost:" + fmt.Sprint(port))
	cmd := exec.Command("python", scriptPath, "-r", executeDir, "-p", fmt.Sprint(port))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting server: %v", err)
	}
	return nil
}

func Export() error {
	platform := "Win"
	args := "--headless --quit --export-debug " + platform
	return util.RunCommandInDir(projectDir, cmdPath, args)
}

func BuildDll() {
	rawdir, _ := os.Getwd()
	os.Chdir(goDir)
	envVars := []string{"CGO_ENABLED=1"}
	util.RunGolang(envVars, "build", "-o", libPath, "-buildmode=c-shared")
	os.Chdir(rawdir)
}

func BuildWasm() {
	rawdir, _ := os.Getwd()
	dir := path.Join(projectDir, ".builds/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdspx.wasm")
	os.Chdir(goDir)
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	util.RunGolang(envVars, "build", "-o", filePath)
	os.Chdir(rawdir)
}

func ExportBuild(platform string) error {
	println("start export: platform =", platform, " projectDir =", projectDir)
	os.MkdirAll(filepath.Join(projectDir, ".builds", strings.ToLower(platform)), os.ModePerm)
	cmd := exec.Command(cmdPath, "--headless", "--quit", "--path", projectDir, "--export-debug", platform)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error exporting to web:", err)
	}
	return err
}
