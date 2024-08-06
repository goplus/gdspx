package manager

import (
	. "godot-ext/gd4go/internal/engine"

	"grow.graphics/gd"
)

var (
	managers []IManager
)

func addManager[T IManager](mgr T) T {
	mgr.Init(Root, KeepAlive)
	managers = append(managers, mgr)
	return mgr
}
func createMgrs() {
	AudioMgr = addManager(&audioMgr{})
	AnimationMgr = addManager(&animationMgr{})
	PhysicMgr = addManager(&physicMgr{})
	InputMgr = addManager(&inputMgr{})
	RenderMgr = addManager(&renderMgr{})
	SpriteMgr = addManager(&spriteMgr{})
}

func InitMgrs() {
	createMgrs()
	for _, mgr := range managers {
		mgr.Ready()
	}
}

func TickMgrs(delta gd.Float) {
	for _, mgr := range managers {
		mgr.Process(delta)
	}
}
