package sprite

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Player struct {
	Sprite
	timer      float32
	moveSpeed  float32
	OnDieEvent *Event0
}

func (pself *Player) OnStart() {
	pself.moveSpeed = 600
	pself.OnDieEvent = NewEvent0()
}
func (pself *Player) OnUpdate(delta float32) {
	pself.timer += delta
	if InputMgr.GetKey(KeyCode.W) {
		pself.Move(0, pself.moveSpeed*delta)
	}
	if InputMgr.GetKey(KeyCode.S) {
		pself.Move(0, -pself.moveSpeed*delta)
	}
	if InputMgr.GetKey(KeyCode.D) {
		pself.Move(pself.moveSpeed*delta, 0)
	}
	if InputMgr.GetKey(KeyCode.A) {
		pself.Move(-pself.moveSpeed*delta, 0)
	}
}


func (pself *Player) OnHit() {
	pself.OnDieEvent.Trigger()
}
