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
	obj := SpriteMgr.CreateSprite("TestSprite中文aaA")
	SpriteMgr.CreateSprite("TestSprite中文aBBBBA")
	SpriteMgr.UpdateZIndex(obj, 1)

}

func onUpdate(delta float64) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
