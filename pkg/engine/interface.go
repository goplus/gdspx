package engine

var (
	AudioMgr  IAudioMgr
	UIMgr     IUIMgr
	PhysicMgr IPhysicMgr
	InputMgr  IInputMgr
	SpriteMgr ISpriteMgr
)

type IManager interface {
	Init(root Node)
	OnStart()
	OnUpdate(delta float64)
	OnDestroy()
}

type EngineCallbackInfo struct {
	OnEngineStart   func()
	OnEngineUpdate  func(float64)
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

type IAudioMgr interface {
	PlayAudio(path string)
	SetAudioVolume(volume float64)
	GetAudioVolume() float64
	IsMusicPlaying() bool
	PlayMusic(path string)
	SetMusicVolume(volume float64)
	GetMusicVolume() float64
	PauseMusic()
	ResumeMusic()
	GetMusicTimer() float64
	SetMusicTimer(time float64)
}

type IUIMgr interface {
	CreateButton(path string, rect Rect2, text string) Object
	CreateLabel(path string, rect Rect2, text string) Object
	CreateImage(path string, rect Rect2, color Color) Object
	CreateSlider(path string, rect Rect2, value float64) Object
	CreateToggle(path string, rect Rect2, value bool) Object
	CreateInput(path string, rect Rect2, text string) Object
	GetType(obj Object) int64
	SetInteractable(obj Object, interactable bool)
	GetInteractable(obj Object) bool
	SetText(obj Object, text string)
	GetText(obj Object) string
	SetRect(obj Object, rect Rect2)
	GetRect(obj Object) Rect2
	SetColor(obj Object, color Color)
	GetColor(obj Object) Color
	SetFontSize(obj Object, size float64)
	GetFontSize(obj Object) float64
	SetVisible(obj Object, visible bool)
	GetVisible(obj Object) bool
}

type IPhysicMgr interface {
	SetGravity(gravity float64)
	GetGravity() float64
	SetVelocity(obj Object, velocity Vec2)
	GetVelocity(obj Object) Vec2
	SetMass(obj Object, mass float64)
	GetMass(obj Object) float64
	AddForce(obj Object, force Vec2)
	AddImpulse(obj Object, impulse Vec2)
	SetCollisionLayer(obj Object, layer int64)
	GetCollisionLayer(obj Object) int64
	SetCollisionMask(obj Object, mask int64)
	GetCollisionMask(obj Object) int64
	GetColliderType(obj Object) int64
	AddColliderRect(obj Object, center Vec2, size Vec2)
	AddColliderCircle(obj Object, center Vec2, radius float64)
	AddColliderCapsule(obj Object, center Vec2, size Vec2)
	SetTrigger(obj Object, trigger bool)
	IsTrigger(obj Object) bool
	SetCollisionEnabled(obj Object, enabled bool)
	IsCollisionEnabled(obj Object) bool
}

type IInputMgr interface {
	GetMousePos() Vec2
	GetMouseState(obj int64) bool
	GetKeyState(key int64) int64
	GetAxis(axis string) float64
	IsActionPressed(action string) bool
	IsActionJustPressed(action string) bool
	IsActionJustReleased(action string) bool
}

type ISpriteMgr interface {
	CreateSprite(path string) Object
	CloneSprite(obj Object) Object
	DestroySprite(obj Object) bool
	IsSpriteAlive(obj Object) bool
	SetPosition(obj Object, pos Vec2)
	SetRotation(obj Object, rot float64)
	SetScale(obj Object, scale Vec2)
	GetPosition(obj Object) Vec2
	GetRotation(obj Object) float64
	GetScale(obj Object) Vec2
	SetColor(obj Object, color Color)
	GetColor(obj Object) Color
	UpdateTexture(obj Object, path string)
	GetTexture(obj Object) string
	SetVisible(obj Object, visible bool)
	GetVisible(obj Object) bool
	UpdateZIndex(obj Object, z int64)

	// TODO animation
}
