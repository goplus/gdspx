package main

import (
	"gdspx-demo02/pkg/game"

	. "github.com/realdream-ai/gdspx/pkg/engine"
	"github.com/realdream-ai/gdspx/pkg/gdspx"
)

func main() {
	game.RegisterTypes()
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   game.OnStart,
		OnEngineUpdate:  game.OnUpdate,
		OnEngineDestroy: game.OnDestroy,
	})
}
