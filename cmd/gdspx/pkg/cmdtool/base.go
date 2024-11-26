package cmdtool

import (
	"embed"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/util"
)

var (
	curCmd     ICmdTool
	version    string
	appName    string
	serverPort int = 8005

	proejctFS  embed.FS
	targetDir  string
	projectDir string
	goDir      = ""

	cmdPath string
	libPath string

	binPostfix = ""
)

type ICmdTool interface {
	Register()
	CheckCmd(ext ...string) error
	CheckEnv() error
	SetupEnv(appName, version string, fs embed.FS, fsRelDir string, targetDir string) error

	// util
	ShowHelpInfo()
	Clear()

	// check
	OnBeforeCheck() error

	// pc
	BuildDll() error
	Run(string) error
	Export() error

	// web
	BuildWasm() error
	RunWeb() error
	ExportWeb() error
	StopWeb() error
	CheckExportWeb() error

	// check
	OnAfterCheck() error
}

type BaseCmdTool struct {
}

func (pself *BaseCmdTool) Register() {
	curCmd = pself
}
func (pself *BaseCmdTool) OnBeforeCheck() error {
	return nil
}
func (pself *BaseCmdTool) OnAfterCheck() error {
	return nil
}
func (pself *BaseCmdTool) CheckCmd(ext ...string) (err error) {
	if len(os.Args) <= 1 {
		pself.ShowHelpInfo()
		return
	}
	if !CheckCmd() {
		println("invalid cmd, please refer to help")
		pself.ShowHelpInfo()
	}
	return
}

func (pself *BaseCmdTool) CheckEnv() (err error) {
	return
}

func (pself *BaseCmdTool) SetupEnv(binName, pversion string, fs embed.FS, fsRelDir string, pTargetDir string) (err error) {
	targetDir = pTargetDir
	appName = binName
	proejctFS = fs
	version = pversion
	PrepareEnv(fsRelDir, targetDir)
	err = SetupEnv()
	return
}

func (pself *BaseCmdTool) ShowHelpInfo() {
	ShowHelpInfo(appName, version)
	return
}

func (pself *BaseCmdTool) Clear() {
	Clear()
	return

}

func (pself *BaseCmdTool) BuildDll() (err error) {
	BuildDll()
	return
}

func (pself *BaseCmdTool) Run(arg string) (err error) {
	Run(arg)
	return
}
func (pself *BaseCmdTool) Export() (err error) {
	Export()
	return
}

func (pself *BaseCmdTool) BuildWasm() (err error) {
	BuildWasm()
	return
}
func (pself *BaseCmdTool) RunWeb() (err error) {
	if !util.IsFileExist(filepath.Join(projectDir, ".builds", "web")) {
		ExportWeb()
	}
	RunWeb()
	return
}
func (pself *BaseCmdTool) ExportWeb() (err error) {
	ExportWeb()
	return
}
func (pself *BaseCmdTool) StopWeb() (err error) {
	StopWeb()
	return

}
func (pself *BaseCmdTool) CheckExportWeb() (err error) {
	CheckExportWeb()
	return
}
