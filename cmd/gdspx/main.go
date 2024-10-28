package main

import (
	"fmt"
	"os"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/impl"
)

func main() {
	impl.CheckPresetEnvironment()
	impl.TargetDir = "."
	if len(os.Args) > 2 {
		impl.TargetDir = os.Args[2]
	}
	if len(os.Args) <= 1 {
		showHelpInfo()
		return
	}

	switch os.Args[1] {
	case "help", "version":
		showHelpInfo()
		return
	case "clear":
		impl.ClearGdspx(impl.TargetDir)
		return
	case "stopweb":
		impl.StopWebServer()
		return
	case "init":
		impl.PrepareGoEnv()
	}
	if err := execCmds(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func execCmds() error {
	return impl.ExecCmds(impl.BuildDll)
}

func showHelpInfo() {
	impl.ShowHelpInfo("gdspx")
}
