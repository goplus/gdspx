package sprite

import (
	. "gdspx-demo01/pkg/define"

	. "github.com/realdream-ai/gdspx/pkg/engine"
)

type Bullet struct {
	Sprite
}

func (pself *Bullet) OnStart() {
}

func (pself *Bullet) OnUpdate(delta float64) {
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
