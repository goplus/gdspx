/*------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
package wrap

import (
	"fmt"
	"reflect"
	. "godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/pkg/engine"
)
func BindMgr(mgrs []IManager) {
	for _, mgr := range mgrs {
		switch v := mgr.(type) {
		case IAudioMgr:
			AudioMgr = v

		case IInputMgr:
			InputMgr = v

		case IPhysicMgr:
			PhysicMgr = v

		case ISpriteMgr:
			SpriteMgr = v

		case IUiMgr:
			UiMgr = v

		default:
			panic(fmt.Sprintf("engine init error : unknown manager type %s", reflect.TypeOf(mgr).String()))
		}
	}
}

type audioMgr struct {
	baseMgr
}
type inputMgr struct {
	baseMgr
}
type physicMgr struct {
	baseMgr
}
type spriteMgr struct {
	baseMgr
}
type uiMgr struct {
	baseMgr
}


func createMgrs() []IManager {
	addManager(&audioMgr{})
	addManager(&inputMgr{})
	addManager(&physicMgr{})
	addManager(&spriteMgr{})
	addManager(&uiMgr{})
	return mgrs
}

// call gdextension interface functions

func (pself *audioMgr) PlayAudio(path string) {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	CallAudioPlayAudio(arg0)
}
func (pself *audioMgr) SetAudioVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	CallAudioSetAudioVolume(arg0)
}
func (pself *audioMgr) GetAudioVolume() float32  {
	retValue := CallAudioGetAudioVolume()
	return ToFloat32(retValue)
}
func (pself *audioMgr) IsMusicPlaying() bool  {
	retValue := CallAudioIsMusicPlaying()
	return ToBool(retValue)
}
func (pself *audioMgr) PlayMusic(path string) {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	CallAudioPlayMusic(arg0)
}
func (pself *audioMgr) SetMusicVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	CallAudioSetMusicVolume(arg0)
}
func (pself *audioMgr) GetMusicVolume() float32  {
	retValue := CallAudioGetMusicVolume()
	return ToFloat32(retValue)
}
func (pself *audioMgr) PauseMusic() {
	CallAudioPauseMusic()
}
func (pself *audioMgr) ResumeMusic() {
	CallAudioResumeMusic()
}
func (pself *audioMgr) GetMusicTimer() float32  {
	retValue := CallAudioGetMusicTimer()
	return ToFloat32(retValue)
}
func (pself *audioMgr) SetMusicTimer(time float32) {
	arg0 := ToGdFloat(time)
	CallAudioSetMusicTimer(arg0)
}
func (pself *inputMgr) GetMousePos() Vec2  {
	retValue := CallInputGetMousePos()
	return ToVec2(retValue)
}
func (pself *inputMgr) GetMouseState(mouse_id int64) bool  {
	arg0 := ToGdInt(mouse_id)
	retValue := CallInputGetMouseState(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) GetKeyState(key int64) int64  {
	arg0 := ToGdInt(key)
	retValue := CallInputGetKeyState(arg0)
	return ToInt64(retValue)
}
func (pself *inputMgr) GetAxis(axis string) float32  {
	arg0Str := NewCString(axis)
	arg0 := arg0Str.ToGdString() 
	retValue := CallInputGetAxis(arg0)
	return ToFloat32(retValue)
}
func (pself *inputMgr) IsActionPressed(action string) bool  {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString() 
	retValue := CallInputIsActionPressed(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) IsActionJustPressed(action string) bool  {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString() 
	retValue := CallInputIsActionJustPressed(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) IsActionJustReleased(action string) bool  {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString() 
	retValue := CallInputIsActionJustReleased(arg0)
	return ToBool(retValue)
}
func (pself *physicMgr) Raycast(from Vec2, to Vec2, collision_mask int64) Object  {
	arg0 := ToGdVec2(from)
	arg1 := ToGdVec2(to)
	arg2 := ToGdInt(collision_mask)
	retValue := CallPhysicRaycast(arg0, arg1, arg2)
	return ToObject(retValue)
}
func (pself *spriteMgr) CreateSprite(path string) Object  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	retValue := CallSpriteCreateSprite(arg0)
	return ToObject(retValue)
}
func (pself *spriteMgr) CloneSprite(obj Object) Object  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteCloneSprite(arg0)
	return ToObject(retValue)
}
func (pself *spriteMgr) DestroySprite(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteDestroySprite(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsSpriteAlive(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsSpriteAlive(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetPosition(obj Object, pos Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(pos)
	CallSpriteSetPosition(arg0, arg1)
}
func (pself *spriteMgr) SetRotation(obj Object, rot float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(rot)
	CallSpriteSetRotation(arg0, arg1)
}
func (pself *spriteMgr) SetScale(obj Object, scale Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(scale)
	CallSpriteSetScale(arg0, arg1)
}
func (pself *spriteMgr) GetPosition(obj Object) Vec2  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetPosition(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) GetRotation(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetRotation(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) GetScale(obj Object) Vec2  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetScale(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	CallSpriteSetColor(arg0, arg1)
}
func (pself *spriteMgr) GetColor(obj Object) Color  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetColor(arg0)
	return ToColor(retValue)
}
func (pself *spriteMgr) SetTexture(obj Object, path string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString() 
	CallSpriteSetTexture(arg0, arg1)
}
func (pself *spriteMgr) GetTexture(obj Object) string  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetTexture(arg0)
	return ToString(retValue)
}
func (pself *spriteMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	CallSpriteSetVisible(arg0, arg1)
}
func (pself *spriteMgr) GetVisible(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetVisible(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) GetZIndex(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetZIndex(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetZIndex(obj Object, z int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(z)
	CallSpriteSetZIndex(arg0, arg1)
}
func (pself *spriteMgr) PlayAnim(obj Object, p_name string, p_custom_scale float32, p_from_end bool) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(p_name)
	arg1 := arg1Str.ToGdString() 
	arg2 := ToGdFloat(p_custom_scale)
	arg3 := ToGdBool(p_from_end)
	CallSpritePlayAnim(arg0, arg1, arg2, arg3)
}
func (pself *spriteMgr) PlayBackwardsAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(p_name)
	arg1 := arg1Str.ToGdString() 
	CallSpritePlayBackwardsAnim(arg0, arg1)
}
func (pself *spriteMgr) PauseAnim(obj Object) {
	arg0 := ToGdObj(obj)
	CallSpritePauseAnim(arg0)
}
func (pself *spriteMgr) StopAnim(obj Object) {
	arg0 := ToGdObj(obj)
	CallSpriteStopAnim(arg0)
}
func (pself *spriteMgr) IsPlayingAnim(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsPlayingAnim(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(p_name)
	arg1 := arg1Str.ToGdString() 
	CallSpriteSetAnim(arg0, arg1)
}
func (pself *spriteMgr) GetAnim(obj Object) string  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnim(arg0)
	return ToString(retValue)
}
func (pself *spriteMgr) SetAnimFrame(obj Object, p_frame int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(p_frame)
	CallSpriteSetAnimFrame(arg0, arg1)
}
func (pself *spriteMgr) GetAnimFrame(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimFrame(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetAnimSpeedScale(obj Object, p_speed_scale float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(p_speed_scale)
	CallSpriteSetAnimSpeedScale(arg0, arg1)
}
func (pself *spriteMgr) GetAnimSpeedScale(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimSpeedScale(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) GetAnimPlayingSpeed(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimPlayingSpeed(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetAnimCentered(obj Object, p_center bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_center)
	CallSpriteSetAnimCentered(arg0, arg1)
}
func (pself *spriteMgr) IsAnimCentered(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimCentered(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnimOffset(obj Object, p_offset Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(p_offset)
	CallSpriteSetAnimOffset(arg0, arg1)
}
func (pself *spriteMgr) GetAnimOffset(obj Object) Vec2  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimOffset(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetAnimFlipH(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	CallSpriteSetAnimFlipH(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedH(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimFlippedH(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnimFlipV(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	CallSpriteSetAnimFlipV(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedV(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimFlippedV(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetGravity(obj Object, gravity float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(gravity)
	CallSpriteSetGravity(arg0, arg1)
}
func (pself *spriteMgr) GetGravity(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetGravity(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetMass(obj Object, mass float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(mass)
	CallSpriteSetMass(arg0, arg1)
}
func (pself *spriteMgr) GetMass(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetMass(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) AddForce(obj Object, force Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(force)
	CallSpriteAddForce(arg0, arg1)
}
func (pself *spriteMgr) AddImpulse(obj Object, impulse Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(impulse)
	CallSpriteAddImpulse(arg0, arg1)
}
func (pself *spriteMgr) SetCollisionLayer(obj Object, layer int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(layer)
	CallSpriteSetCollisionLayer(arg0, arg1)
}
func (pself *spriteMgr) GetCollisionLayer(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetCollisionLayer(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetCollisionMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	CallSpriteSetCollisionMask(arg0, arg1)
}
func (pself *spriteMgr) GetCollisionMask(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetCollisionMask(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetTriggerLayer(obj Object, layer int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(layer)
	CallSpriteSetTriggerLayer(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerLayer(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetTriggerLayer(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetTriggerMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	CallSpriteSetTriggerMask(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerMask(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetTriggerMask(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetColliderRect(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	CallSpriteSetColliderRect(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetColliderCircle(obj Object, center Vec2, radius float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdFloat(radius)
	CallSpriteSetColliderCircle(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetColliderCapsule(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	CallSpriteSetColliderCapsule(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetCollisionEnabled(obj Object, enabled bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(enabled)
	CallSpriteSetCollisionEnabled(arg0, arg1)
}
func (pself *spriteMgr) IsCollisionEnabled(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsCollisionEnabled(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetTriggerRect(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	CallSpriteSetTriggerRect(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerCircle(obj Object, center Vec2, radius float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdFloat(radius)
	CallSpriteSetTriggerCircle(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerCapsule(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	CallSpriteSetTriggerCapsule(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerEnabled(obj Object, trigger bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(trigger)
	CallSpriteSetTriggerEnabled(arg0, arg1)
}
func (pself *spriteMgr) IsTriggerEnabled(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsTriggerEnabled(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) CreateButton(path string, rect Rect2, text string) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2Str := NewCString(text)
	arg2 := arg2Str.ToGdString() 
	retValue := CallUICreateButton(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) CreateLabel(path string, rect Rect2, text string) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2Str := NewCString(text)
	arg2 := arg2Str.ToGdString() 
	retValue := CallUICreateLabel(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) CreateImage(path string, rect Rect2, color Color) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2 := ToGdColor(color)
	retValue := CallUICreateImage(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) CreateSlider(path string, rect Rect2, value float32) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2 := ToGdFloat(value)
	retValue := CallUICreateSlider(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) CreateToggle(path string, rect Rect2, value bool) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2 := ToGdBool(value)
	retValue := CallUICreateToggle(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) CreateInput(path string, rect Rect2, text string) int64  {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString() 
	arg1 := ToGdRect2(rect)
	arg2Str := NewCString(text)
	arg2 := arg2Str.ToGdString() 
	retValue := CallUICreateInput(arg0, arg1, arg2)
	return ToInt64(retValue)
}
func (pself *uiMgr) GetType(obj Object) int64  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetType(arg0)
	return ToInt64(retValue)
}
func (pself *uiMgr) SetInteractable(obj Object, interactable bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(interactable)
	CallUISetInteractable(arg0, arg1)
}
func (pself *uiMgr) GetInteractable(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetInteractable(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) SetText(obj Object, text string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(text)
	arg1 := arg1Str.ToGdString() 
	CallUISetText(arg0, arg1)
}
func (pself *uiMgr) GetText(obj Object) string  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetText(arg0)
	return ToString(retValue)
}
func (pself *uiMgr) SetRect(obj Object, rect Rect2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdRect2(rect)
	CallUISetRect(arg0, arg1)
}
func (pself *uiMgr) GetRect(obj Object) Rect2  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetRect(arg0)
	return ToRect2(retValue)
}
func (pself *uiMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	CallUISetColor(arg0, arg1)
}
func (pself *uiMgr) GetColor(obj Object) Color  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetColor(arg0)
	return ToColor(retValue)
}
func (pself *uiMgr) SetFontSize(obj Object, size float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(size)
	CallUISetFontSize(arg0, arg1)
}
func (pself *uiMgr) GetFontSize(obj Object) float32  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetFontSize(arg0)
	return ToFloat32(retValue)
}
func (pself *uiMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	CallUISetVisible(arg0, arg1)
}
func (pself *uiMgr) GetVisible(obj Object) bool  {
	arg0 := ToGdObj(obj)
	retValue := CallUIGetVisible(arg0)
	return ToBool(retValue)
}