package manager

import (
	. "godot-ext/gdspx/internal/engine"
)

type baseMgr struct {
	Node Node
}

func (pself *baseMgr) Init(node Node) {
	pself.Node = node
	//println("init manager", node.GetName())
}

func (pself *baseMgr) Ready() {

}

func (pself *baseMgr) Process(delta float32) {
}
