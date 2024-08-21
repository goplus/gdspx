package game

import (
	. "gdspx-demo01/pkg/sprites"
	. "godot-ext/gdspx/pkg/engine"
	"math/rand"
)

var (
	WinWidth  = float32(1080)
	WinHeight = float32(1920)
	timer	 = float32(0)
)


func onGameStart() {
	obj := CreateSprite[Aircraft]()
	obj.SetPosition(Vec2{0, -WinHeight / 2.0})
}

func onGameUpdate(delta float32) {
	timer+= delta
	if timer > 1 {
		timer = 0
		obj := CreateSprite[Enemy]()
		value := (rand.Float32() *2 -1) * WinWidth / 2.0
		obj.SetPosition(Vec2{value , WinHeight / 2.0})
	}
}

func onGameDestroy() {

}