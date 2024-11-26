package cmdtool

import (
	"embed"
	"fmt"
	"os"
)

func RunCmd(cmd ICmdTool, appName, version string, fs embed.FS, fsRelDir string, dstRelDir string, ext ...string) (err error) {
	cmd.Register()
	err = cmd.CheckCmd(ext...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	err = cmd.CheckEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	switch os.Args[1] {
	case "help", "version":
		cmd.ShowHelpInfo()
		return
	case "clear":
		cmd.Clear()
		return
	case "stopweb":
		cmd.StopWeb()
		return
	case "init":
		return
	}

	err = cmd.OnBeforeCheck(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	targetDirArg := "."
	if len(os.Args) > 2 {
		targetDirArg = os.Args[2]
	}
	err = cmd.SetupEnv(appName, version, fs, fsRelDir, targetDirArg, dstRelDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	switch os.Args[1] {
	case "editor", "run", "export", "build":
		cmd.BuildDll()
	case "buildweb", "runweb", "exportweb":
		cmd.BuildWasm()
	}

	switch os.Args[1] {
	case "editor":
		err = cmd.Run("-e")
	case "run":
		err = cmd.Run("")
	case "export":
		err = cmd.Export()
	case "runweb":
		err = cmd.RunWeb()
	case "exportweb":
		err = cmd.ExportWeb()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	err = cmd.OnAfterCheck(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	return
}
