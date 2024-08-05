package gd4spx

import (
	"goplus/gd4spx/launcher"

	"grow.graphics/gd"
)

func RegisterEngineTypes(godot gd.Lifetime) {
	gd.Register[launcher.EngineNode](godot)
}
