package manager

import (
	"godot-ext/gd4go/internal/engine"
)

type spriteMgr struct {
	engine.BasicMgr
}

func (pself *spriteMgr) CreateSprite() {
	//mob, ok := gd.As[gd.RigidBody2D](Temporary, m.MobScene.Instantiate(m.Temporary, 0))
}
