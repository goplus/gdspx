package manager

import (
	. "github.com/godot-go/godot-go/pkg/builtin"
)

type baseMgr struct {
	Node Node
}

func (pself *baseMgr) AddChildNode(node Node, name string) Node {
	node.SetName_StrExt(name)
	//pself.Node.AddChild(node, true, NodeInternalMode(0))
	return node
}

func (pself *baseMgr) Init(node Node) {
	pself.Node = node
}

func (pself *baseMgr) Ready() {

}

func (pself *baseMgr) Process(delta float32) {
}
