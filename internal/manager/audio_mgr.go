package manager

import (
	"grow.graphics/gd"
)

type audioMgr struct {
	baseMgr
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

func (pself *audioMgr) Process(delta gd.Float) {}
