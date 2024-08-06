package manager

import (
	"grow.graphics/gd"
)

type baseMgr struct {
	Root      gd.Node
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime
}

func (pself *baseMgr) Init(root gd.Node, keepAlive gd.Lifetime) {
	pself.Root = root
	pself.KeepAlive = keepAlive
}

func (pself *baseMgr) Ready() {

}

func (pself *baseMgr) Process(delta gd.Float) {
}
