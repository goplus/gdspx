package gdspx

import (
	inengine "godot-ext/gdspx/internal/engine"
	. "godot-ext/gdspx/pkg/engine"
)

func LinkEngine(callback EngineCallbackInfo) {
	inengine.Link(EngineCallbackInfo(callback))
}
