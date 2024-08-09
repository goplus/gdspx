package main

import (
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	gdspx.LinkEngine(gdspx.EngineCallbacks{
		OnEngineStart:   onStart,
		OnEngineUpdate:  onUpdate,
		OnEngineDestroy: onDestory,
	})
}

func onStart() {
	println("onEngineStart")
	SpriteMgr.CreateSprite("TestSprite中文")
}

func onUpdate(delta float64) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
