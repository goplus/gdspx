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
