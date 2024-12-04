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

type returnDirect struct{}

func (p *returnDirect) Error() string {
	return ""
}
func (pself *CmdTool) OnBeforeCheck(cmd string) error {
	switch cmd {
	case "updatemod":
		impl.UpdateMod(
			"github.com/realdream-ai/gdspx ", "./spx",
			"github.com/realdream-ai/gdspx/cmd/gdspx ",
			[]string{"", "cmd/spx", "cmd/ispx"},
			[]string{"cmd/spx/template/project/go.mod.txt"})
		return &returnDirect{}
	}
	return nil
}

func main() {
	cmd := &CmdTool{}
	cmdtool.RunCmd(cmd, "gdspx", version, proejct_fs, "template/project", "", "updatemod")
}
