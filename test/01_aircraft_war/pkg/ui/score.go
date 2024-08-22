package ui

import (
	"fmt"
	. "godot-ext/gdspx/pkg/engine"
)

type UiScore struct {
	UiNode
	value int64
}

func (pself *UiScore) OnStart() {
	pself.value = 0
}
func (pself *UiScore) SetValue(value int64) {
	pself.value = value
	pself.refresh()
}
func (pself *UiScore) AddValue(value int64) {
	pself.value += value
	pself.refresh()
}

func (pself *UiScore) refresh() {
	txt := fmt.Sprintf("Score: %d", pself.value)
	pself.SetText(txt)
}