package main

import (
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	gdspx.LinkEngine(gdspx.EngineCallbacks{
		OnEngineStart: onStart,
	})
}

func onStart() {
	println("onEngineStart")
}
