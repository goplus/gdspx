package launcher

import (
	. "godot-ext/gd4go/internal/engine"
	"godot-ext/gd4go/internal/manager"

	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

var (
	mgrs []IManager
)

func initEngine(pself *EngineNode) {
	KeepAlive = pself.KeepAlive
	Temporary = gd.NewLifetime(KeepAlive)
	Root = pself.Super().AsNode()

	mgrs = manager.InitMgrs()
	for _, mgr := range mgrs {
		mgr.Ready()
	}
}

func tickEngine(delta gd.Float) {
	Temporary.End()
	Temporary = gd.NewLifetime(KeepAlive)
	for _, mgr := range mgrs {
		mgr.Process(delta)
	}
}
