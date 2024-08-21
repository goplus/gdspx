package engine

type Sprite struct {
	Id Object
}

func NewSprite(path string) *Sprite {
	pself := &Sprite{}
	pself.Id = SpriteMgr.CreateSprite(path)
	return pself
}
func (pself *Sprite) GetId( ) Object{
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
