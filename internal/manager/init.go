package manager

import (
	. "godot-ext/gd4go/internal/engine"
)

func addManager[T IManager](managers []IManager, mgr T) T {
	mgr.Init(Root, KeepAlive)
	managers = append(managers, mgr)
	return mgr
}
func InitMgrs() []IManager {
	var managers []IManager

	AudioMgr = addManager(managers, &audioMgr{})
	AnimationMgr = addManager(managers, &animationMgr{})
	PhysicMgr = addManager(managers, &physicMgr{})
	InputMgr = addManager(managers, &inputMgr{})
	RenderMgr = addManager(managers, &renderMgr{})
	SpriteMgr = addManager(managers, &spriteMgr{})

	return managers
}
