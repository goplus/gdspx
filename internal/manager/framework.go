package manager

import (
	. "godot-ext/gdspx/internal/engine"
	"reflect"
)

var (
	managers []IManager
)

func addManager[T IManager](mgr T) T {
	typeName := reflect.TypeOf(mgr).Elem().Name()
	println("add manager", typeName)
	mgr.Init(nil)
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

func TickMgrs(delta float32) {
	for _, mgr := range managers {
		mgr.Process(delta)
	}
}
