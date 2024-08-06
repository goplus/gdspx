package launcher

import (
	. "godot-ext/gd4go/internal/engine"
	"godot-ext/gd4go/internal/manager"
)

var (
	mgrs []IManager
)

func initEngine(pself *EngineNode) {
	manager.InitMgrs()
}

func tickEngine(delta float32) {
	manager.TickMgrs(delta)
}
