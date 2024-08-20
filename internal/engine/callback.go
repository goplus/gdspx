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
	infos.OnUIPressed = onUIPressed
	infos.OnUIReleased = onUIReleased
	infos.OnUIHovered = onUIHovered
	infos.OnUIClicked = onUIClicked
	infos.OnUIToggle = onUIToggle
	infos.OnUITextChanged = onUITextChanged

	return infos
}

func onSpriteReady(id int64)     {}
func onSpriteUpdated(id int64)   {}
func onSpriteDestroyed(id int64) {}

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

func onTriggerEnter(id int64, oid int64) {}
func onTriggerStay(id int64, oid int64)  {}
func onTriggerExit(id int64, oid int64)  {}

// UI
func onUIPressed(id int64)                  {}
func onUIReleased(id int64)                 {}
func onUIHovered(id int64)                  {}
func onUIClicked(id int64)                  {}
func onUIToggle(id int64, isOn bool)        {}
func onUITextChanged(id int64, text string) {}
