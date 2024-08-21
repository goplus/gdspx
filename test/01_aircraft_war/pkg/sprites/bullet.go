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
	pself.Move(0,2000 * delta)
	if(pself.GetPosY() > WinHeight){
		pself.Destroy()
	}
}

func (pself *Bullet) OnTriggerEnter(target ISpriter) {
	if enemy, ok := target.(*Enemy); ok {
		enemy.OnHit()
		pself.Destroy()
	}
}