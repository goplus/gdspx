package engine

type UiNode struct {
	Id Object
}

func NewUiNode(path string) *UiNode {
	pself := &UiNode{}
	pself.Id = Object(UiMgr.CreateNode(path)) 
	return pself
}
func (pself *UiNode) Destroy() bool {
	return UiMgr.DestroyNode(pself.Id)
}
func (pself *UiNode) GetId() Object{
	return pself.Id
}
func (pself *UiNode) SetId(id Object) {
	pself.Id = id
}
func (pself *UiNode) OnStart() {
}
func (pself *UiNode) OnUpdate(delta float32) {
}
func (pself *UiNode) OnDestroy() {
}

