package launcher

import (
	. "godot-ext/gd4go/internal/engine"
	. "godot-ext/gd4go/internal/manager"

	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

var (
	managers []IManager
)

func addManager[T IManager](mgr T) T {
	mgr.Init(Root, KeepAlive)
	managers = append(managers, mgr)
	return mgr
}
func initEngine(pself *EngineNode) {
	KeepAlive = pself.KeepAlive
	Temporary = gd.NewLifetime(KeepAlive)
	Root = pself.Super().AsNode()

	AudioUtil = addManager(&AudioMgr{})
	AnimationUtil = addManager(&AnimationMgr{})
	PhysicUtil = addManager(&PhysicMgr{})
	InputUtil = addManager(&InputMgr{})
	RenderUtil = addManager(&RenderMgr{})
	SpriteUtil = addManager(&SpriteMgr{})

	for _, mgr := range managers {
		mgr.Ready()
	}
}

func tickEngine(delta gd.Float) {
	Temporary.End()
	Temporary = gd.NewLifetime(KeepAlive)
	for _, mgr := range managers {
		mgr.Process(delta)
	}
}
