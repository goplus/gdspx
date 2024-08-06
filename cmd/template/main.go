package main

import (
	"godot-ext/gd4go/pkg/gd4go"

	"grow.graphics/gd/gdextension"
)

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
	gd4go.InitEngine(godot)
}
