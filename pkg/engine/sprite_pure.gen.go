//go:build pure_engine
// +build pure_engine

/*------------------------------------------------------------------------------
//   This code was generated by template sprite.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "sprite.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

package engine

import (
	. "github.com/realdream-ai/mathf"
)

func (pself *Sprite) AddForce(force Vec2) {
}

func (pself *Sprite) AddImpulse(impulse Vec2) {
}

func (pself *Sprite) CheckCollision(target Object, is_src_trigger bool, is_dst_trigger bool) bool {
	var _val bool
	return _val
}

func (pself *Sprite) CheckCollisionWithPoint(point Vec2, is_trigger bool) bool {
	var _val bool
	return _val
}

func (pself *Sprite) CloneSprite() Object {
	var _val Object
	return _val
}

func (pself *Sprite) CreateSprite(path string) Object {
	var _val Object
	return _val
}

func (pself *Sprite) DestroySprite() bool {
	var _val bool
	return _val
}

func (pself *Sprite) GetAnim() string {
	var _val string
	return _val
}

func (pself *Sprite) GetAnimFrame() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) GetAnimOffset() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetAnimPlayingSpeed() float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetAnimSpeedScale() float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetChildPosition(path string) Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetChildRotation(path string) float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetChildScale(path string) Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetCollisionLayer() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) GetCollisionMask() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) GetColor() Color {
	var _val Color
	return _val
}

func (pself *Sprite) GetFloorNormal() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetGravity() float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetLastMotion() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetMass() float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetPosition() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetPositionDelta() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetRealVelocity() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetRenderScale() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetRotation() float64 {
	var _val float64
	return _val
}

func (pself *Sprite) GetScale() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetTexture() string {
	var _val string
	return _val
}

func (pself *Sprite) GetTriggerLayer() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) GetTriggerMask() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) GetVelocity() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetVisible() bool {
	var _val bool
	return _val
}

func (pself *Sprite) GetWallNormal() Vec2 {
	var _val Vec2
	return _val
}

func (pself *Sprite) GetZIndex() int64 {
	var _val int64
	return _val
}

func (pself *Sprite) IsAnimCentered() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsAnimFlippedH() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsAnimFlippedV() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsCollisionEnabled() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnCeiling() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnCeilingOnly() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnFloor() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnFloorOnly() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnWall() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsOnWallOnly() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsPlayingAnim() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsSpriteAlive() bool {
	var _val bool
	return _val
}

func (pself *Sprite) IsTriggerEnabled() bool {
	var _val bool
	return _val
}

func (pself *Sprite) MoveAndSlide() {
}

func (pself *Sprite) PauseAnim() {
}

func (pself *Sprite) PlayAnim(p_name string, p_speed float64, isLoop bool, p_revert bool) {
}

func (pself *Sprite) PlayBackwardsAnim(p_name string) {
}

func (pself *Sprite) SetAnim(p_name string) {
}

func (pself *Sprite) SetAnimCentered(p_center bool) {
}

func (pself *Sprite) SetAnimFlipH(p_flip bool) {
}

func (pself *Sprite) SetAnimFlipV(p_flip bool) {
}

func (pself *Sprite) SetAnimFrame(p_frame int64) {
}

func (pself *Sprite) SetAnimOffset(p_offset Vec2) {
}

func (pself *Sprite) SetAnimSpeedScale(p_speed_scale float64) {
}

func (pself *Sprite) SetChildPosition(path string, pos Vec2) {
}

func (pself *Sprite) SetChildRotation(path string, rot float64) {
}

func (pself *Sprite) SetChildScale(path string, scale Vec2) {
}

func (pself *Sprite) SetColliderCapsule(center Vec2, size Vec2) {
}

func (pself *Sprite) SetColliderCircle(center Vec2, radius float64) {
}

func (pself *Sprite) SetColliderRect(center Vec2, size Vec2) {
}

func (pself *Sprite) SetCollisionEnabled(enabled bool) {
}

func (pself *Sprite) SetCollisionLayer(layer int64) {
}

func (pself *Sprite) SetCollisionMask(mask int64) {
}

func (pself *Sprite) SetColor(color Color) {
}

func (pself *Sprite) SetDontDestroyOnLoad() {
}

func (pself *Sprite) SetGravity(gravity float64) {
}

func (pself *Sprite) SetMass(mass float64) {
}

func (pself *Sprite) SetPhysicProcess(is_on bool) {
}

func (pself *Sprite) SetPosition(pos Vec2) {
}

func (pself *Sprite) SetProcess(is_on bool) {
}

func (pself *Sprite) SetRenderScale(scale Vec2) {
}

func (pself *Sprite) SetRotation(rot float64) {
}

func (pself *Sprite) SetScale(scale Vec2) {
}

func (pself *Sprite) SetTexture(path string) {
}

func (pself *Sprite) SetTextureAltas(path string, rect2 Rect2) {
}

func (pself *Sprite) SetTextureAltasDirect(path string, rect2 Rect2) {
}

func (pself *Sprite) SetTextureDirect(path string) {
}

func (pself *Sprite) SetTriggerCapsule(center Vec2, size Vec2) {
}

func (pself *Sprite) SetTriggerCircle(center Vec2, radius float64) {
}

func (pself *Sprite) SetTriggerEnabled(trigger bool) {
}

func (pself *Sprite) SetTriggerLayer(layer int64) {
}

func (pself *Sprite) SetTriggerMask(mask int64) {
}

func (pself *Sprite) SetTriggerRect(center Vec2, size Vec2) {
}

func (pself *Sprite) SetTypeName(type_name string) {
}

func (pself *Sprite) SetVelocity(velocity Vec2) {
}

func (pself *Sprite) SetVisible(visible bool) {
}

func (pself *Sprite) SetZIndex(z int64) {
}

func (pself *Sprite) StopAnim() {
}
