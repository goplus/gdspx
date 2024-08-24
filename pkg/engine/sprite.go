package engine

type Sprite struct {
	Id                  Object
	OnTriggerEnterEvent *Event1[ISpriter]
	OnTriggerExitEvent  *Event1[ISpriter]
}

func (pself *Sprite) onCreate() {
	pself.OnTriggerEnterEvent = NewEvent1[ISpriter]()
	pself.OnTriggerExitEvent = NewEvent1[ISpriter]()
}

func (pself *Sprite) V_OnTriggerEnter(other ISpriter) {
	pself.OnTriggerEnterEvent.Trigger(other)
	pself.OnTriggerEnter(other)
}

func (pself *Sprite) V_OnTriggerExit(other ISpriter) {
	pself.OnTriggerExitEvent.Trigger(other)
	pself.OnTriggerExit(other)
}

func (pself *Sprite) OnTriggerEnter(ISpriter) {}

func (pself *Sprite) OnTriggerExit(ISpriter) {}

func (pself *Sprite) GetId() Object {
	return pself.Id
}
func (pself *Sprite) SetId(id Object) {
	pself.Id = id
}
func (pself *Sprite) Destroy() bool {
	return SpriteMgr.DestroySprite(pself.Id)
}
func (pself *Sprite) OnStart() {
}
func (pself *Sprite) OnUpdate(delta float32) {
}
func (pself *Sprite) OnDestroy() {
}

func (pself *Sprite) AddPos(deltaX, deltaY float32) {
	pos := pself.GetPosition()
	pos.X += deltaX
	pos.Y += deltaY
	pself.SetPosition(pos)
}
func (pself *Sprite) AddPosX(delta float32)  {
	pself.AddPos(delta,0)
}

func (pself *Sprite) AddPosY(delta float32)  {
	pself.AddPos(0,delta)
}

func (pself *Sprite) GetPosX() float32 {
	return pself.GetPosition().X
}

func (pself *Sprite) GetPosY() float32 {
	return pself.GetPosition().Y
}

func (pself *Sprite) SetPosX(value float32) {
	pos := pself.GetPosition()
	pos.X = value
	pself.SetPosition(pos)
}

func (pself *Sprite) SetPosY(value float32) {
	pos := pself.GetPosition()
	pos.Y = value
	pself.SetPosition(pos)
}

func (pself *Sprite) AddVel(deltaX, deltaY float32) {
	pos := pself.GetVelocity()
	pos.X += deltaX
	pos.Y += deltaY
	pself.SetVelocity(pos)
}
func (pself *Sprite) AddVelX(delta float32)  {
	pself.AddVel(delta,0)
}

func (pself *Sprite) AddVelY(delta float32)  {
	pself.AddVel(0,delta)
}

func (pself *Sprite) GetVelX() float32 {
	return pself.GetVelocity().X
}

func (pself *Sprite) GetVelY() float32 {
	return pself.GetVelocity().Y
}

func (pself *Sprite) SetVelX(value float32) {
	pos := pself.GetVelocity()
	pos.X = value
	pself.SetVelocity(pos)
}

func (pself *Sprite) SetVelY(value float32) {
	pos := pself.GetVelocity()
	pos.Y = value
	pself.SetVelocity(pos)
}

func (pself *Sprite) AddScale(deltaX, deltaY float32) {
	pos := pself.GetScale()
	pos.X += deltaX
	pos.Y += deltaY
	pself.SetScale(pos)
}
func (pself *Sprite) AddScaleX(delta float32)  {
	pself.AddScale(delta,0)
}

func (pself *Sprite) AddScaleY(delta float32)  {
	pself.AddScale(0,delta)
}

func (pself *Sprite) GetScaleX() float32 {
	return pself.GetScale().X
}

func (pself *Sprite) GetScaleY() float32 {
	return pself.GetScale().Y
}

func (pself *Sprite) SetScaleX(value float32) {
	pos := pself.GetScale()
	pos.X = value
	pself.SetScale(pos)
}

func (pself *Sprite) SetScaleY(value float32) {
	pos := pself.GetScale()
	pos.Y = value
	pself.SetScale(pos)
}
