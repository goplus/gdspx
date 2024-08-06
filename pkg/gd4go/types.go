package gd4go

import (
	"godot-ext/gd4go/internal/launcher"

	"github.com/godot-go/godot-go/pkg/core"
)

func RegisterEngineTypes() {
	core.AutoRegisterClassDB[*launcher.EngineNode]()
}
