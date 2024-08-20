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
	obj := SpriteMgr.CreateSprite("")
	obj2 := SpriteMgr.CreateSprite("")
	SpriteMgr.SetPosition(obj, Vec2{100, 100})
	SpriteMgr.SetPosition(obj2, Vec2{200, 200})
	SpriteMgr.SetTexture(obj, "res://icon.png")
	SpriteMgr.SetTexture(obj2, "res://icon.png")
}

func onUpdate(delta float32) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
