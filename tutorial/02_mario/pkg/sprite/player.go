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

var (
	gravity = float32(-980)
	SPEED = float32(300)
	JUMP_VELOCITY = float32(400.0)
)


func (pself *Player) OnStart() {
	pself.moveSpeed = 600
	pself.OnDieEvent = NewEvent0()
}

func (pself *Player) OnUpdate(delta float32) {
	if !pself.IsOnFloor() {
		pself.AddVelY(gravity * delta)
	}
	if InputMgr.GetKey(KeyCode.Space) && pself.IsOnFloor() {
		pself.SetVelY(JUMP_VELOCITY)
	}

	dir := InputMgr.GetAxis("left", "right")
	if dir != 0 {
		pself.SetVelX(dir * SPEED)
	} else {
		pself.SetVelX(MoveToward(pself.GetVelX(), 0, SPEED))
	}
	pself.MoveAndSlide()
}

func (pself *Player) OnHit() {
	pself.OnDieEvent.Trigger()
}
