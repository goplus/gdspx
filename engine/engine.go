package engine

import (
	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

func initEngine(pself *EngineNode) {
	KeepAlive = pself.KeepAlive
	Temporary = gd.NewLifetime(KeepAlive)
	node := pself.Super().AsNode2D()
	Root = node
	rootSelf = pself
	managersList := []IManager{
		&AudioMgr{},
		&AnimationMgr{},
		&PhysicMgr{},
		&InputMgr{},
		&RenderMgr{},
	}
	for _, mgr := range managersList {
		mgr.Init(Root, KeepAlive)
		managers = append(managers, mgr)
	}
	for _, mgr := range managers {
		mgr.Ready()
	}
}

func tickEngine(delta gd.Float) {
	Temporary.End()
	Temporary = gd.NewLifetime(rootSelf.KeepAlive)
	for _, mgr := range managers {
		mgr.Process(delta)
	}
}
