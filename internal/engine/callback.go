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

// sprite
func onSpriteReady(id int64) {
	if sprite, ok := Id2Sprites[Object(id)]; ok {
		sprite.OnStart()
	}
}
func onSpriteUpdated(id int64) {
	println("onSpriteUpdated ", id)
}
func onSpriteDestroyed(id int64) {
	delete(Id2Sprites, Object(id))
}

// input
func onMousePressed(id int64) {
	println("onMousePressed ", id)
}
func onMouseReleased(id int64) {
	println("onMouseReleased ", id)
}
func onKeyPressed(id int64) {
	println("onKeyPressed ", id)
}
func onKeyReleased(id int64) {
	println("onKeyReleased ", id)
}
func onActionPressed(name string) {
	println("onActionPressed ", name)
}
func onActionJustPressed(name string) {
	println("onActionJustPressed ", name)
}
func onActionJustReleased(name string) {
	println("onActionJustReleased ", name)
}
func onAxisChanged(name string, value float32) {
	println("onAxisChanged ", name, value)
}

// physic
func onCollisionEnter(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}
func onCollisionStay(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}
func onCollisionExit(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}

func onTriggerEnter(id int64, oid int64) {
	println("onTriggerEnter ", id, oid)
}
func onTriggerStay(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}
func onTriggerExit(id int64, oid int64) {
	println("onTriggerExit ", id, oid)
}

// ui
func onUiPressed(id int64) {
	println("onUiPressed ", id)
}
func onUiReleased(id int64) {
	println("onUiReleased ", id)
}
func onUiHovered(id int64) {
	println("onUiHovered ", id)
}
func onUiClicked(id int64) {
	println("onUiClicked ", id)
}
func onUiToggle(id int64, isOn bool) {
	println("onUiToggle ", id)
}
func onUiTextChanged(id int64, text string) {
	println("onUiTextChanged ", id, text)

}
