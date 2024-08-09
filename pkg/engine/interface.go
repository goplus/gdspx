package engine

var (
	AudioMgr  IAudioMgr
	UIMgr     IUIMgr
	PhysicMgr IPhysicMgr
	InputMgr  IInputMgr
	SpriteMgr ISpriteMgr
)


type IManager interface {
	Init(root GdNode)
	OnStart()
	OnUpdate(delta float64)
	OnDestroy()
}
type EngineCallbackInfo struct {
	OnEngineStart   func()
	OnEngineUpdate  func(float64)
	OnEngineDestroy func()
}

type IAudioMgr interface {
	PlayAudio(path string)
	SetAudioVolume(volume float64)
	GetAudioVolume()
	IsMusicPlaying() bool
	PlayMusic(path string)
	SetMusicVolume(volume float64)
	GetMusicVolume()
	PauseMusic()
	ResumeMusic()
	GetMusicTimer() float64
	SetMusicTimer(time float64)
}

type IUIMgr interface {
	CreateButton(path string, rect Rect2, text string) GdObject
	CreateLabel(path string, rect Rect2, text string) GdObject
	CreateImage(path string, rect Rect2, color Rect2) GdObject
	CreateSlider(path string, rect Rect2, value float64) GdObject
	CreateToggle(path string, rect Rect2, value bool) GdObject
	CreateInput(path string, rect Rect2, text string) GdObject
	GetType(obj GdObject) GdObject
	SetInteractable(obj GdObject, interactable bool)
	GetInteractable(obj GdObject) bool
	SetText(obj GdObject, text string)
	GetText(obj GdObject) string
	SetRect(obj GdObject, rect Rect2)
	GetRect(obj GdObject) Rect2
	SetColor(obj GdObject, color Rect2)
	GetColor(obj GdObject) Rect2
	SetFontSize(obj GdObject, size float64)
	GetFontSize(obj GdObject) float64
	SetVisible(obj GdObject, visible bool)
	GetVisible(obj GdObject) bool
}
type IPhysicMgr interface {
	SetGravity(gravity float64)
	GetGravity() float64
	SetVelocity(obj GdObject, velocity Vector2)
	GetVelocity(obj GdObject) Vector2
	SetMass(obj GdObject, mass float64)
	GetMass(obj GdObject) float64
	AddForce(obj GdObject, force Vector2)
	AddImpulse(obj GdObject, impulse Vector2)
	SetCollisionLayer(obj GdObject, layer GdObject)
	GetCollisionLayer(obj GdObject) GdObject
	SetCollisionMask(obj GdObject, mask GdObject)
	GetCollisionMask(obj GdObject) GdObject
	GetColliderType(obj GdObject) GdObject
	AddColliderRect(obj GdObject, center Vector2, size Vector2)
	AddColliderCircle(obj GdObject, center Vector2, radius float64)
	AddColliderCapsule(obj GdObject, center Vector2, size Vector2)
	SetTrigger(obj GdObject, trigger bool)
	IsTrigger(obj GdObject) bool
	SetCollisionEnabled(obj GdObject, enabled bool)
	IsCollisionEnabled(obj GdObject) bool
}

type IInputMgr interface {
	GetMousePos() Vector2
	GetMouseState(obj int64) bool
	GetKeyState(key int64) int64
	GetAxis(axis string) float64
	IsActionPressed(action string) bool
	IsActionJustPressed(action string) bool
	IsActionJustReleased(action string) bool
}

type ISpriteMgr interface {
	CreateSprite(path string) GdObject
	CloneSprite(obj GdObject) GdObject
	DestroySprite(obj GdObject) bool
	IsSpriteAlive(obj GdObject) bool
	SetPosition(obj GdObject, pos Vector2)
	SetRotation(obj GdObject, rot Vector2)
	SetScale(obj GdObject, scale Vector2)
	GetPosition(obj GdObject) Vector2
	GetRotation(obj GdObject) Vector2
	GetScale(obj GdObject) Vector2
	SetColor(obj GdObject, color Color)
	GetColor(obj GdObject) Color
	UpdateTexture(obj GdObject, path string)
	GetTexture(obj GdObject) string
	SetVisible(obj GdObject, visible bool)
	GetVisible(obj GdObject) bool
	UpdateZIndex(obj GdObject, z GdObject)
}
