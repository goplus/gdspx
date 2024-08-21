package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Aircraft struct {
	Sprite
	timer float32
	moveSpeed float32
}

func (pself *Aircraft) OnStart() {
	pself.moveSpeed = 300
}
func (pself *Aircraft) OnUpdate(delta float32) {
	pself.timer+= delta
	if pself.timer > 0.3 {
		pself.timer = 0
		bullet := CreateSprite[Bullet]()
		pos := pself.GetPosition()
		bullet.SetPosition(Vec2{pos.X , pos.Y + 150})
	}
	
	if InputMgr.GetKeyState(KeyCode.W) == 1 {
		pself.Move(0, pself.moveSpeed * delta)
	}
	if InputMgr.GetKeyState(KeyCode.S) == 1 {
		pself.Move(0, -pself.moveSpeed * delta)
	}
	if InputMgr.GetKeyState(KeyCode.D) == 1 {
		pself.Move( pself.moveSpeed * delta,0)
	}
	if InputMgr.GetKeyState(KeyCode.A) == 1 {
		pself.Move(-pself.moveSpeed * delta,0)
	}
}


func (pself *Aircraft) OnDestory() {
	println("Aircraft OnDestory")
}
