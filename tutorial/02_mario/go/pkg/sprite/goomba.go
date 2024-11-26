package sprite

import (
	. "gdspx-demo02/pkg/define"

	. "github.com/realdream-ai/gdspx/pkg/engine"
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
	if PhysicMgr.CheckCollision(pself.GetPosition(), pself.GetPosition().AddX(Signf(pself.horizontalSpeed)*6), CollisionLayer_Pipe, true, true) {
		println("Collision")
		pself.horizontalSpeed = -pself.horizontalSpeed
	}
}

func (pself *Goomba) OnHit() {
}

func (pself *Goomba) Die() {
	pself.DisablePhysic()
	pself.is_visible = false
	TweenPos(pself, Vec2{pself.GetPosX(), pself.GetPosY() - 500}, 2, func() {
		pself.Destroy()
	})
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
