package impl

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"
)

var (
	binPostfix = ""
)

func BuildFromSource(dstBinPath string) {
	println("======== Building e from source =======")
	// Install SCons and Ninja
	installPythonPackages()

	isNeedDelete := true
	// Check if the "godot" directory exists
	if _, err := os.Stat("godot"); os.IsNotExist(err) {
		fmt.Println("Godot directory not found. Creating and initializing...")
		if err := os.Mkdir("godot", 0755); err != nil {
			log.Fatalf("Failed to create godot directory: %v", err)
		}

		if err := os.Chdir("godot"); err != nil {
			log.Fatalf("Failed to change to godot directory: %v", err)
		}

		runCommand("git", "init")
		runCommand("git", "remote", "add", "origin", "https://github.com/realdream-ai/godot.git")
		runCommand("git", "fetch", "--depth", "1", "origin", "spx"+version)
		runCommand("git", "checkout", "spx"+version)

		fmt.Println("Godot repository setup complete.")
	} else {
		if err := os.Chdir("godot"); err != nil {
			log.Fatalf("Failed to change to godot directory: %v", err)
		}

		fmt.Println("Godot directory already exists.")
		isNeedDelete = false
	}

	// Check the operating system
	if runtime.GOOS == "windows" {
		runCommand("scons", "target=editor", "arch=x86_64", "vsproj=yes", "dev_build=yes")
	} else {
		runCommand("scons", "target=editor", "arch=x86_64", "dev_build=yes")
	}

	// Change back to the parent directory
	if err := os.Chdir(".."); err != nil {
		log.Fatalf("Failed to change to parent directory: %v", err)
	}

	fmt.Println("Build engine done")
	// Copy the binary
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = getGoEnv()
	}

	fmt.Printf("Destination binary path: %s\n", dstBinPath)
	filePath := findFirstMatchingFile("godot/bin", "godot.*.editor.dev.*", "console")
	if filePath == "" {
		fmt.Println("No matching file found. " + filePath)
		if isNeedDelete {
			os.RemoveAll("./godot")
		}
	}

	if err := CopyFile(filePath, dstBinPath); err != nil {
		log.Fatalf("Failed to copy binary: %v", err)
	}
	if isNeedDelete {
		os.RemoveAll("./godot")
	}
}

func findFirstMatchingFile(dir, pattern, exclude string) string {
	var foundFile string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if match, _ := filepath.Match(pattern, info.Name()); match {
			if !strings.Contains(info.Name(), exclude) {
				foundFile = path
				return filepath.SkipDir
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return foundFile
}

// Helper function to run a command
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command %s failed: %v", name, err)
	}
}

// Install SCons and Ninja using pip
func installPythonPackages() {
	runCommand("pip", "install", "scons==4.7.0")
}

func CheckPresetEnvironment() {
	if !isPythonInstalled("python3") && !isPythonInstalled("python") {
		fmt.Println("Python is not installed. Please install Python first, python version should >= 3.8")
		os.Exit(1)
	}
}

// Helper function to check if a specific Python command exists
func isPythonInstalled(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// Get the Go environment variable
func getGoEnv() string {
	cmd := exec.Command("go", "env", "GOPATH")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to get GOPATH: %v", err)
	}
	return string(output)
}

// Copy a file
func CopyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, input, 0755); err != nil {
		return err
	}
	return nil
}

func downloadGd4spx(fileName string, dstPath string) error {
	url := "https://118.195.190.67/bin/" + fileName
	dest := dstPath
	fmt.Println("Downloading file... ", url, "=>", dest)
	err := downloadFile(url, dest)
	return err
}

func checkAndGetBinPath() (string, error) {
	if runtime.GOOS == "windows" {
		binPostfix = ".exe"
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	dstBinPath := gopath + "/bin/gd4spx" + version + binPostfix
	gd4spx, err := exec.LookPath("gd4spx")
	if err == nil {
		if current, err := exec.Command(gd4spx, "--version").CombinedOutput(); err == nil {
			if strings.HasPrefix(string(current), version+".") {
				return gd4spx, nil
			}
		}
	}
	info, err := os.Stat(dstBinPath)
	if os.IsNotExist(err) {
		BuildFromSource(dstBinPath)
	} else if err != nil {
		return "", err
	} else {
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(dstBinPath, 0755); err != nil {
				return "", err
			}
		}
	}
	return dstBinPath, nil
}
