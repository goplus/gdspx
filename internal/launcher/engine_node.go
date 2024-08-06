package launcher

import (
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type EngineNode struct {
	NodeImpl
}

func (pself *EngineNode) Ready() {
	initEngine(pself)
}

func (pself *EngineNode) Process(delta float32) {
	tickEngine(delta)
}
