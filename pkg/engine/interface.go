package engine

type ILifeCycle interface {
	OnStart()
	OnFixedUpdate(delta float32)
	OnUpdate(delta float32)
	OnDestroy()
}
type IManager interface {
	ILifeCycle
	Init(root Node)
}

type ISpriter interface {
	ILifeCycle
	onCreate()
	GetId() Object
	SetId(Object)
	Destroy() bool

	OnTriggerEnter(ISpriter)
	V_OnTriggerEnter(ISpriter)

	OnTriggerExit(ISpriter)
	V_OnTriggerExit(ISpriter)
}
type IUiNode interface {
	ILifeCycle
	onCreate()
	GetId() Object
	SetId(Object)
	Destroy() bool
	OnUiClick()
	V_OnUiClick()

	OnUiPressed()
	V_OnUiPressed()

	OnUiReleased()
	V_OnUiReleased()

	OnUiHovered()
	V_OnUiHovered()

	OnUiToggle(isOn bool)
	V_OnUiToggle(isOn bool)

	OnUiTextChanged(txt string)
	V_OnUiTextChanged(txt string)
}
type EngineCallbackInfo struct {
	OnEngineStart       func()
	OnEngineUpdate      func(float32)
	OnEngineFixedUpdate func(float32)
	OnEngineDestroy     func()
}

type CallbackInfo struct {
	EngineCallbackInfo

	// life cycle
	OnSpriteReady        func(int64)
	OnSpriteUpdated      func(float32)
	OnSpriteFixedUpdated func(float32)
	OnSpriteDestroyed    func(int64)

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

	OnSpriteScreenEntered     func(int64)
	OnSpriteScreenExited      func(int64)
	OnSpriteVfxFinished       func(int64)
	OnSpriteAnimationFinished func(int64)
	OnSpriteAnimationLooped   func(int64)
	OnSpriteFrameChanged      func(int64)
	OnSpriteAnimationChanged  func(int64)
	OnSpriteFramesSetChanged  func(int64)
}
