package launcher

import (
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
)

type EngineNode struct {
	NodeImpl
}

func (pself *EngineNode) V_ready() {
	initEngine(pself)
}

func (pself *EngineNode) V_process(delta float32) {
	tickEngine(delta)
}
