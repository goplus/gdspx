package main

import (
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	gdspx.LinkEngine(gdspx.EngineCallbacks{
		OnEngineStart: onStart,
	})
}

func onStart() {
	println("onEngineStart")
	SpriteMgr.CreateSprite("TestSprite")
}
