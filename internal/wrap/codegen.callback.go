package wrap

import (
	. "godot-ext/gdspx/pkg/engine"
)

type spriteMgr struct {
	baseMgr
}
type uiMgr struct {
	baseMgr
}
type audioMgr struct {
	baseMgr
}
type inputMgr struct {
	baseMgr
}
type physicMgr struct {
	baseMgr
}

func createMgrs() []IManager {
	addManager(&audioMgr{})
	addManager(&uiMgr{})
	addManager(&physicMgr{})
	addManager(&inputMgr{})
	addManager(&spriteMgr{})
	return mgrs
}
