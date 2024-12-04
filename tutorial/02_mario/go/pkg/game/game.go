package game

import (
	. "gdspx-demo02/pkg/sprite"

	. "github.com/realdream-ai/gdspx/pkg/engine"
	. "github.com/realdream-ai/mathf"
)

var (
	timer = float64(0)
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

func OnUpdate(delta float64) {
}

func OnDestroy() {

}
