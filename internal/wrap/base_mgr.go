package wrap

import (
	"fmt"
	. "godot-ext/gdspx/pkg/engine"
)

type baseMgr struct {
	Node Node
}

func (pself *baseMgr) Init(node Node) {
	pself.Node = node
	//println("init manager", node.GetName())
}

func (pself *baseMgr) OnStart() {
}
func (pself *baseMgr) OnUpdate(delta float64) {
}
func (pself *baseMgr) OnDestroy() {
}

func (mgr *baseMgr) logf(format string, v ...any) (n int, err error) {
	return fmt.Printf(format, v...)
}
func (mgr *baseMgr) log(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
