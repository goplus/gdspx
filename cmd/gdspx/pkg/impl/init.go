package impl

import (
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
)

var (
	binPostfix = ""
)

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

func downloadPack(dstDir, tagName, postfix string) error {
	urlHeader := "https://github.com/realdream-ai/godot/releases/download/"
	fileName := tagName + postfix
	url := urlHeader + tagName + "/" + fileName
	// download pc
	err := downloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download web
	fileName = tagName + "_web.zip"
	url = urlHeader + tagName + "/" + fileName
	err = downloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download webpack
	fileName = tagName + "_webpack.zip"
	url = urlHeader + tagName + "/" + fileName
	err = downloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	return err
}

func checkAndGetBinPath() (string, error) {
	if runtime.GOOS == "windows" {
		binPostfix = "_win.exe"
	} else if runtime.GOOS == "darwin" {
		binPostfix = "_darwin"
	} else if runtime.GOOS == "linux" {
		binPostfix = "_linux"
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	version := Version
	dstDir := gopath + "/bin"
	tagName := "gdspx" + version
	dstFileName := tagName + binPostfix
	dstBinPath := path.Join(dstDir, dstFileName)
	gdspx, err := exec.LookPath(dstFileName)
	if err == nil {
		if _, err := exec.Command(gdspx, "--version").CombinedOutput(); err == nil {
			return gdspx, nil
		}
	}

	info, err := os.Stat(dstBinPath)
	if os.IsNotExist(err) {
		println("Downloading gdspx pack...")
		err := downloadPack(dstDir, tagName, binPostfix)
		if err != nil {
			print("downloadPack error:" + err.Error())
			return "", err
		}
		if err := os.Chmod(dstBinPath, 0755); err != nil {
			return "", err
		}
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
