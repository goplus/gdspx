package engine

import (
	. "godot-ext/gdspx/internal/core"
	"godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/internal/wrap"
)

var (
	mgrs     []IManager
	callback EngineCallbackInfo
)

func Link(engineCallback EngineCallbackInfo) []IManager {
	ok := ffi.BindAPI()
	if !ok {
		panic("godot not found")
	}
	mgrs = CreateMgrs()
	callback = engineCallback
	RegisterCallbacks(EngineCallbackInfo{
		OnEngineStart:   onEngineStart,
		OnEngineUpdate:  onEngineUpdate,
		OnEngineDestroy: onEngineDestroy,
	})
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
