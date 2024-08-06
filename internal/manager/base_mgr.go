package manager

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

type baseMgr struct {
	Node Node
}

func (pself *baseMgr) Init(node Node) {
	pself.Node = node
	println("init manager", node.GetName().ToGoString())
}

func (pself *baseMgr) Ready() {

}

func (pself *baseMgr) Process(delta float32) {
}
