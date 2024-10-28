package sprite

import (
	. "github.com/realdream-ai/gdspx/pkg/engine"
)

type Actor struct {
	Sprite
}

func (pself *Actor) DisablePhysic() {
	pself.SetCollisionLayer(0)
	pself.SetCollisionMask(0)
	pself.SetTriggerLayer(0)
	pself.SetTriggerLayer(0)
}
