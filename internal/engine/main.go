package engine

import (
	"godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/internal/wrap"
	. "godot-ext/gdspx/pkg/engine"
)

var (
	mgrs     []IManager
	callback EngineCallbackInfo
	sprites   = make([]ISpriter, 0)
	timer	 = float32(0)
)

func Link(engineCallback EngineCallbackInfo) []IManager {
	ok := ffi.Link()
	if !ok {
		panic("godot bind symbol failed!")
	}
	mgrs = CreateMgrs()
	callback = engineCallback
	infos := bindCallbacks()
	RegisterCallbacks(infos)
	BindMgr(mgrs)
	InitEngine()
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

func onEngineUpdate(delta float32) {
	for _, mgr := range mgrs {
		mgr.OnUpdate(delta)
	}
	TimeSinceGameStart += delta
	sprites = sprites[:0]
	for _, sprite := range Id2Sprites {
		sprites = append(sprites, sprite)
	}
	for _, sprite := range sprites {
		sprite.OnUpdate(delta)
	}
	if callback.OnEngineUpdate != nil {
		callback.OnEngineUpdate(delta)
	}
}
func onEngineDestroy() {
	if callback.OnEngineDestroy != nil {
		callback.OnEngineDestroy()
	}
	sprites = sprites[:0]
	for _, sprite := range Id2Sprites {
		sprites = append(sprites, sprite)
	}
	for _, sprite := range sprites {
		sprite.OnDestroy()
	}
	for _, mgr := range mgrs {
		mgr.OnDestroy()
	}
}
