package wrap

import (
	. "godot-ext/gdspx/internal/core"
)

type CallbackInfo struct {
	EngineCallbackInfo

	// life cycle
	OnSpriteReady     func(int)
	OnSpriteUpdated   func(int)
	OnSpriteDestroyed func(int)

	// input
	OnMousePressed       func(int)
	OnMouseReleased      func(int)
	OnKeyPressed         func(int)
	OnKeyReleased        func(int)
	OnActionPressed      func(string)
	OnActionJustPressed  func(string)
	OnActionJustReleased func(string)
	OnAxisChanged        func(string, float64)

	// physic
	OnCollisionEnter func(int, int64)
	OnCollisionStay  func(int, int64)
	OnCollisionExit  func(int, int64)

	OnTriggerEnter func(int, int64)
	OnTriggerStay  func(int, int64)
	OnTriggerExit  func(int, int64)

	// UI
	OnUIPressed     func(int)
	OnUIReleased    func(int)
	OnUIHovered     func(int)
	OnUIClicked     func(int)
	OnUIToggle      func(int, bool)
	OnUITextChanged func(int, string)
}

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
