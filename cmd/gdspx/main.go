package main

import (
	"fmt"
	"godot-ext/gdspx/cmd/gdspx/pkg/impl"
	"os"
)

func main() {
	impl.CheckPresetEnvironment()
	impl.TargetDir = "."
	if len(os.Args) > 2 {
		impl.TargetDir = os.Args[2]
	}
	if len(os.Args) <= 1 {
		impl.ShowHelpInfo()
		return
	}

	switch os.Args[1] {
	case "help", "version":
		impl.ShowHelpInfo()
		return
	case "init":
		impl.PrepareGoEnv("", "")
	}
	if err := wrap(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func wrap() error {
	// look for a go.mod file
	gd4spxPath, project, libPath, err := impl.SetupEnv()
	if err != nil {
		return err
	}

	switch os.Args[1] {
	case "init":
		return nil
	case "run", "editor", "export", "build":
		impl.BuildDll(project, libPath)
	case "buildweb", "exportweb":
		impl.BuildWasm(project)
	}

	switch os.Args[1] {
	case "run":
		return impl.RunGdspx(gd4spxPath, project, "")
	case "editor":
		return impl.RunGdspx(gd4spxPath, project, "-e")
	}
	return nil
}
