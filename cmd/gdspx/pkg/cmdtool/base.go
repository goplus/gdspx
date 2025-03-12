package cmdtool

import (
	"embed"
	"os"
	"path"
	"path/filepath"

	_ "embed"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/util"
)

var (
	curCmd     ICmdTool
	appName    string
	Version    string
	ServerPort int = 8005

	proejctFS  embed.FS
	TargetDir  string
	ProjectDir string
	GoDir      = ""

	CmdPath string
	LibPath string

	projectRelPath = "/project"
	binPostfix     = ""
)

type ICmdTool interface {
	Register()
	CheckCmd(ext ...string) error
	CheckEnv() error
	SetupEnv(appName, version string, fs embed.FS, fsRelDir string, TargetDir string, projectRelPath string) error

	// import
	ShouldReimport() bool
	Reimport()
	// util
	ShowHelpInfo()
	Clear()

	// check
	OnBeforeCheck(cmd string) error

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

	// mobile
	ExportIos() error
	ExportApk() error

	// check
	OnAfterCheck(cmd string) error
}

type BaseCmdTool struct {
}

func (pself *BaseCmdTool) Register() {
	curCmd = pself
}
func (pself *BaseCmdTool) OnBeforeCheck(cmd string) error {
	return nil
}
func (pself *BaseCmdTool) OnAfterCheck(cmd string) error {
	return nil
}
func (pself *BaseCmdTool) CheckCmd(ext ...string) (err error) {
	if len(os.Args) <= 1 {
		pself.ShowHelpInfo()
		return
	}
	if !CheckCmd(ext...) {
		println("invalid cmd, please refer to help")
		pself.ShowHelpInfo()
	}
	return
}

func (pself *BaseCmdTool) CheckEnv() (err error) {
	return
}

func (pself *BaseCmdTool) SetupEnv(binName, pversion string, fs embed.FS, fsRelDir string, pTargetDir string, pProjectRelPath string) (err error) {
	TargetDir = pTargetDir
	appName = binName
	proejctFS = fs
	Version = pversion
	projectRelPath = pProjectRelPath
	err = SetupEnv()
	PrepareEnv(fsRelDir, ProjectDir)
	if pself.ShouldReimport() {
		pself.Reimport()
	}
	return
}
func (pself *BaseCmdTool) ShouldReimport() bool {
	return !util.IsFileExist(path.Join(ProjectDir, ".godot/uid_cache.bin"))
}
func (pself *BaseCmdTool) Reimport() {
	ImportProj()
}

func (pself *BaseCmdTool) ShowHelpInfo() {
	ShowHelpInfo(appName, Version)
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
	if !util.IsFileExist(filepath.Join(ProjectDir, ".builds", "web")) {
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

func (pself *BaseCmdTool) ExportApk() (err error) {
	ExportApk()
	return
}

func (pself *BaseCmdTool) ExportIos() (err error) {
	ExportIos()
	return
}
