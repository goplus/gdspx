package cmdtool

import (
	"errors"
	"fmt"
	"go/build"
	"log"
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
		"exportapk", "exportios",
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
func ExportIos() error {
	// First build the iOS libraries
	if err := buildIosLibraries(); err != nil {
		return fmt.Errorf("failed to build iOS libraries: %w", err)
	}

	// Set up paths
	ipaPath := filepath.Join(ProjectDir, ".builds", "ios", "Game.ipa")
	buildDir := filepath.Dir(ipaPath)

	// Create builds directory if it doesn't exist
	if err := os.MkdirAll(buildDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create build directory: %w", err)
	}

	// Check if Godot binary exists
	if _, err := os.Stat(CmdPath); os.IsNotExist(err) {
		return fmt.Errorf("Godot binary not found at %s", CmdPath)
	}

	// Check if project file exists
	projectFilePath := filepath.Join(ProjectDir, "project.godot")
	if _, err := os.Stat(projectFilePath); os.IsNotExist(err) {
		return fmt.Errorf("Godot project file not found at %s", projectFilePath)
	}

	// Import project to ensure resources are up to date
	fmt.Println("Importing project resources...")
	cmd := exec.Command(CmdPath, "--headless", "--path", ProjectDir, "--editor", "--quit")
	if err := cmd.Run(); err != nil {
		fmt.Println("Warning: project import may have issues:", err)
	}

	// Export the project to IPA
	fmt.Println("Exporting Godot project to IPA...")
	cmd = exec.Command(CmdPath, "--headless", "--path", ProjectDir, "--export-debug", "iOS", ipaPath)

	// Capture standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("IPA export failed: %w", err)
	}

	// Check if IPA was created
	if _, err := os.Stat(ipaPath); os.IsNotExist(err) {
		return fmt.Errorf("IPA export failed: file not created at %s", ipaPath)
	}

	log.Println("IPA export completed successfully!", ipaPath)

	log.Println("Try to install ipa to devices...")
	// install ipa to device
	cmd = exec.Command("ios-deploy", "--bundle", ipaPath)

	// Capture standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("IPA install failed: %w", err)
	}
	return nil
}

func buildIosLibraries() error {
	// Configuration variables
	frameworkName := "gdspx"
	libDir := filepath.Join(ProjectDir, "lib")
	xcframeworkPath := filepath.Join(libDir, "lib"+frameworkName+".ios.xcframework")
	buildDir := filepath.Join(ProjectDir, ".godot", "tmp", "gobuild")
	simulatorDir := filepath.Join(buildDir, "simulator")
	deviceDir := filepath.Join(buildDir, "device")
	headersDir := filepath.Join(buildDir, "headers")
	goSrcDir := filepath.Join(ProjectDir, "go")

	// Create directories
	os.RemoveAll(buildDir)
	os.RemoveAll(xcframeworkPath)
	for _, dir := range []string{simulatorDir, deviceDir, libDir, headersDir} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	fmt.Println("📦 Building Go libraries for iOS...")

	// Create a dummy header file with the required exports
	headerContent := `#ifndef LIBGDSPX_H
#define LIBGDSPX_H

#include <stdlib.h>

// GDExtension initialization function
void GDExtensionInit(void *p_interface, const void *p_library, void *r_initialization);

#endif // LIBGDSPX_H
`
	if err := os.WriteFile(filepath.Join(headersDir, "libgdspx.h"), []byte(headerContent), 0644); err != nil {
		return fmt.Errorf("failed to create header file: %w", err)
	}

	// Copy C headers to the headers directory
	headerFiles, err := filepath.Glob(filepath.Join(goSrcDir, "*.h"))
	if err != nil {
		return fmt.Errorf("failed to find header files: %w", err)
	}
	for _, headerFile := range headerFiles {
		destFile := filepath.Join(headersDir, filepath.Base(headerFile))
		if err := util.CopyFile(headerFile, destFile); err != nil {
			return fmt.Errorf("failed to copy header file %s: %w", headerFile, err)
		}
	}

	// Get SDK paths
	simulatorSdkPath, err := exec.Command("xcrun", "--sdk", "iphonesimulator", "--show-sdk-path").Output()
	if err != nil {
		return fmt.Errorf("failed to get simulator SDK path: %w", err)
	}
	deviceSdkPath, err := exec.Command("xcrun", "--sdk", "iphoneos", "--show-sdk-path").Output()
	if err != nil {
		return fmt.Errorf("failed to get device SDK path: %w", err)
	}

	// Disable signal handling in Go for iOS
	os.Setenv("GODEBUG", "cgocheck=0,asyncpreemptoff=1,panicnil=1")

	// Build for iOS Simulator (x86_64)
	fmt.Println("🔨 Building for iOS Simulator (x86_64)...")
	cmd := exec.Command("go", "build", "-tags=ios", "-buildmode=c-archive", "-trimpath", "-ldflags=-w -s", "-o", filepath.Join(simulatorDir, "libgdspx-x86_64.a"), ".")
	cmd.Dir = goSrcDir
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=1",
		"GOOS=darwin",
		"GOARCH=amd64",
		"CGO_CFLAGS=-isysroot "+strings.TrimSpace(string(simulatorSdkPath))+" -mios-simulator-version-min=12.0 -arch x86_64 -fembed-bitcode",
		"CGO_LDFLAGS=-isysroot "+strings.TrimSpace(string(simulatorSdkPath))+" -mios-simulator-version-min=12.0 -arch x86_64",
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for iOS Simulator (x86_64): %w", err)
	}

	// Build for iOS Simulator (arm64)
	fmt.Println("🔨 Building for iOS Simulator (arm64)...")
	cmd = exec.Command("go", "build", "-tags=ios", "-buildmode=c-archive", "-trimpath", "-ldflags=-w -s", "-o", filepath.Join(simulatorDir, "libgdspx-arm64-sim.a"), ".")
	cmd.Dir = goSrcDir
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=1",
		"GOOS=darwin",
		"GOARCH=arm64",
		"CGO_CFLAGS=-isysroot "+strings.TrimSpace(string(simulatorSdkPath))+" -mios-simulator-version-min=12.0 -arch arm64 -fembed-bitcode",
		"CGO_LDFLAGS=-isysroot "+strings.TrimSpace(string(simulatorSdkPath))+" -mios-simulator-version-min=12.0 -arch arm64",
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for iOS Simulator (arm64): %w", err)
	}

	// Build for iOS Device (arm64)
	fmt.Println("🔨 Building for iOS Device (arm64)...")
	cmd = exec.Command("go", "build", "-tags=ios", "-buildmode=c-archive", "-trimpath", "-ldflags=-w -s", "-o", filepath.Join(deviceDir, "libgdspx-arm64.a"), ".")
	cmd.Dir = goSrcDir
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=1",
		"GOOS=darwin",
		"GOARCH=arm64",
		"CGO_CFLAGS=-isysroot "+strings.TrimSpace(string(deviceSdkPath))+" -mios-version-min=12.0 -arch arm64 -fembed-bitcode",
		"CGO_LDFLAGS=-isysroot "+strings.TrimSpace(string(deviceSdkPath))+" -mios-version-min=12.0 -arch arm64",
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for iOS Device (arm64): %w", err)
	}

	// Create a fat binary for simulator (combines arm64 and x86_64)
	fmt.Println("🔗 Creating fat binary for simulator...")
	cmd = exec.Command("lipo", "-create", "-output", filepath.Join(simulatorDir, "libgdspx.a"),
		filepath.Join(simulatorDir, "libgdspx-x86_64.a"),
		filepath.Join(simulatorDir, "libgdspx-arm64-sim.a"))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create fat binary for simulator: %w", err)
	}

	// Create XCFramework
	fmt.Println("🎁 Creating XCFramework...")
	cmd = exec.Command("xcrun", "xcodebuild", "-create-xcframework",
		"-library", filepath.Join(simulatorDir, "libgdspx.a"), "-headers", headersDir,
		"-library", filepath.Join(deviceDir, "libgdspx-arm64.a"), "-headers", headersDir,
		"-output", xcframeworkPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create XCFramework: %w", err)
	}

	// Clean up temporary build files
	fmt.Println("🧹 Cleaning up temporary build files...")
	os.RemoveAll(buildDir)

	fmt.Println("✅ Successfully built libgdspx.ios.xcframework!")
	fmt.Println("📍 Location:", xcframeworkPath)

	return nil
}

func ExportApk() error {
	// First build the dynamic libraries for Android
	if err := buildAndroidLibraries(); err != nil {
		return fmt.Errorf("failed to build Android libraries: %w", err)
	}

	// Set up paths
	apkPath := filepath.Join(ProjectDir, ".builds", "android", "game.apk")
	buildDir := filepath.Dir(apkPath)

	// Create builds directory if it doesn't exist
	if err := os.MkdirAll(buildDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create build directory: %w", err)
	}

	// Check if Godot binary exists
	if _, err := os.Stat(CmdPath); os.IsNotExist(err) {
		return fmt.Errorf("Godot binary not found at %s", CmdPath)
	}

	// Check if project file exists
	projectFilePath := filepath.Join(ProjectDir, "project.godot")
	if _, err := os.Stat(projectFilePath); os.IsNotExist(err) {
		return fmt.Errorf("Godot project file not found at %s", projectFilePath)
	}

	BuildDll()
	// Import project to ensure resources are up to date
	fmt.Println("Importing project resources...")
	cmd := exec.Command(CmdPath, "--headless", "--path", ProjectDir, "--editor", "--quit")
	if err := cmd.Run(); err != nil {
	}

	// Export the project to APK
	fmt.Println("Exporting Godot project to APK...")
	cmd = exec.Command(CmdPath, "--headless", "--path", ProjectDir, "--export-debug", "Android", apkPath)

	// Capture standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("APK export failed: %w", err)
	}

	// Check if APK was created
	if _, err := os.Stat(apkPath); os.IsNotExist(err) {
		fmt.Println("APK export failed: file not created at ", apkPath)
		return nil
	}
	log.Println("APK export completed successfully!", apkPath)

	// Check if adb is available
	_, err := exec.LookPath("adb")
	if err != nil {
		fmt.Println("adb command not found. Please ensure Android SDK platform tools are installed and in your PATH")
		return nil
	}

	// Check if any Android device is connected
	cmd = exec.Command("adb", "devices")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("failed to check for connected devices:", err)
		return nil
	}

	if !strings.Contains(string(output), "device\n") {
		fmt.Println("no Android device connected. Please connect a device and enable USB debugging")
		return nil
	}

	// Install the APK
	fmt.Println("Installing APK...")
	cmd = exec.Command("adb", "install", "-r", apkPath)
	if err := cmd.Run(); err != nil {
		fmt.Println("APK installation failed:", err)
		return nil
	}

	fmt.Println("APK installation successful!")
	return nil
}

func buildAndroidLibraries() error {
	libDir := filepath.Join(ProjectDir, "lib")
	goDir := filepath.Join(ProjectDir, "go")

	// Check if ANDROID_NDK_ROOT is set
	androidNdkRoot := os.Getenv("ANDROID_NDK_ROOT")
	if androidNdkRoot == "" {
		fmt.Println("ANDROID_NDK_ROOT environment variable is not set")
		return nil
	}

	// Detect system architecture and OS
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Set host tag based on OS and architecture
	hostTag := ""
	switch osName {
	case "windows":
		hostTag = "windows-x86_64"
	case "linux":
		if arch == "amd64" {
			hostTag = "linux-x86_64"
		} else if arch == "arm64" {
			hostTag = "linux-aarch64"
		} else {
			return fmt.Errorf("unsupported Linux architecture: %s", arch)
		}
	case "darwin":
		hostTag = "darwin-x86_64"
	default:
		return fmt.Errorf("unsupported operating system: %s", osName)
	}

	// Create lib directory if it doesn't exist
	if err := os.MkdirAll(libDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create lib directory: %w", err)
	}

	// Set NDK toolchain path and minimum SDK version
	ndkToolchain := filepath.Join(androidNdkRoot, "toolchains", "llvm", "prebuilt", hostTag, "bin")
	minSdk := "21"

	// Build for arm64-v8a
	fmt.Println("Building for arm64-v8a...")
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", filepath.Join(libDir, "libgdspx-android-arm64.so"), ".")
	cmd.Dir = goDir
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=arm64",
		"CC="+filepath.Join(ndkToolchain, "aarch64-linux-android"+minSdk+"-clang"),
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for arm64-v8a: %w", err)
	}

	// Build for armeabi-v7a
	fmt.Println("Building for armeabi-v7a...")
	cmd = exec.Command("go", "build", "-buildmode=c-shared", "-o", filepath.Join(libDir, "libgdspx-android-arm32.so"), ".")
	cmd.Dir = goDir
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=arm",
		"CC="+filepath.Join(ndkToolchain, "armv7a-linux-androideabi"+minSdk+"-clang"),
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build for armeabi-v7a: %w", err)
	}

	fmt.Println("Build android so completed successfully!")
	return nil
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
