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
	SpriteMgr.SetPosition(obj,Vec2{100,100})
	SpriteMgr.SetPosition(obj2,Vec2{100,200})
	SpriteMgr.UpdateZIndex(obj, 1)

}

func onUpdate(delta float64) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
