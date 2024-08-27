package sprite

import (
	. "gdspx-demo02/pkg/define"
	. "godot-ext/gdspx/pkg/engine"
)

type Goomba struct {
	Enemy
	horizontalSpeed float32
	verticalSpeed   float32
	is_visible      bool
}

func (pself *Goomba) OnStart() {
	pself.horizontalSpeed = 20
	pself.verticalSpeed = 100
	pself.is_visible = false
}

func (pself *Goomba) OnUpdate(delta float32) {
	if !pself.is_visible {
		return
	}
	pself.AddPosX(pself.horizontalSpeed * delta)
	if !pself.IsOnFloor() {
		pself.AddVelY(-pself.verticalSpeed * delta)
	}
	if PhysicMgr.CheckCollision(pself.GetPosition(), pself.GetPosition().AddX(Signf(pself.horizontalSpeed)*12), CollisionLayer_Pipe, true, true) {
		println("Collision")
		pself.horizontalSpeed = -pself.horizontalSpeed
	}
}

func (pself *Goomba) OnHit() {
}

func (pself *Goomba) Die() {
	pself.Destroy()
}

func (pself *Goomba) OnScreenExited() {
	pself.Destroy()
}

func (pself *Goomba) OnScreenEntered() {
	pself.SetProcess(true)
	pself.is_visible = true
}

func (pself *Goomba) OnTriggerEnter(target ISpriter) {

}
