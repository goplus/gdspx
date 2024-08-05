package engine

import (
	"grow.graphics/gd"
)

type BasicMgr struct {
	Root      gd.Node2D
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime
}
