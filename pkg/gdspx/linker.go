package gdspx

import (
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/internal/engine"
)
type EngineCallbacks EngineCallbackInfo

func LinkEngine(callback EngineCallbacks) {
	engine.Link(EngineCallbackInfo(callback))
}
