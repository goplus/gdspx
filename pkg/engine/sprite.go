package engine

type Sprite struct {
	Id Object
}

func NewSprite(path string) *Sprite {
	pself := &Sprite{}
	pself.Id = SpriteMgr.CreateSprite(path)
	return pself
}
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

func (pself *Sprite) OnTriggerEnter(ISpriter) {}

func (pself *Sprite) OnTriggerExit(ISpriter) {}

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
