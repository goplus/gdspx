package demo

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

// @autobind signal "start_game"
type HUD struct {
	CanvasLayerImpl
	ScoreLabel   Label  `gdbind:"ScoreLabel"`
	MessageLabel Label  `gdbind:"MessageLabel"`
	MessageTimer Timer  `gdbind:"MessageTimer"`
	StartButton  Button `gdbind:"StartButton"`
}
