package game

import (
	. "gdspx-demo02/pkg/sprite"

	. "github.com/realdream-ai/gdspx/pkg/engine"
)

var (
	timer = float32(0)
)

func RegisterTypes() {
	RegisterSpriteType[Player]()
	RegisterSpriteType[Enemy]()
	RegisterSpriteType[Bullet]()
	RegisterSpriteType[Goomba]()

}

func OnStart() {
	CameraMgr.SetCameraZoom(Vec2{2.75, 2.75})
	player := CreateSprite[Player]()
	player.SetPosition(Vec2{-175, 0})

}

func OnUpdate(delta float32) {
}

func OnDestroy() {

}
