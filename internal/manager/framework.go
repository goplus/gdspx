package manager

import (
	. "godot-ext/gdspx/internal/engine"
	"reflect"
)

var (
	managers []IManager
)

func createNode(typeName string) Node {
	return 0 // TODO
}

func addManager[T IManager](mgr T) T {
	typeName := reflect.TypeOf(mgr).Elem().Name()
	println("add manager", typeName)
	node := createNode("SpxNode")

	// TODO(jiepengtan) fix godot-go bug
	//Root.AddChild(node, true, NodeInternalMode(0))
	mgr.Init(node)
	managers = append(managers, mgr)
	return mgr
}
func createMgrs() {
	ResMgr = addManager(&resMgr{})
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
