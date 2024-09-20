package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"
)

const version = "4.2.2"

var (
	//go:embed template/project/project.gd4spx
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
)

func main() {
	switch os.Args[1] {
	case "gen":
		generateCode()
		return
	case "init":
		if err := setupFile(false, "go.mod", go_mode_txt); err != nil {
			panic(err)
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
	gd4spxPath, err := useGd4spx()
	if err != nil {
		return fmt.Errorf("gd4spx requires Godot v%s to be installed as a binary at $GOPATH/bin/gd4spx-%s: %w", version, err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
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

	project := "./project"
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

	if err := setup(gd4spxPath, wd, project); err != nil {
		return err
	}

	switch os.Args[1] {
	case "init":
		return nil
	case "run", "editor", "export", "build":
		buildDll(project + "/lib/" + libraryName)
	case "run_web", "build_web", "export_web":
		buildWasm(project + "/lib/" + libraryName)
	}
	switch os.Args[1] {
	case "run":
		return runGd4spx(gd4spxPath, project, "")
	case "editor":
		return runGd4spx(gd4spxPath, project, "-e")
	}
	return nil
}

func buildDll(path string) {
	runGolang(nil, "build", "-buildmode=c-shared", "-o", path)
}

func buildWasm(path string) {
	envVars := []string{"GOOS=js", "GOARCH=wasm"}
	runGolang(envVars, "build", "-tags", "platform_web", "-o", path)
}

func setup(gd4spxPath string, wd, project string) error {
	if err := os.MkdirAll(project+"/.godot", 0755); err != nil {
		return err
	}
	if err := setupFile(false, project+"/main.tscn", main_tscn); err != nil {
		return err
	}
	if err := setupFile(false, project+"/project.gd4spx", project_gd4spx, filepath.Base(wd)); err != nil {
		return err
	}
	if err := setupFile(true, project+"/library.gdextension", library_gdextension); err != nil {
		return err
	}
	if err := setupFile(false, project+"/.godot/extension_list.cfg", extension_list_cfg); err != nil {
		return err
	}
	_, err := os.Stat(project + "/.godot")
	if os.IsNotExist(err) {
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

	if envVars == nil {
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

func downloadGd4spx(fileName string, dstPath string) error {
	url := "https://118.195.190.67/bin/" + fileName
	dest := dstPath
	fmt.Println("Downloading file... ", url, "=>", dest)
	err := downloadFile(url, dest)
	return err
}

func useGd4spx() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	dstBinPath := gopath + "/bin/gd4spx" + version
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
		switch runtime.GOOS {
		case "android":
			return "echo", nil
		case "darwin":
			downloadGd4spx("gd4spx"+version+".darwin.x86_64", dstBinPath)
		case "linux":
			fmt.Println("gd: downloading Godot v" + version + " stable for linux")
			downloadGd4spx("gd4spx"+version+".linux.x86_64", dstBinPath)
		default:
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(gopath+"/bin/gd4spx"+version, 0755); err != nil {
				return "", err
			}
		}
	}
	return dstBinPath, nil
}
