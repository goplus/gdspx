package engine

import (
	. "godot-ext/gdspx/pkg/engine"
)

func bindCallbacks() CallbackInfo {
	infos := CallbackInfo{}
	infos.OnEngineStart = onEngineStart
	infos.OnEngineUpdate = onEngineUpdate
	infos.OnEngineDestroy = onEngineDestroy
	infos.OnSpriteReady = onSpriteReady
	infos.OnSpriteUpdated = onSpriteUpdated
	infos.OnSpriteDestroyed = onSpriteDestroyed

	// input
	infos.OnMousePressed = onMousePressed
	infos.OnMouseReleased = onMouseReleased
	infos.OnKeyPressed = onKeyPressed
	infos.OnKeyReleased = onKeyReleased
	infos.OnActionPressed = onActionPressed
	infos.OnActionJustPressed = onActionJustPressed
	infos.OnActionJustReleased = onActionJustReleased
	infos.OnAxisChanged = onAxisChanged

	// physic
	infos.OnCollisionEnter = onCollisionEnter
	infos.OnCollisionStay = onCollisionStay
	infos.OnCollisionExit = onCollisionExit

	infos.OnTriggerEnter = onTriggerEnter
	infos.OnTriggerStay = onTriggerStay
	infos.OnTriggerExit = onTriggerExit

	// ui
	infos.OnUiPressed = onUiPressed
	infos.OnUiReleased = onUiReleased
	infos.OnUiHovered = onUiHovered
	infos.OnUiClicked = onUiClicked
	infos.OnUiToggle = onUiToggle
	infos.OnUiTextChanged = onUiTextChanged

	return infos
}

func onSpriteReady(id int64) {
	println("onSpriteReady ", id)
}
func onSpriteUpdated(id int64) {
	println("onSpriteUpdated ", id)
}
func onSpriteDestroyed(id int64) {
	println("onSpriteDestroyed ", id)
}

// input
func onMousePressed(id int64)                  {}
func onMouseReleased(id int64)                 {}
func onKeyPressed(id int64)                    {}
func onKeyReleased(id int64)                   {}
func onActionPressed(name string)              {}
func onActionJustPressed(name string)          {}
func onActionJustReleased(name string)         {}
func onAxisChanged(name string, value float32) {}

// physic
func onCollisionEnter(id int64, oid int64) {}
func onCollisionStay(id int64, oid int64)  {}
func onCollisionExit(id int64, oid int64)  {}

func onTriggerEnter(id int64, oid int64) {
	println("onTriggerEnter ", id, oid)
}
func onTriggerStay(id int64, oid int64) {}
func onTriggerExit(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}

// UI
func onUiPressed(id int64)           {}
func onUiReleased(id int64)          {}
func onUiHovered(id int64)           {}
func onUiClicked(id int64)           {}
func onUiToggle(id int64, isOn bool) {}
func onUiTextChanged(id int64, text string) {

}
