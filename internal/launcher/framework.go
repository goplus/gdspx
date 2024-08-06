package launcher

import (
	. "godot-ext/gd4go/internal/engine"
	"godot-ext/gd4go/internal/manager"

	"grow.graphics/gd"
)

var (
	mgrs []IManager
)

func initEngine(pself *EngineNode) {
	KeepAlive = pself.KeepAlive
	Temporary = gd.NewLifetime(KeepAlive)
	Root = pself.Super().AsNode()

	manager.InitMgrs()
}

func tickEngine(delta gd.Float) {
	Temporary.End()
	Temporary = gd.NewLifetime(KeepAlive)
	manager.TickMgrs(delta)
}
