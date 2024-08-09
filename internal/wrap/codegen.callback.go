package wrap

import (
	. "godot-ext/gdspx/pkg/engine"
)

type CallbackInfo struct {
	EngineCallbackInfo

	// life cycle
	OnSpriteReady     func(int64)
	OnSpriteUpdated   func(int64)
	OnSpriteDestroyed func(int64)

	// input
	OnMousePressed       func(int64)
	OnMouseReleased      func(int64)
	OnKeyPressed         func(int64)
	OnKeyReleased        func(int64)
	OnActionPressed      func(string)
	OnActionJustPressed  func(string)
	OnActionJustReleased func(string)
	OnAxisChanged        func(string, float64)

	// physic
	OnCollisionEnter func(int64, int64)
	OnCollisionStay  func(int64, int64)
	OnCollisionExit  func(int64, int64)

	OnTriggerEnter func(int64, int64)
	OnTriggerStay  func(int64, int64)
	OnTriggerExit  func(int64, int64)

	// UI
	OnUIPressed     func(int64)
	OnUIReleased    func(int64)
	OnUIHovered     func(int64)
	OnUIClicked     func(int64)
	OnUIToggle      func(int64, bool)
	OnUITextChanged func(int64, string)
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
