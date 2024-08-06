package engine

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

type IManager interface {
	Init(root Node)
	Ready()
	Process(delta float32)
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
