package engine

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
	CreateButton(path string, rect Rect2, text string) int64
	CreateLabel(path string, rect Rect2, text string) int64
	CreateImage(path string, rect Rect2, color Rect2) int64
	CreateSlider(path string, rect Rect2, value float64) int64
	CreateToggle(path string, rect Rect2, value bool) int64
	CreateInput(path string, rect Rect2, text string) int64
	GetType(id int64) int64
	SetInteractable(id int64, int64eractable bool)
	GetInteractable(id int64) bool
	SetText(id int64, text string)
	GetText(id int64) string
	SetRect(id int64, rect Rect2)
	GetRect(id int64) Rect2
	SetColor(id int64, color Rect2)
	GetColor(id int64) Rect2
	SetFontSize(id int64, size float64)
	GetFontSize(id int64) float64
	SetVisible(id int64, visible bool)
	GetVisible(id int64) bool
}
type IPhysicMgr interface {
	SetGravity(gravity float64)
	GetGravity() float64
	SetVelocity(id int64, velocity Vector2)
	GetVelocity(id int64) Vector2
	SetMass(id int64, mass float64)
	GetMass(id int64) float64
	AddForce(id int64, force Vector2)
	AddImpulse(id int64, impulse Vector2)
	SetCollisionLayer(id int64, layer int64)
	GetCollisionLayer(id int64) int64
	SetCollisionMask(id int64, mask int64)
	GetCollisionMask(id int64) int64
	GetColliderType(id int64) int64
	AddColliderRect(id int64, center Vector2, size Vector2)
	AddColliderCircle(id int64, center Vector2, radius float64)
	AddColliderCapsule(id int64, center Vector2, size Vector2)
	SetTrigger(id int64, trigger bool)
	IsTrigger(id int64) bool
	SetCollisionEnabled(id int64, enabled bool)
	IsCollisionEnabled(id int64) bool
}

type IInputMgr interface {
	GetMousePos() Vector2
	GetMouseState(id int64) bool
	GetKeyState(key int64) int64
	GetAxis(axis string) float64
	IsActionPressed(action string) bool
	IsActionJustPressed(action string) bool
	IsActionJustReleased(action string) bool
}

type ISpriteMgr interface {
	CreateSprite(path string) int64
	CloneSprite(id int64) int64
	DestroySprite(id int64) bool
	IsSpriteAlive(id int64) bool
	SetPosition(id int64, pos Vector2)
	SetRotation(id int64, rot Vector2)
	SetScale(id int64, scale Vector2)
	GetPosition(id int64) Vector2
	GetRotation(id int64) Vector2
	GetScale(id int64) Vector2
	SetColor(id int64, color Color)
	GetColor(id int64) Color
	UpdateTexture(id int64, path string)
	GetTexture(id int64) string
	SetVisible(id int64, visible bool)
	GetVisible(id int64) bool
	UpdateZIndex(id int64, z int64)
}
