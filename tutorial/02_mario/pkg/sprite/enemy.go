package sprite

import (
	. "gdspx-demo02/pkg/define"
	. "godot-ext/gdspx/pkg/engine"
)

type Enemy struct {
	Sprite
	IsDied     bool
	dyingTimer float32
}

func (pself *Enemy) OnUpdate(delta float32) {
	if pself.IsDied {
		pself.dyingTimer -= delta
		if pself.dyingTimer < 0 {
			pself.Destroy()
		}
		return
	}
	pself.Move(0, -100*delta)
	if pself.GetPosY() < -WinHeight {
		pself.Destroy()
	}
}

func (pself *Enemy) OnTriggerEnter(target ISpriter) {
	if item, ok := target.(*Aircraft); ok {
		item.OnHit()
	}
}

func (pself *Enemy) OnHit() {
	pself.dyingTimer = 0.2
	pself.IsDied = true
	Score += 100
	pself.PlayAnim("die", 3, false)
}