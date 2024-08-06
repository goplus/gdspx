package gdspx

import (
	"godot-ext/gdspx/internal/launcher"

	"github.com/godot-go/godot-go/pkg/core"
)

func RegisterEngineTypes() {
	core.AutoRegisterClassDB[*launcher.EngineNode]()
}
