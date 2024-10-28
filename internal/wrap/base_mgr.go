package wrap

import (
	"fmt"

	. "github.com/realdream-ai/gdspx/pkg/engine"
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
func (pself *baseMgr) OnUpdate(delta float32) {
}
func (pself *baseMgr) OnFixedUpdate(delta float32) {
}
func (pself *baseMgr) OnDestroy() {
}

func (mgr *baseMgr) logf(format string, v ...any) (n int, err error) {
	return fmt.Printf(format, v...)
}
func (mgr *baseMgr) log(a ...any) (n int, err error) {
	return fmt.Println(a...)
}
