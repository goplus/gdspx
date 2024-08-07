package launcher

import (
	. "godot-ext/gdspx/internal/engine"
)

type EngineNode struct {
	Node
}

func (pself *EngineNode) V_ready() {
	initEngine(pself)
}

func (pself *EngineNode) V_process(delta float32) {
	tickEngine(delta)
}
