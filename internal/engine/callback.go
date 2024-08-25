package engine

import (
	. "godot-ext/gdspx/pkg/engine"
)

func bindCallbacks() CallbackInfo {
	infos := CallbackInfo{}
	infos.OnEngineStart = onEngineStart
	infos.OnEngineUpdate = onEngineUpdate
	infos.OnEngineFixedUpdate = onEngineFixedUpdate
	infos.OnEngineDestroy = onEngineDestroy

	infos.OnSceneSpriteInstantiated = onSceneSpriteInstantiated

	infos.OnSpriteReady = onSpriteReady
	infos.OnSpriteUpdated = onSpriteUpdated
	infos.OnSpriteFixedUpdated = onSpriteFixedUpdated
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

	infos.OnSpriteScreenExited = onSpriteScreenExited
	infos.OnSpriteVfxFinished = onSpriteVfxFinished
	infos.OnSpriteAnimationFinished = onSpriteAnimationFinished
	infos.OnSpriteAnimationLooped = onSpriteAnimationLooped
	infos.OnSpriteFrameChanged = onSpriteFrameChanged
	infos.OnSpriteAnimationChanged = onSpriteAnimationChanged
	infos.OnSpriteFramesSetChanged = onSpriteFramesSetChanged

	return infos
}
func onSceneSpriteInstantiated(id int64, type_name string) {
	BindSceneInstantiatedSprite(Object(id), type_name)
}

// sprite
func onSpriteReady(id int64) {
	if sprite, ok := Id2Sprites[Object(id)]; ok {
		sprite.OnStart()
	}
}
func onSpriteUpdated(delta float32) {
	println("onSpriteUpdated ", delta)
}
func onSpriteFixedUpdated(delta float32) {
	println("onSpriteFixedUpdated ", delta)
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
	if sprite, ok := Id2Sprites[Object(id)]; ok {
		if other, ok2 := Id2Sprites[Object(oid)]; ok2 {
			sprite.V_OnTriggerEnter(other)
			sprite.OnTriggerEnter(other)
		}
	}
}
func onTriggerStay(id int64, oid int64) {
}
func onTriggerExit(id int64, oid int64) {
	if sprite, ok := Id2Sprites[Object(id)]; ok {
		if other, ok2 := Id2Sprites[Object(oid)]; ok2 {
			sprite.V_OnTriggerExit(other)
			sprite.OnTriggerExit(other)
		}
	}
}

// ui
func onUiPressed(id int64) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiPressed()
		node.OnUiPressed()
	}
}
func onUiReleased(id int64) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiReleased()
		node.OnUiReleased()
	}
}
func onUiHovered(id int64) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiHovered()
		node.OnUiHovered()
	}
}
func onUiClicked(id int64) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiClick()
		node.OnUiClick()
	}
}
func onUiToggle(id int64, isOn bool) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiToggle(isOn)
		node.OnUiToggle(isOn)
	}
}
func onUiTextChanged(id int64, text string) {
	if node, ok := Id2UiNodes[Object(id)]; ok {
		node.V_OnUiTextChanged(text)
		node.OnUiTextChanged(text)
	}
}

func onSpriteScreenExited(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteVfxFinished(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteAnimationFinished(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteAnimationLooped(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteFrameChanged(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteAnimationChanged(id int64) {
	println("onEngineFixedUpdate ", id)
}
func onSpriteFramesSetChanged(id int64) {
	println("onEngineFixedUpdate ", id)
}
