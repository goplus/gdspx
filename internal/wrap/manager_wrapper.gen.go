//go:build !platform_web

/*
------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package wrap

import (
	"fmt"
	. "godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/pkg/engine"
	"reflect"
)

func BindMgr(mgrs []IManager) {
	for _, mgr := range mgrs {
		switch v := mgr.(type) {
		case IAudioMgr:
			AudioMgr = v

		case ICameraMgr:
			CameraMgr = v

		case IInputMgr:
			InputMgr = v

		case IPhysicMgr:
			PhysicMgr = v

		case IPlatformMgr:
			PlatformMgr = v

		case ISceneMgr:
			SceneMgr = v

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
type cameraMgr struct {
	baseMgr
}
type inputMgr struct {
	baseMgr
}
type physicMgr struct {
	baseMgr
}
type platformMgr struct {
	baseMgr
}
type sceneMgr struct {
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
	addManager(&cameraMgr{})
	addManager(&inputMgr{})
	addManager(&physicMgr{})
	addManager(&platformMgr{})
	addManager(&sceneMgr{})
	addManager(&spriteMgr{})
	addManager(&uiMgr{})
	return mgrs
}

// call gdextension interface functions

func (pself *audioMgr) PlayAudio(path string) {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	CallAudioPlayAudio(arg0)
}
func (pself *audioMgr) SetAudioVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	CallAudioSetAudioVolume(arg0)
}
func (pself *audioMgr) GetAudioVolume() float32 {
	retValue := CallAudioGetAudioVolume()
	return ToFloat32(retValue)
}
func (pself *audioMgr) IsMusicPlaying() bool {
	retValue := CallAudioIsMusicPlaying()
	return ToBool(retValue)
}
func (pself *audioMgr) PlayMusic(path string) {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	CallAudioPlayMusic(arg0)
}
func (pself *audioMgr) SetMusicVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	CallAudioSetMusicVolume(arg0)
}
func (pself *audioMgr) GetMusicVolume() float32 {
	retValue := CallAudioGetMusicVolume()
	return ToFloat32(retValue)
}
func (pself *audioMgr) PauseMusic() {
	CallAudioPauseMusic()
}
func (pself *audioMgr) ResumeMusic() {
	CallAudioResumeMusic()
}
func (pself *audioMgr) GetMusicTimer() float32 {
	retValue := CallAudioGetMusicTimer()
	return ToFloat32(retValue)
}
func (pself *audioMgr) SetMusicTimer(time float32) {
	arg0 := ToGdFloat(time)
	CallAudioSetMusicTimer(arg0)
}
func (pself *cameraMgr) GetCameraPosition() Vec2 {
	retValue := CallCameraGetCameraPosition()
	return ToVec2(retValue)
}
func (pself *cameraMgr) SetCameraPosition(position Vec2) {
	arg0 := ToGdVec2(position)
	CallCameraSetCameraPosition(arg0)
}
func (pself *cameraMgr) GetCameraZoom() Vec2 {
	retValue := CallCameraGetCameraZoom()
	return ToVec2(retValue)
}
func (pself *cameraMgr) SetCameraZoom(size Vec2) {
	arg0 := ToGdVec2(size)
	CallCameraSetCameraZoom(arg0)
}
func (pself *cameraMgr) GetViewportRect() Rect2 {
	retValue := CallCameraGetViewportRect()
	return ToRect2(retValue)
}
func (pself *inputMgr) GetMousePos() Vec2 {
	retValue := CallInputGetMousePos()
	return ToVec2(retValue)
}
func (pself *inputMgr) GetKey(key int64) bool {
	arg0 := ToGdInt(key)
	retValue := CallInputGetKey(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) GetMouseState(mouse_id int64) bool {
	arg0 := ToGdInt(mouse_id)
	retValue := CallInputGetMouseState(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) GetKeyState(key int64) int64 {
	arg0 := ToGdInt(key)
	retValue := CallInputGetKeyState(arg0)
	return ToInt64(retValue)
}
func (pself *inputMgr) GetAxis(neg_action string, pos_action string) float32 {
	arg0Str := NewCString(neg_action)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1Str := NewCString(pos_action)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallInputGetAxis(arg0, arg1)
	return ToFloat32(retValue)
}
func (pself *inputMgr) IsActionPressed(action string) bool {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallInputIsActionPressed(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) IsActionJustPressed(action string) bool {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallInputIsActionJustPressed(arg0)
	return ToBool(retValue)
}
func (pself *inputMgr) IsActionJustReleased(action string) bool {
	arg0Str := NewCString(action)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallInputIsActionJustReleased(arg0)
	return ToBool(retValue)
}
func (pself *physicMgr) Raycast(from Vec2, to Vec2, collision_mask int64) Object {
	arg0 := ToGdVec2(from)
	arg1 := ToGdVec2(to)
	arg2 := ToGdInt(collision_mask)
	retValue := CallPhysicRaycast(arg0, arg1, arg2)
	return ToObject(retValue)
}
func (pself *physicMgr) CheckCollision(from Vec2, to Vec2, collision_mask int64, collide_with_areas bool, collide_with_bodies bool) bool {
	arg0 := ToGdVec2(from)
	arg1 := ToGdVec2(to)
	arg2 := ToGdInt(collision_mask)
	arg3 := ToGdBool(collide_with_areas)
	arg4 := ToGdBool(collide_with_bodies)
	retValue := CallPhysicCheckCollision(arg0, arg1, arg2, arg3, arg4)
	return ToBool(retValue)
}
func (pself *platformMgr) SetWindowSize(width int64, height int64) {
	arg0 := ToGdInt(width)
	arg1 := ToGdInt(height)
	CallPlatformSetWindowSize(arg0, arg1)
}
func (pself *platformMgr) GetWindowSize() Vec2 {
	retValue := CallPlatformGetWindowSize()
	return ToVec2(retValue)
}
func (pself *platformMgr) SetWindowTitle(title string) {
	arg0Str := NewCString(title)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	CallPlatformSetWindowTitle(arg0)
}
func (pself *platformMgr) GetWindowTitle() string {
	retValue := CallPlatformGetWindowTitle()
	return ToString(retValue)
}
func (pself *platformMgr) SetWindowFullscreen(enable bool) {
	arg0 := ToGdBool(enable)
	CallPlatformSetWindowFullscreen(arg0)
}
func (pself *platformMgr) IsWindowFullscreen() bool {
	retValue := CallPlatformIsWindowFullscreen()
	return ToBool(retValue)
}
func (pself *platformMgr) SetDebugMode(enable bool) {
	arg0 := ToGdBool(enable)
	CallPlatformSetDebugMode(arg0)
}
func (pself *platformMgr) IsDebugMode() bool {
	retValue := CallPlatformIsDebugMode()
	return ToBool(retValue)
}
func (pself *sceneMgr) ChangeSceneToFile(path string) {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	CallSceneChangeSceneToFile(arg0)
}
func (pself *sceneMgr) ReloadCurrentScene() int64 {
	retValue := CallSceneReloadCurrentScene()
	return ToInt64(retValue)
}
func (pself *sceneMgr) UnloadCurrentScene() {
	CallSceneUnloadCurrentScene()
}
func (pself *spriteMgr) SetDontDestroyOnLoad(obj Object) {
	arg0 := ToGdObj(obj)
	CallSpriteSetDontDestroyOnLoad(arg0)
}
func (pself *spriteMgr) SetProcess(obj Object, is_on bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(is_on)
	CallSpriteSetProcess(arg0, arg1)
}
func (pself *spriteMgr) SetPhysicProcess(obj Object, is_on bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(is_on)
	CallSpriteSetPhysicProcess(arg0, arg1)
}
func (pself *spriteMgr) SetChildPosition(obj Object, path string, pos Vec2) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	arg2 := ToGdVec2(pos)
	CallSpriteSetChildPosition(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildPosition(obj Object, path string) Vec2 {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallSpriteGetChildPosition(arg0, arg1)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetChildRotation(obj Object, path string, rot float32) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	arg2 := ToGdFloat(rot)
	CallSpriteSetChildRotation(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildRotation(obj Object, path string) float32 {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallSpriteGetChildRotation(arg0, arg1)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetChildScale(obj Object, path string, scale Vec2) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	arg2 := ToGdVec2(scale)
	CallSpriteSetChildScale(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildScale(obj Object, path string) Vec2 {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallSpriteGetChildScale(arg0, arg1)
	return ToVec2(retValue)
}
func (pself *spriteMgr) CheckCollision(obj Object, target Object, is_src_trigger bool, is_dst_trigger bool) bool {
	arg0 := ToGdObj(obj)
	arg1 := ToGdObj(target)
	arg2 := ToGdBool(is_src_trigger)
	arg3 := ToGdBool(is_dst_trigger)
	retValue := CallSpriteCheckCollision(arg0, arg1, arg2, arg3)
	return ToBool(retValue)
}
func (pself *spriteMgr) CheckCollisionWithPoint(obj Object, point Vec2, is_trigger bool) bool {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(point)
	arg2 := ToGdBool(is_trigger)
	retValue := CallSpriteCheckCollisionWithPoint(arg0, arg1, arg2)
	return ToBool(retValue)
}
func (pself *spriteMgr) CreateSprite(path string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallSpriteCreateSprite(arg0)
	return ToObject(retValue)
}
func (pself *spriteMgr) CloneSprite(obj Object) Object {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteCloneSprite(arg0)
	return ToObject(retValue)
}
func (pself *spriteMgr) DestroySprite(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteDestroySprite(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsSpriteAlive(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsSpriteAlive(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetPosition(obj Object, pos Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(pos)
	CallSpriteSetPosition(arg0, arg1)
}
func (pself *spriteMgr) GetPosition(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetPosition(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetRotation(obj Object, rot float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(rot)
	CallSpriteSetRotation(arg0, arg1)
}
func (pself *spriteMgr) GetRotation(obj Object) float32 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetRotation(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetScale(obj Object, scale Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(scale)
	CallSpriteSetScale(arg0, arg1)
}
func (pself *spriteMgr) GetScale(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetScale(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	CallSpriteSetColor(arg0, arg1)
}
func (pself *spriteMgr) GetColor(obj Object) Color {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetColor(arg0)
	return ToColor(retValue)
}
func (pself *spriteMgr) SetTexture(obj Object, path string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	CallSpriteSetTexture(arg0, arg1)
}
func (pself *spriteMgr) GetTexture(obj Object) string {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetTexture(arg0)
	return ToString(retValue)
}
func (pself *spriteMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	CallSpriteSetVisible(arg0, arg1)
}
func (pself *spriteMgr) GetVisible(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetVisible(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) GetZIndex(obj Object) int64 {
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
	defer arg1Str.Destroy()
	arg2 := ToGdFloat(p_custom_scale)
	arg3 := ToGdBool(p_from_end)
	CallSpritePlayAnim(arg0, arg1, arg2, arg3)
}
func (pself *spriteMgr) PlayBackwardsAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(p_name)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
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
func (pself *spriteMgr) IsPlayingAnim(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsPlayingAnim(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(p_name)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	CallSpriteSetAnim(arg0, arg1)
}
func (pself *spriteMgr) GetAnim(obj Object) string {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnim(arg0)
	return ToString(retValue)
}
func (pself *spriteMgr) SetAnimFrame(obj Object, p_frame int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(p_frame)
	CallSpriteSetAnimFrame(arg0, arg1)
}
func (pself *spriteMgr) GetAnimFrame(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimFrame(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetAnimSpeedScale(obj Object, p_speed_scale float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(p_speed_scale)
	CallSpriteSetAnimSpeedScale(arg0, arg1)
}
func (pself *spriteMgr) GetAnimSpeedScale(obj Object) float32 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimSpeedScale(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) GetAnimPlayingSpeed(obj Object) float32 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimPlayingSpeed(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetAnimCentered(obj Object, p_center bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_center)
	CallSpriteSetAnimCentered(arg0, arg1)
}
func (pself *spriteMgr) IsAnimCentered(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimCentered(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnimOffset(obj Object, p_offset Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(p_offset)
	CallSpriteSetAnimOffset(arg0, arg1)
}
func (pself *spriteMgr) GetAnimOffset(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetAnimOffset(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) SetAnimFlipH(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	CallSpriteSetAnimFlipH(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedH(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimFlippedH(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetAnimFlipV(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	CallSpriteSetAnimFlipV(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedV(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsAnimFlippedV(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) SetVelocity(obj Object, velocity Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(velocity)
	CallSpriteSetVelocity(arg0, arg1)
}
func (pself *spriteMgr) GetVelocity(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetVelocity(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) IsOnFloor(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnFloor(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsOnFloorOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnFloorOnly(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsOnWall(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnWall(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsOnWallOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnWallOnly(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsOnCeiling(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnCeiling(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) IsOnCeilingOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsOnCeilingOnly(arg0)
	return ToBool(retValue)
}
func (pself *spriteMgr) GetLastMotion(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetLastMotion(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) GetPositionDelta(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetPositionDelta(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) GetFloorNormal(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetFloorNormal(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) GetWallNormal(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetWallNormal(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) GetRealVelocity(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetRealVelocity(arg0)
	return ToVec2(retValue)
}
func (pself *spriteMgr) MoveAndSlide(obj Object) {
	arg0 := ToGdObj(obj)
	CallSpriteMoveAndSlide(arg0)
}
func (pself *spriteMgr) SetGravity(obj Object, gravity float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(gravity)
	CallSpriteSetGravity(arg0, arg1)
}
func (pself *spriteMgr) GetGravity(obj Object) float32 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetGravity(arg0)
	return ToFloat32(retValue)
}
func (pself *spriteMgr) SetMass(obj Object, mass float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(mass)
	CallSpriteSetMass(arg0, arg1)
}
func (pself *spriteMgr) GetMass(obj Object) float32 {
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
func (pself *spriteMgr) GetCollisionLayer(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetCollisionLayer(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetCollisionMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	CallSpriteSetCollisionMask(arg0, arg1)
}
func (pself *spriteMgr) GetCollisionMask(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetCollisionMask(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetTriggerLayer(obj Object, layer int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(layer)
	CallSpriteSetTriggerLayer(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerLayer(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteGetTriggerLayer(arg0)
	return ToInt64(retValue)
}
func (pself *spriteMgr) SetTriggerMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	CallSpriteSetTriggerMask(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerMask(obj Object) int64 {
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
func (pself *spriteMgr) IsCollisionEnabled(obj Object) bool {
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
func (pself *spriteMgr) IsTriggerEnabled(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallSpriteIsTriggerEnabled(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) CreateNode(path string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallUiCreateNode(arg0)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateButton(path string, text string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1Str := NewCString(text)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallUiCreateButton(arg0, arg1)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateLabel(path string, text string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1Str := NewCString(text)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallUiCreateLabel(arg0, arg1)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateImage(path string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	retValue := CallUiCreateImage(arg0)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateToggle(path string, value bool) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1 := ToGdBool(value)
	retValue := CallUiCreateToggle(arg0, arg1)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateSlider(path string, value float32) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1 := ToGdFloat(value)
	retValue := CallUiCreateSlider(arg0, arg1)
	return ToObject(retValue)
}
func (pself *uiMgr) CreateInput(path string, text string) Object {
	arg0Str := NewCString(path)
	arg0 := arg0Str.ToGdString()
	defer arg0Str.Destroy()
	arg1Str := NewCString(text)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	retValue := CallUiCreateInput(arg0, arg1)
	return ToObject(retValue)
}
func (pself *uiMgr) DestroyNode(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallUiDestroyNode(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) GetType(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetType(arg0)
	return ToInt64(retValue)
}
func (pself *uiMgr) SetText(obj Object, text string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(text)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	CallUiSetText(arg0, arg1)
}
func (pself *uiMgr) GetText(obj Object) string {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetText(arg0)
	return ToString(retValue)
}
func (pself *uiMgr) SetTexture(obj Object, path string) {
	arg0 := ToGdObj(obj)
	arg1Str := NewCString(path)
	arg1 := arg1Str.ToGdString()
	defer arg1Str.Destroy()
	CallUiSetTexture(arg0, arg1)
}
func (pself *uiMgr) GetTexture(obj Object) string {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetTexture(arg0)
	return ToString(retValue)
}
func (pself *uiMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	CallUiSetColor(arg0, arg1)
}
func (pself *uiMgr) GetColor(obj Object) Color {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetColor(arg0)
	return ToColor(retValue)
}
func (pself *uiMgr) SetFontSize(obj Object, size int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(size)
	CallUiSetFontSize(arg0, arg1)
}
func (pself *uiMgr) GetFontSize(obj Object) int64 {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetFontSize(arg0)
	return ToInt64(retValue)
}
func (pself *uiMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	CallUiSetVisible(arg0, arg1)
}
func (pself *uiMgr) GetVisible(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetVisible(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) SetInteractable(obj Object, interactable bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(interactable)
	CallUiSetInteractable(arg0, arg1)
}
func (pself *uiMgr) GetInteractable(obj Object) bool {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetInteractable(arg0)
	return ToBool(retValue)
}
func (pself *uiMgr) SetRect(obj Object, rect Rect2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdRect2(rect)
	CallUiSetRect(arg0, arg1)
}
func (pself *uiMgr) GetRect(obj Object) Rect2 {
	arg0 := ToGdObj(obj)
	retValue := CallUiGetRect(arg0)
	return ToRect2(retValue)
}
