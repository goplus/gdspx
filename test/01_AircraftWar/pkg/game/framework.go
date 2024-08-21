package game

import (
	. "gdspx-demo01/pkg/sprites"
)

var (
	sprites   = make([]ISpriter, 0)
)

func OnStart() {
	onGameStart()
}

func OnUpdate(delta float32) {
	onGameUpdate(delta)
	sprites = sprites[:0]
	for _, sprite := range Id2Sprites {
		sprites = append(sprites, sprite)
	}
	for _, sprite := range sprites {
		sprite.OnUpdate(delta)
	}
}

func OnDestroy() {
	onGameDestroy()
	sprites = sprites[:0]
	for _, sprite := range Id2Sprites {
		sprites = append(sprites, sprite)
	}
	for _, sprite := range sprites {
		sprite.OnDestroy()
	}
}
