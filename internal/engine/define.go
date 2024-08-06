package engine

import (
	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

var (
	Root      gd.Node
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime

	// managers
	AudioMgr     IAudioMgr
	AnimationMgr IAnimationMgr
	PhysicMgr    IPhysicMgr
	RenderMgr    IRenderMgr
	InputMgr     IInputMgr
	SpriteMgr    ISpriteMgr
)
