package engine

func (pself *Sprite) IsAlive() bool {
	return SpriteMgr.IsSpriteAlive(pself.Id)
}
func (pself *Sprite) SetPosition(pos Vec2) {
	SpriteMgr.SetPosition(pself.Id, pos)
}
func (pself *Sprite) SetRotation(rot float32) {
	SpriteMgr.SetRotation(pself.Id, rot)
}
func (pself *Sprite) SetScale(scale Vec2) {
	SpriteMgr.SetScale(pself.Id, scale)
}
func (pself *Sprite) GetPosition() Vec2 {
	return SpriteMgr.GetPosition(pself.Id)
}
func (pself *Sprite) GetRotation() float32 {
	return SpriteMgr.GetRotation(pself.Id)
}
func (pself *Sprite) GetScale() Vec2 {
	return SpriteMgr.GetScale(pself.Id)
}
func (pself *Sprite) SetColor(color Color) {
	SpriteMgr.SetColor(pself.Id, color)
}
func (pself *Sprite) GetColor() Color {
	return SpriteMgr.GetColor(pself.Id)
}
func (pself *Sprite) SetTexture(path string) {
	SpriteMgr.SetTexture(pself.Id, path)
}
func (pself *Sprite) GetTexture() string {
	return SpriteMgr.GetTexture(pself.Id)
}
func (pself *Sprite) SetVisible(visible bool) {
	SpriteMgr.SetVisible(pself.Id, visible)
}
func (pself *Sprite) GetVisible() bool {
	return SpriteMgr.GetVisible(pself.Id)
}
func (pself *Sprite) GetZIndex() int64 {
	return SpriteMgr.GetZIndex(pself.Id)
}
func (pself *Sprite) SetZIndex(z int64) {
	SpriteMgr.SetZIndex(pself.Id, z)
}
func (pself *Sprite) PlayAnim(p_name string, p_custom_scale float32, p_from_end bool) {
	SpriteMgr.PlayAnim(pself.Id, p_name, p_custom_scale, p_from_end)
}
func (pself *Sprite) PlayBackwardsAnim(p_name string) {
	SpriteMgr.PlayBackwardsAnim(pself.Id, p_name)
}
func (pself *Sprite) PauseAnim() {
	SpriteMgr.PauseAnim(pself.Id)
}
func (pself *Sprite) StopAnim() {
	SpriteMgr.StopAnim(pself.Id)
}
func (pself *Sprite) IsPlayingAnim() bool {
	return SpriteMgr.IsPlayingAnim(pself.Id)
}
func (pself *Sprite) SetAnim(p_name string) {
	SpriteMgr.SetAnim(pself.Id, p_name)
}
func (pself *Sprite) GetAnim() string {
	return SpriteMgr.GetAnim(pself.Id)
}
func (pself *Sprite) SetAnimFrame(p_frame int64) {
	SpriteMgr.SetAnimFrame(pself.Id, p_frame)
}
func (pself *Sprite) GetAnimFrame() int64 {
	return SpriteMgr.GetAnimFrame(pself.Id)
}
func (pself *Sprite) SetAnimSpeedScale(p_speed_scale float32) {
	SpriteMgr.SetAnimSpeedScale(pself.Id, p_speed_scale)
}
func (pself *Sprite) GetAnimSpeedScale() float32 {
	return SpriteMgr.GetAnimSpeedScale(pself.Id)
}
func (pself *Sprite) GetAnimPlayingSpeed() float32 {
	return SpriteMgr.GetAnimPlayingSpeed(pself.Id)
}
func (pself *Sprite) SetAnimCentered(p_center bool) {
	SpriteMgr.SetAnimCentered(pself.Id, p_center)
}
func (pself *Sprite) IsAnimCentered() bool {
	return SpriteMgr.IsAnimCentered(pself.Id)
}
func (pself *Sprite) SetAnimOffset(p_offset Vec2) {
	SpriteMgr.SetAnimOffset(pself.Id, p_offset)
}
func (pself *Sprite) GetAnimOffset() Vec2 {
	return SpriteMgr.GetAnimOffset(pself.Id)
}
func (pself *Sprite) SetAnimFlipH(p_flip bool) {
	SpriteMgr.SetAnimFlipH(pself.Id, p_flip)
}
func (pself *Sprite) IsAnimFlippedH() bool {
	return SpriteMgr.IsAnimFlippedH(pself.Id)
}
func (pself *Sprite) SetAnimFlipV(p_flip bool) {
	SpriteMgr.SetAnimFlipV(pself.Id, p_flip)
}
func (pself *Sprite) IsAnimFlippedV() bool {
	return SpriteMgr.IsAnimFlippedV(pself.Id)
}
func (pself *Sprite) SetGravity(gravity float32) {
	SpriteMgr.SetGravity(pself.Id, gravity)
}
func (pself *Sprite) GetGravity() float32 {
	return SpriteMgr.GetGravity(pself.Id)
}
func (pself *Sprite) SetMass(mass float32) {
	SpriteMgr.SetMass(pself.Id, mass)
}
func (pself *Sprite) GetMass() float32 {
	return SpriteMgr.GetMass(pself.Id)
}
func (pself *Sprite) AddForce(force Vec2) {
	SpriteMgr.AddForce(pself.Id, force)
}
func (pself *Sprite) AddImpulse(impulse Vec2) {
	SpriteMgr.AddImpulse(pself.Id, impulse)
}
func (pself *Sprite) SetCollisionLayer(layer int64) {
	SpriteMgr.SetCollisionLayer(pself.Id, layer)
}
func (pself *Sprite) GetCollisionLayer() int64 {
	return SpriteMgr.GetCollisionLayer(pself.Id)
}
func (pself *Sprite) SetCollisionMask(mask int64) {
	SpriteMgr.SetCollisionMask(pself.Id, mask)
}
func (pself *Sprite) GetCollisionMask() int64 {
	return SpriteMgr.GetCollisionMask(pself.Id)
}
func (pself *Sprite) SetTriggerLayer(layer int64) {
	SpriteMgr.SetTriggerLayer(pself.Id, layer)
}
func (pself *Sprite) GetTriggerLayer() int64 {
	return SpriteMgr.GetTriggerLayer(pself.Id)
}
func (pself *Sprite) SetTriggerMask(mask int64) {
	SpriteMgr.SetTriggerMask(pself.Id, mask)
}
func (pself *Sprite) GetTriggerMask() int64 {
	return SpriteMgr.GetTriggerMask(pself.Id)
}
func (pself *Sprite) SetColliderRect(center Vec2, size Vec2) {
	SpriteMgr.SetColliderRect(pself.Id, center, size)
}
func (pself *Sprite) SetColliderCircle(center Vec2, radius float32) {
	SpriteMgr.SetColliderCircle(pself.Id, center, radius)
}
func (pself *Sprite) SetColliderCapsule(center Vec2, size Vec2) {
	SpriteMgr.SetColliderCapsule(pself.Id, center, size)
}
func (pself *Sprite) SetCollisionEnabled(enabled bool) {
	SpriteMgr.SetCollisionEnabled(pself.Id, enabled)
}
func (pself *Sprite) IsCollisionEnabled() bool {
	return SpriteMgr.IsCollisionEnabled(pself.Id)
}
func (pself *Sprite) SetTriggerRect(center Vec2, size Vec2) {
	SpriteMgr.SetTriggerRect(pself.Id, center, size)
}
func (pself *Sprite) SetTriggerCircle(center Vec2, radius float32) {
	SpriteMgr.SetTriggerCircle(pself.Id, center, radius)
}
func (pself *Sprite) SetTriggerCapsule(center Vec2, size Vec2) {
	SpriteMgr.SetTriggerCapsule(pself.Id, center, size)
}
func (pself *Sprite) SetTriggerEnabled(trigger bool) {
	SpriteMgr.SetTriggerEnabled(pself.Id, trigger)
}
func (pself *Sprite) IsTriggerEnabled() bool {
	return SpriteMgr.IsTriggerEnabled(pself.Id)
}
func (pself *Sprite) SetVelocity(velocity Vec2){
	SpriteMgr.SetVelocity(pself.Id, velocity)
}
func (pself *Sprite) GetVelocity() Vec2 {
	return SpriteMgr.GetVelocity(pself.Id)
}
func (pself *Sprite) IsOnFloor() bool {
	return SpriteMgr.IsOnFloor(pself.Id)
}
func (pself *Sprite) IsOnFloorOnly() bool {
	return SpriteMgr.IsOnFloorOnly(pself.Id)
}
func (pself *Sprite) IsOnWall() bool {
	return SpriteMgr.IsOnWall(pself.Id)
}
func (pself *Sprite) IsOnWallOnly() bool {
	return SpriteMgr.IsOnWallOnly(pself.Id)
}
func (pself *Sprite) IsOnCeiling() bool {
	return SpriteMgr.IsOnCeiling(pself.Id)
}
func (pself *Sprite) IsOnCeilingOnly() bool {
	return SpriteMgr.IsOnCeilingOnly(pself.Id)
}
func (pself *Sprite) GetLastMotion() Vec2 {
	return SpriteMgr.GetLastMotion(pself.Id)
}
func (pself *Sprite) GetPositionDelta() Vec2 {
	return SpriteMgr.GetPositionDelta(pself.Id)
}
func (pself *Sprite) GetFloorNormal() Vec2 {
	return SpriteMgr.GetFloorNormal(pself.Id)
}
func (pself *Sprite) GetWallNormal() Vec2 {
	return SpriteMgr.GetWallNormal(pself.Id)
}
func (pself *Sprite) GetRealVelocity() Vec2 {
	return SpriteMgr.GetRealVelocity(pself.Id)
}
func (pself *Sprite) MoveAndSlide() {
	SpriteMgr.MoveAndSlide(pself.Id)
}