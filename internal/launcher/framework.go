package launcher

import (
	. "godot-ext/gdspx/internal/engine"
	"godot-ext/gdspx/internal/manager"
)

var (
	mgrs []IManager
)

func initEngine(pself *EngineNode) {
	Root = 0 // TODO(jiepengtan): get root node from godot
	manager.InitMgrs()
}

func tickEngine(delta float32) {
	manager.TickMgrs(delta)
}
