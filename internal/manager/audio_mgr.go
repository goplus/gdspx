package manager

import (
	. "godot-ext/gdspx/internal/node"
)

type audioMgr struct {
	baseMgr
	Music *SpxAudioPalyer
	Audio *SpxAudioPalyer
}

func (pself *audioMgr) PlayAudio(name string) {
	//ResourceLoader.Load(name)
}
func (pself *audioMgr) SetAudioVolume(volume float32) {
	pself.Audio.SetVolumeDb(volume)
}
func (pself *audioMgr) GetAudioVolume() float32 {
	return pself.Audio.GetVolumeDb()
}

func (pself *audioMgr) PlayMusic(name string) {
	// TODO(jiepengtan): implement this
}
func (pself *audioMgr) SetMusicStreamPause(isPause bool) {
	pself.Music.SetStreamPaused(isPause)
}
func (pself *audioMgr) GetMusicStreamPause() bool {
	return pself.Music.GetStreamPaused()
}
func (pself *audioMgr) StopMusic() {
	pself.Music.Stop()
}
func (pself *audioMgr) SetMusicVolume(volume float32) {
	pself.Music.SetVolumeDb(volume)
}
func (pself *audioMgr) GetMusicVolume() float32 {
	return pself.Music.GetVolumeDb()
}
func (pself *audioMgr) SetMusicTime(time float32) {
	pself.Music.Seek(time)
}
func (pself *audioMgr) GetMusicTime() float32 {
	// TODO(jiepengtan): implement this
	return 0
}

func (pself *audioMgr) Ready() {
	pself.Music = createNode("SpxAudioPalyer").(*SpxAudioPalyer)
	pself.Audio = createNode("SpxAudioPalyer").(*SpxAudioPalyer)
}
