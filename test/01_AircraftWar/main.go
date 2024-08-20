package main

import (
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   onStart,
		OnEngineUpdate:  onUpdate,
		OnEngineDestroy: onDestory,
	})
}

func onStart() {
	println("onEngineStart")
	obj := NewSprite("res://Assets/Prefabs/Enemy.tscn")
	obj.SetPosition(Vec2{-540, 0})
}

func onUpdate(delta float32) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
