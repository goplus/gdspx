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
	AudioUtil = &AudioMgr{BasicMgr{Root: node, KeepAlive: pself.KeepAlive}}
	AnimationUtil = &AnimationMgr{BasicMgr{Root: node, KeepAlive: pself.KeepAlive}}
	PhysicUtil = &PhysicMgr{BasicMgr{Root: node, KeepAlive: pself.KeepAlive}}
	InputUtil = &InputMgr{BasicMgr{Root: node, KeepAlive: pself.KeepAlive}}
	RenderUtil = &RenderMgr{BasicMgr{Root: node, KeepAlive: pself.KeepAlive}}
	println("initEngine")

}

func tickEngine(delta gd.Float) {
	Temporary.End()
	Temporary = gd.NewLifetime(rootSelf.KeepAlive)
}
