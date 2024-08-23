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

func (pself *Sprite) Move(deltaX, deltaY float32) {
	pos := pself.GetPosition()
	pos.X += deltaX
	pos.Y += deltaY
	pself.SetPosition(pos)
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
