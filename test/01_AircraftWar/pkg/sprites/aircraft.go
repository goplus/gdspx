package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Aircraft struct {
	Sprite
}

func (pself *Aircraft) OnStart() {
	println("Aircraft OnStart")
}
func (pself *Aircraft) OnUpdate(delta float32) {
	println("Aircraft OnUpdate")
}
func (pself *Aircraft) OnDestory() {
	println("Aircraft OnDestory")
}
