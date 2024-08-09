package wrap

import (
	"godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/pkg/engine"
)

func (pself *physicMgr) SetGravity(gravity float64) {
	pself.logf("Setting gravity to %f\n", gravity)
}

func (pself *physicMgr) GetGravity() float64 {
	return 0
}

func (pself *physicMgr) SetVelocity(id GdObject, velocity Vector2) {
	pself.logf("Setting velocity of %d to %f, %f\n", id, velocity.X, velocity.Y)
}

func (pself *physicMgr) GetVelocity(id GdObject) Vector2 {
	return Vector2{}
}

func (pself *physicMgr) SetMass(id GdObject, mass float64) {
	pself.logf("Setting mass of %d to %f\n", id, mass)
}

func (pself *physicMgr) GetMass(id GdObject) float64 {
	return 0
}

func (pself *physicMgr) AddForce(id GdObject, force Vector2) {
	pself.logf("Adding force to %d %f, %f\n", id, force.X, force.Y)
}

func (pself *physicMgr) AddImpulse(id GdObject, impulse Vector2) {
	pself.logf("Adding impulse to %d %f, %f\n", id, impulse.X, impulse.Y)
}

func (pself *physicMgr) SetCollisionLayer(id GdObject, layer GdObject) {
	pself.logf("Setting collision layer of %d to %d\n", id, layer)
}

func (pself *physicMgr) GetCollisionLayer(id GdObject) GdObject {
	return 0
}

func (pself *physicMgr) SetCollisionMask(id GdObject, mask GdObject) {
	pself.logf("Setting collision mask of %d to %d\n", id, mask)
}

func (pself *physicMgr) GetCollisionMask(id GdObject) GdObject {
	return 0
}

func (pself *physicMgr) GetColliderType(id GdObject) GdObject {
	return 0
}

func (pself *physicMgr) AddColliderRect(id GdObject, center Vector2, size Vector2) {
	pself.logf("Adding collider rect to %d with center %f, %f and size %f, %f\n", id, center.X, center.Y, size.X, size.Y)
}

func (pself *physicMgr) AddColliderCircle(id GdObject, center Vector2, radius float64) {
	pself.logf("Adding collider circle to %d with center %f, %f and radius %f\n", id, center.X, center.Y, radius)
}

func (pself *physicMgr) AddColliderCapsule(id GdObject, center Vector2, size Vector2) {
	pself.logf("Adding collider capsule to %d with center %f, %f and size %f, %f\n", id, center.X, center.Y, size.X, size.Y)
}

func (pself *physicMgr) SetTrigger(id GdObject, trigger bool) {
	pself.logf("Setting trigger of %d to %t\n", id, trigger)
}

func (pself *physicMgr) IsTrigger(id GdObject) bool {
	return false
}

func (pself *physicMgr) SetCollisionEnabled(id GdObject, enabled bool) {
	pself.logf("Setting collision enabled of %d to %t\n", id, enabled)
}

func (pself *physicMgr) IsCollisionEnabled(id GdObject) bool {
	return false
}

func (pself *spriteMgr) CreateSprite(path string) GdObject {
	cpath := ffi.NewCString(path)
	// TODO (jiepengtan): free cstring
	//defer cpath.Destroy()
	ffi.CallSpriteCreateSprite(cpath.ToGdString())
	return 0
}

func (pself *spriteMgr) CloneSprite(id GdObject) GdObject {
	return 0
}

func (pself *spriteMgr) DestroySprite(id GdObject) bool {
	return false
}

func (pself *spriteMgr) IsSpriteAlive(id GdObject) bool {
	return false
}

func (pself *spriteMgr) SetPosition(id GdObject, pos Vector2) {
	pself.logf("Updating position of %d to %f, %f", id, pos.X, pos.Y)
}

func (pself *spriteMgr) SetRotation(id GdObject, rot Vector2) {
	pself.logf("Updating rotation of %d to %v", id, rot)
}

func (pself *spriteMgr) SetScale(id GdObject, scale Vector2) {
	pself.logf("Updating scale of %d to %f, %f", id, scale.X, scale.Y)
}

func (pself *spriteMgr) GetPosition(id GdObject) Vector2 {
	return Vector2{}
}

func (pself *spriteMgr) GetRotation(id GdObject) Vector2 {
	return Vector2{}
}

func (pself *spriteMgr) GetScale(id GdObject) Vector2 {
	return Vector2{}
}

func (pself *spriteMgr) SetColor(id GdObject, color Color) {
	pself.logf("Updating color of %d to (%s)", id, color.R, color.G, color.B, color.A)
}

func (pself *spriteMgr) GetColor(id GdObject) Color {
	return Color{}
}

func (pself *spriteMgr) UpdateTexture(id GdObject, path string) {
	pself.logf("Updating texture of %d to %s", id, path)
}

func (pself *spriteMgr) GetTexture(id GdObject) string {
	return ""
}

func (pself *spriteMgr) SetVisible(id GdObject, visible bool) {
	pself.logf("Updating visibility of %d to %v", id, visible)
}

func (pself *spriteMgr) GetVisible(id GdObject) bool {
	return false
}

func (pself *spriteMgr) UpdateZIndex(id GdObject, z GdObject) {
	pself.logf("Updating z index of %d to %d", id, z)
}

func (pself *uiMgr) CreateButton(path string, rect Rect2, text string) GdObject   { return 0 }
func (pself *uiMgr) CreateLabel(path string, rect Rect2, text string) GdObject    { return 0 }
func (pself *uiMgr) CreateImage(path string, rect Rect2, color Rect2) GdObject    { return 0 }
func (pself *uiMgr) CreateSlider(path string, rect Rect2, value float64) GdObject { return 0 }
func (pself *uiMgr) CreateToggle(path string, rect Rect2, value bool) GdObject    { return 0 }
func (pself *uiMgr) CreateInput(path string, rect Rect2, text string) GdObject    { return 0 }

func (pself *uiMgr) GetType(id GdObject) GdObject                        { return 0 }
func (pself *uiMgr) SetInteractable(id GdObject, GdObjecteractable bool) {}
func (pself *uiMgr) GetInteractable(id GdObject) bool                    { return false }

func (pself *uiMgr) SetText(id GdObject, text string)      {}
func (pself *uiMgr) GetText(id GdObject) string            { return "" }
func (pself *uiMgr) SetRect(id GdObject, rect Rect2)       {}
func (pself *uiMgr) GetRect(id GdObject) Rect2             { return Rect2{} }
func (pself *uiMgr) SetColor(id GdObject, color Rect2)     {}
func (pself *uiMgr) GetColor(id GdObject) Rect2            { return Rect2{} }
func (pself *uiMgr) SetFontSize(id GdObject, size float64) {}
func (pself *uiMgr) GetFontSize(id GdObject) float64       { return 0 }
func (pself *uiMgr) SetVisible(id GdObject, visible bool)  {}
func (pself *uiMgr) GetVisible(id GdObject) bool           { return false }

func (pself *audioMgr) PlayAudio(path string) {
	pself.logf("Playing %v", path)
}

func (pself *audioMgr) SetAudioVolume(volume float64) {
	pself.logf("Setting volume to %v", volume)
}

func (pself *audioMgr) GetAudioVolume() {
	pself.log("GetAudioVolume")
}

func (pself *audioMgr) IsMusicPlaying() bool {
	return false
}

func (pself *audioMgr) PlayMusic(path string) {
	pself.logf("Playing music %v", path)
}

func (pself *audioMgr) SetMusicVolume(volume float64) {
	pself.logf("Setting music volume to %v", volume)
}

func (pself *audioMgr) GetMusicVolume() {
	pself.log("GetMusicVolume")
}

func (pself *audioMgr) PauseMusic() {
	pself.log("PauseMusic")
}

func (pself *audioMgr) ResumeMusic() {
	pself.log("ResumeMusic")
}

func (pself *audioMgr) GetMusicTimer() float64 {
	return 0
}

func (pself *audioMgr) SetMusicTimer(time float64) {
	pself.logf("Setting music timer to %v", time)
}

func (pself *inputMgr) GetMousePos() Vector2 {
	return Vector2{}
}

func (pself *inputMgr) GetMouseState(id int64) bool {
	return false
}

func (pself *inputMgr) GetKeyState(key int64) int64 {
	return 0
}

func (pself *inputMgr) GetAxis(axis string) float64 {
	return 0
}

func (pself *inputMgr) IsActionPressed(action string) bool {
	return false
}

func (pself *inputMgr) IsActionJustPressed(action string) bool {
	return false
}

func (pself *inputMgr) IsActionJustReleased(action string) bool {
	return false
}
