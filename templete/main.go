package main

import (
	"goplus/gd4spx/engine"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
	gd.Register[engine.EngineNode](godot)

}
