package manager

import (
	"godot-ext/gd4go/internal/engine"
)

type SpriteMgr struct {
	engine.BasicMgr
}

func (pself *SpriteMgr) CreateSprite() {
	//mob, ok := gd.As[gd.RigidBody2D](Temporary, m.MobScene.Instantiate(m.Temporary, 0))
}
