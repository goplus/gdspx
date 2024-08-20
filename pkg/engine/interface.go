package engine

type IManager interface {
	Init(root Node)
	OnStart()
	OnUpdate(delta float32)
	OnDestroy()
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
	OnUIPressed     func(int64)
	OnUIReleased    func(int64)
	OnUIHovered     func(int64)
	OnUIClicked     func(int64)
	OnUIToggle      func(int64, bool)
	OnUITextChanged func(int64, string)
}
