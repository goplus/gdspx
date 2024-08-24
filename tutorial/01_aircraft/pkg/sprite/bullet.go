package sprite

import (
	. "gdspx-demo01/pkg/define"
	. "godot-ext/gdspx/pkg/engine"
)

type Bullet struct {
	Sprite
}

func (pself *Bullet) OnStart() {
}

func (pself *Bullet) OnUpdate(delta float32) {
	pself.AddPos(0, 2000*delta)
	if pself.GetPosY() > WinHeight {
		pself.Destroy()
	}
}

func (pself *Bullet) OnTriggerEnter(target ISpriter) {
	if enemy, ok := target.(*Enemy); ok {
		enemy.OnHit()
		pself.Destroy()
	}
}
