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
	fmt.Println((cmdName + " Version = " + Version + "\n") + strings.ReplaceAll(msg, "#CMDNAME", cmdName))
}

func PrepareEnv(fsRelDir, dstDir string) {
	util.CopyDir(proejctFS, fsRelDir, dstDir, false)
	rawDir, _ := os.Getwd()
	os.Chdir(GoDir)
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
	binPostfix, CmdPath, err = impl.CheckAndGetAppPath(appName, Version)
	if err != nil {
		return fmt.Errorf(appName+"requires engine to be installed as a binary at $GOPATH/bin/: %w", err)
	}
	ProjectDir, _ = filepath.Abs(path.Join(TargetDir, projectRelPath))
	GoDir, _ = filepath.Abs(ProjectDir + "/go")

	var libraryName = fmt.Sprintf(appName+"-%v-%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryName += ".dll"
	case "darwin":
		libraryName += ".dylib"
	default:
		libraryName += ".so"
	}
	LibPath, _ = filepath.Abs(path.Join(ProjectDir, "lib", libraryName))
	return nil
}

func ImportProj() error {
	curCmd.BuildDll()
	fmt.Println(" ================= Importing ... ================= ")
	cmd := exec.Command(CmdPath, "--import", "--headless")
	cmd.Dir = ProjectDir
	cmd.Start()
	cmd.Wait()
	return nil
}

func ExportWebEditor() error {
	gopath := build.Default.GOPATH
	editorZipPath := path.Join(gopath, "bin", appName+Version+"_web.zip")
	dstPath := path.Join(ProjectDir, ".builds/web")
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
	if !util.IsFileExist(path.Join(ProjectDir, ".builds/web")) {
		return ExportWeb()
	}
	return nil
}

func Run(args string) error {
	return util.RunCommandInDir(ProjectDir, CmdPath, args)
}

func ExportWeb() error {
	curCmd.BuildDll()
	// Delete gdextension
	os.RemoveAll(filepath.Join(ProjectDir, "lib"))
	os.Remove(filepath.Join(ProjectDir, ".godot", "extension_list.cfg"))
	os.Remove(filepath.Join(ProjectDir, "gdx.gdextension"))
	curCmd.BuildWasm()
	err := ExportBuild("Web")
	return err
}

func Clear() {
	os.RemoveAll(filepath.Join(ProjectDir, "lib"))
	os.RemoveAll(filepath.Join(ProjectDir, ".godot"))
	os.RemoveAll(filepath.Join(ProjectDir, ".build"))
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
	port := ServerPort
	curCmd.StopWeb()
	scriptPath := filepath.Join(ProjectDir, ".godot", "gdspx_web_server.py")
	scriptPath = strings.ReplaceAll(scriptPath, "\\", "/")
	executeDir := filepath.Join(ProjectDir, ".builds/web")
	executeDir = strings.ReplaceAll(executeDir, "\\", "/")
	println("web server running at http://127.0.0.1:" + fmt.Sprint(port))
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
	return util.RunCommandInDir(ProjectDir, CmdPath, args)
}

func BuildDll() {
	rawdir, _ := os.Getwd()
	os.Chdir(GoDir)
	envVars := []string{"CGO_ENABLED=1"}
	util.RunGolang(envVars, "build", "-o", LibPath, "-buildmode=c-shared")
	os.Chdir(rawdir)
}

func BuildWasm() {
	rawdir, _ := os.Getwd()
	dir := path.Join(ProjectDir, ".builds/web/")
	os.MkdirAll(dir, 0755)
	filePath := path.Join(dir, "gdspx.wasm")
	os.Chdir(GoDir)
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	util.RunGolang(envVars, "build", "-o", filePath)
	os.Chdir(rawdir)
}

func ExportBuild(platform string) error {
	println("start export: platform =", platform, " ProjectDir =", ProjectDir)
	os.MkdirAll(filepath.Join(ProjectDir, ".builds", strings.ToLower(platform)), os.ModePerm)
	cmd := exec.Command(CmdPath, "--headless", "--quit", "--path", ProjectDir, "--export-debug", platform)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error exporting to web:", err)
	}
	return err
}
