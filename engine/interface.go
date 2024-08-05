package engine

import (
	"grow.graphics/gd"
)

var (
	Root      gd.Node
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime

	AudioUtil     IAudioMgr
	AnimationUtil IAnimationMgr
	PhysicUtil    IPhysicMgr
	RenderUtil    IRenderMgr
	InputUtil     IInputMgr
	SpriteUtil    ISpriteMgr

	managers []IManager
)

type IAudioMgr interface{}
type IAnimationMgr interface{}
type IPhysicMgr interface{}
type IInputMgr interface{}
type IRenderMgr interface{}
type ISpriteMgr interface{}
