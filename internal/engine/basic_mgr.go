package engine

import (
	"grow.graphics/gd"
)

type IManager interface {
	Init(root gd.Node, keepAlive gd.Lifetime)
	Ready()
	Process(delta gd.Float)
}

type BasicMgr struct {
	Root      gd.Node
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime
}

func (pself *BasicMgr) Init(root gd.Node, keepAlive gd.Lifetime) {
	pself.Root = root
	pself.KeepAlive = keepAlive
}

func (pself *BasicMgr) Ready() {

}

func (pself *BasicMgr) Process(delta gd.Float) {
}
