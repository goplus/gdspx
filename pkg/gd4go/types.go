package gd4go

import (
	"godot-ext/gd4go/internal/launcher"

	"grow.graphics/gd"
)

func RegisterEngineTypes(godot gd.Lifetime) {
	gd.Register[launcher.EngineNode](godot)
}
