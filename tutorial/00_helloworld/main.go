package main

import (
	. "github.com/realdream-ai/gdspx/pkg/engine"
	"github.com/realdream-ai/gdspx/pkg/gdspx"
)

func main() {
	RegisterTypes()
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   onStart,
		OnEngineUpdate:  onUpdate,
		OnEngineDestroy: onDestory,
	})
}

func RegisterTypes() {

}
func onStart() {
	println("hello world!")
}

func onUpdate(delta float32) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("goodbye world!")
}
