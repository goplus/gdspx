package main

import (
	"gdspx-demo01/pkg/game"
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	game.RegisterTypes()
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   game.OnStart,
		OnEngineUpdate:  game.OnUpdate,
		OnEngineDestroy: game.OnDestroy,
	})
}
