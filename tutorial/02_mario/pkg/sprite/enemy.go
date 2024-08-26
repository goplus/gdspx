package sprite

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Enemy struct {
	Sprite
}

func (pself *Enemy) OnStart() {
	println("Enemy.OnStart", pself.GetPosX(), pself.GetPosY())
}

func (pself *Enemy) OnUpdate(delta float32) {
	pself.AddPosX(100 * delta)
}

func (pself *Enemy) OnHit() {
}

func (pself *Enemy) Die() {
	println("Enemy.Die")
}