package main

import (
	"goplus/gd4spx"

	"grow.graphics/gd/gdextension"
)

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
	gd4spx.RegisterEngineTypes(godot)
}
