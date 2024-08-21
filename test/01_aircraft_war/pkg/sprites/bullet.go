package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
	. "gdspx-demo01/pkg/define"
)

type Bullet struct {
	Sprite
}

func (pself *Bullet) OnStart() {
	println("Bullet OnStart", pself.GetId())
}
func (pself *Bullet) OnUpdate(delta float32) {
	pos := pself.GetPosition()
	pos.Y += 2000 * delta
	pself.SetPosition(pos)
	if(pos.Y > WinHeight){
		println("Bullet OnDestory " , pself.GetId())
		pself.Destroy()
	}
}