package sprite

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Goomba struct {
	Enemy
	horizontalSpeed float32
	verticalSpeed float32
}

func (pself *Goomba) OnStart() {
	pself.horizontalSpeed = 20
	pself.verticalSpeed = 100
}

func (pself *Goomba) OnUpdate(delta float32) {
	pself.AddPosX(pself.horizontalSpeed * delta)
	if !pself.IsOnFloor(){
		pself.AddVelY(-pself.verticalSpeed * delta)
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
}


func (pself *Goomba) OnTriggerEnter(target ISpriter) {

}