package engine

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

var (
	Root Node
	// managers
	ResMgr       IResMgr
	AudioMgr     IAudioMgr
	AnimationMgr IAnimationMgr
	PhysicMgr    IPhysicMgr
	RenderMgr    IRenderMgr
	InputMgr     IInputMgr
	SpriteMgr    ISpriteMgr
)
