package gdspx

import (
	"godot-ext/gdspx/internal/launcher"
	"godot-ext/gdspx/internal/node"

	"github.com/godot-go/godot-go/pkg/core"
)

func RegisterEngineTypes() {
	core.AutoRegisterClassDB[*launcher.EngineNode]()
	// register nodes
	core.AutoRegisterClassDB[*node.SpxNode]()
	core.AutoRegisterClassDB[*node.SpxAudioPalyer]()

}
