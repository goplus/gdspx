package engine

type ILifeCycle interface {
	OnStart()
	OnUpdate(delta float32)
	OnDestroy()
}
type IManager interface {
	ILifeCycle
	Init(root Node)
}

type ISpriter interface {
	ILifeCycle
	SetId(Object)
	OnTriggerEnter(ISpriter)
	OnTriggerExit(ISpriter)
}
type IUiNode interface {
	ILifeCycle
	SetId(Object)
}
type EngineCallbackInfo struct {
	OnEngineStart   func()
	OnEngineUpdate  func(float32)
	OnEngineDestroy func()
}

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
	OnAxisChanged        func(string, float32)

	// physic
	OnCollisionEnter func(int64, int64)
	OnCollisionStay  func(int64, int64)
	OnCollisionExit  func(int64, int64)

	OnTriggerEnter func(int64, int64)
	OnTriggerStay  func(int64, int64)
	OnTriggerExit  func(int64, int64)

	// UI
	OnUiReady       func(int64)
	OnUiUpdated     func(int64)
	OnUiDestroyed   func(int64)
	OnUiPressed     func(int64)
	OnUiReleased    func(int64)
	OnUiHovered     func(int64)
	OnUiClicked     func(int64)
	OnUiToggle      func(int64, bool)
	OnUiTextChanged func(int64, string)
}
