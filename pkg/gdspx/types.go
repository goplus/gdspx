package gdspx

import (
	. "godot-ext/gdspx/internal/core"
	"godot-ext/gdspx/internal/engine"
)

type EngineCallbacks EngineCallbackInfo

var (
	AudioMgr  IAudioMgr
	UIMgr     IUIMgr
	PhysicMgr IPhysicMgr
	InputMgr  IInputMgr
	SpriteMgr ISpriteMgr
)

func LinkEngine(callback EngineCallbacks) {
	mgrs := engine.Link(EngineCallbackInfo(callback))
	for _, mgr := range mgrs {
		switch v := mgr.(type) {
		case IAudioMgr:
			AudioMgr = v
		case IUIMgr:
			UIMgr = v
		case IPhysicMgr:
			PhysicMgr = v
		case IInputMgr:
			InputMgr = v
		case ISpriteMgr:
			SpriteMgr = v
		default:
			panic("engine init error : unknown manager")
		}
	}
}
