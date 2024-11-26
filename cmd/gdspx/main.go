package main

import (
	"embed"
	_ "embed"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/cmdtool"
	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/impl"
)

var (
	//go:embed template/project/*
	proejct_fs embed.FS

	//go:embed template/version
	version string
)

type CmdTool struct {
	cmdtool.BaseCmdTool
}

func (pself *CmdTool) OnBeforeCheck() error {
	impl.UpdateMod(
		"github.com/realdream-ai/gdspx ", "./spx",
		"github.com/realdream-ai/gdspx/cmd/gdspx ",
		[]string{"", "cmd/spx", "cmd/ispx"},
		[]string{"cmd/spx/template/project/go.mod.txt"})
	return nil
}

func main() {
	cmd := &CmdTool{}
	cmdtool.RunCmd(cmd, "gdspx", version, proejct_fs, "template/project")
}
