package game

import (
	. "gdspx-demo02/pkg/sprite"
	. "godot-ext/gdspx/pkg/engine"
)

var (
	timer        = float32(0)
)

func OnStart() {
	player := CreateSprite[Player]()
	player.SetPosition(Vec2{-175, 62})
}

func OnUpdate(delta float32) {
}

func OnDestroy() {

}
