package manager

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

type audioMgr struct {
	baseMgr
	Music AudioStreamPlayer
	Audio AudioStreamPlayer
}

func (pself *audioMgr) PlayAudio(name string)         {}
func (pself *audioMgr) SetAudioVolume(volume float32) {}
func (pself *audioMgr) GetAudioVolume() float32       { return 0 }

func (pself *audioMgr) PlayMusic(name string)         {}
func (pself *audioMgr) PauseMusic()                   {}
func (pself *audioMgr) StopMusic()                    {}
func (pself *audioMgr) SetMusicVolume(volume float32) {}
func (pself *audioMgr) GetMusicVolume() float32       { return 0 }
func (pself *audioMgr) SetMusicTime(time float32)     {}
func (pself *audioMgr) GetMusicTime() float32         { return 0 }

func (pself *audioMgr) Ready() {
}

func (pself *audioMgr) Process(delta float32) {

}
