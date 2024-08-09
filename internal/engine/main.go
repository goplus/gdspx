package engine

import (
	"reflect"
	"fmt"
	"godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/pkg/engine"
	. "godot-ext/gdspx/internal/wrap"
)

var (
	mgrs     []IManager
	callback EngineCallbackInfo
)

func Link(engineCallback EngineCallbackInfo) []IManager {
	ok := ffi.Link()
	if !ok {
		panic("godot bind symbol failed!")
	}
	mgrs = CreateMgrs()
	callback = engineCallback
	RegisterCallbacks(EngineCallbackInfo{
		OnEngineStart:   onEngineStart,
		OnEngineUpdate:  onEngineUpdate,
		OnEngineDestroy: onEngineDestroy,
	})
	
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
			panic(fmt.Sprintf("engine init error : unknown manager type %s", reflect.TypeOf(mgr).String()))
		}
	}
	return mgrs
}

func onEngineStart() {
	for _, mgr := range mgrs {
		mgr.OnStart()
	}
	if callback.OnEngineStart != nil {
		callback.OnEngineStart()
	}
}

func onEngineUpdate(delta float64) {
	for _, mgr := range mgrs {
		mgr.OnUpdate(delta)
	}
	if callback.OnEngineUpdate != nil {
		callback.OnEngineUpdate(delta)
	}
}
func onEngineDestroy() {
	if callback.OnEngineDestroy != nil {
		callback.OnEngineDestroy()
	}
	for _, mgr := range mgrs {
		mgr.OnDestroy()
	}
}
