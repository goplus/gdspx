package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
	. "gdspx-demo01/pkg/define"
)

type Bullet struct {
	Sprite
}

func (pself *Bullet) OnStart() {
}

func (pself *Bullet) OnUpdate(delta float32) {
	pos := pself.GetPosition()
	pos.Y += 2000 * delta
	pself.SetPosition(pos)
	if(pos.Y > WinHeight){
		pself.Destroy()
	}
}

func (pself *Bullet) OnTriggerEnter(target ISpriter) {
	if enemy, ok := target.(*Enemy); ok {
		enemy.OnHit()
		pself.Destroy()
	}
}