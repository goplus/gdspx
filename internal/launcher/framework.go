package launcher

import (
	. "godot-ext/gdspx/internal/engine"
	"godot-ext/gdspx/internal/manager"
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
