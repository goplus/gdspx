package manager

import (
)

type audioMgr struct {
	baseMgr
	//Music *SpxAudioPalyer
	//Audio *SpxAudioPalyer
}

func (pself *audioMgr) PlayAudio(name string) {
	//ResourceLoader.Load(name)
}
func (pself *audioMgr) SetAudioVolume(volume float32) {
}
func (pself *audioMgr) GetAudioVolume() float32 {
	return 0
}

func (pself *audioMgr) PlayMusic(name string) {
	// TODO(jiepengtan): implement this
}
func (pself *audioMgr) SetMusicStreamPause(isPause bool) {
}
func (pself *audioMgr) GetMusicStreamPause() bool {
	return false
}
func (pself *audioMgr) StopMusic() {
}
func (pself *audioMgr) SetMusicVolume(volume float32) {
}
func (pself *audioMgr) GetMusicVolume() float32 {
	return 0
}
func (pself *audioMgr) SetMusicTime(time float32) {
}
func (pself *audioMgr) GetMusicTime() float32 {
	// TODO(jiepengtan): implement this
	return 0
}

func (pself *audioMgr) Ready() {
}
