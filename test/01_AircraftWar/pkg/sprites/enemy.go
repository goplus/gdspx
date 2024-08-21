package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Enemy struct {
	Sprite
}
func (pself *Enemy) OnStart() {
	println("Enemy OnStart")
}
func (pself *Enemy) OnUpdate(delta float32) {
	println("Enemy OnUpdate",pself.GetId())
	pos := pself.GetPosition()
	pos.Y -= 10 * delta
	pself.SetPosition(pos)
}
func (pself *Enemy) OnDestory() {
	println("Enemy OnDestory")
}
