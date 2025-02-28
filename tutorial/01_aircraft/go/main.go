package main

import "C"
import (
	"gdspx-demo01/pkg/game"

	. "github.com/realdream-ai/gdspx/pkg/engine"
	"github.com/realdream-ai/gdspx/pkg/gdspx"
)

func main() {
	doinit()
	game.RegisterTypes()
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   game.OnStart,
		OnEngineUpdate:  game.OnUpdate,
		OnEngineDestroy: game.OnDestroy,
	})
}
