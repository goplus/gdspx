package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Enemy struct {
	Sprite
}

func (pself *Enemy) OnUpdate(delta float32) {
	pos := pself.GetPosition()
	pos.Y -= 100 * delta
	pself.SetPosition(pos)
}
func (pself *Enemy) OnDestory() {
	println("Enemy OnDestory")
}
