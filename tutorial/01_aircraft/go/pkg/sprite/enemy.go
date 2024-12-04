package sprite

import (
	. "gdspx-demo01/pkg/define"

	. "github.com/realdream-ai/gdspx/pkg/engine"
)

type Enemy struct {
	Sprite
	IsDied     bool
	dyingTimer float64
}

func (pself *Enemy) OnUpdate(delta float64) {
	if pself.IsDied {
		pself.dyingTimer -= delta
		if pself.dyingTimer < 0 {
			pself.Destroy()
		}
		return
	}
	pself.AddPos(0, -100*delta)
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
	pself.PlayAnim("die", 3, false, false)
}
