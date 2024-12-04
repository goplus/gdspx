package cmdtool

import (
	"embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func RunCmd(cmd ICmdTool, appName, version string, fs embed.FS, fsRelDir string, dstRelDir string, ext ...string) (err error) {
	cmd.Register()
	err = cmd.CheckCmd(ext...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	TargetDir = "."
	if len(os.Args) > 2 {
		TargetDir = os.Args[2]
	}
	TargetDir, _ = filepath.Abs(TargetDir)
	ProjectDir, _ = filepath.Abs(path.Join(TargetDir, dstRelDir))

	if len(os.Args) > 3 {
		port := os.Args[3]
		ServerPort, _ = strconv.Atoi(port)
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
		if err.Error() != "" {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		return
	}

	err = cmd.SetupEnv(appName, version, fs, fsRelDir, TargetDir, dstRelDir)
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
