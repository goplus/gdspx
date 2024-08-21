package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Aircraft struct {
	Sprite
	timer float32
}

func (pself *Aircraft) OnStart() {
	println("Aircraft OnStart")
}
func (pself *Aircraft) OnUpdate(delta float32) {
	pself.timer+= delta
	if pself.timer > 0.3 {
		pself.timer = 0
		bullet := CreateSprite[Bullet]()
		pos := pself.GetPosition()
		bullet.SetPosition(Vec2{pos.X , pos.Y + 150})
	}
}
func (pself *Aircraft) OnDestory() {
	println("Aircraft OnDestory")
}
