package engine

import (
	"grow.graphics/gd"
)

type IManager interface {
	Init(root gd.Node, keepAlive gd.Lifetime)
	Ready()
	Process(delta gd.Float)
}

type IAudioMgr interface {
	PlayAudio(name string)
	SetAudioVolume(volume float32)
	GetAudioVolume() float32

	PlayMusic(name string)
	PauseMusic()
	StopMusic()
	SetMusicVolume(volume float32)
	GetMusicVolume() float32
	SetMusicTime(time float32)
	GetMusicTime() float32
}

type IAnimationMgr interface {
}
type IPhysicMgr interface{}
type IInputMgr interface{}
type IRenderMgr interface{}
type ISpriteMgr interface{}
