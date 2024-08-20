package game

import (
	. "gdspx-demo01/pkg/sprites"
	. "godot-ext/gdspx/pkg/engine"
)

var (
	WinWidth  = float32(1080)
	WinHeight = float32(1920)
)

func OnStart() {
	obj := CreateSprite[Aircraft]()
	obj.SetPosition(Vec2{0, -WinHeight / 2.0})
	obj2 := CreateSprite[Bullet]()
	obj2.SetPosition(Vec2{0, -WinHeight/2.0 + 100})
	obj3 := CreateSprite[Enemy]()
	obj3.SetPosition(Vec2{0, WinHeight / 2.0})
}

func OnUpdate(delta float32) {
	for _, sprite := range Sprites {
		sprite.OnUpdate(delta)
	}
}

func OnDestroy() {
	for _, sprite := range Sprites {
		sprite.OnDestroy()
	}
}
