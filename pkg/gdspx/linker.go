package gdspx

import (
	inengine "github.com/realdream-ai/gdspx/internal/engine"
	. "github.com/realdream-ai/gdspx/pkg/engine"
)

func LinkEngine(callback EngineCallbackInfo) {
	inengine.Link(EngineCallbackInfo(callback))
}
