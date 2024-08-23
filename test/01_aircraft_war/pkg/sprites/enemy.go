package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
	. "gdspx-demo01/pkg/define"
)

type Enemy struct {
	Sprite
	IsDied bool
	dyingTimer float32
}

func (pself *Enemy) OnUpdate(delta float32) {
	if pself.IsDied {
		pself.dyingTimer -= delta
		if(pself.dyingTimer < 0){
			pself.Destroy()
		}
		return
	}
	pself.Move(0, -100 * delta)
	if pself.GetPosY() < -WinHeight {
		pself.Destroy()
	}
}

func (pself *Enemy) OnHit() {
	pself.dyingTimer = 0.2
	pself.IsDied = true
	Score += 100
	pself.PlayAnim("die", 3, false)
}

func (pself *Enemy) OnDestory() {
	
}
