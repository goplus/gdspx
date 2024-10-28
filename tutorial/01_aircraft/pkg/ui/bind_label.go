package ui

import (
	"fmt"

	. "github.com/realdream-ai/gdspx/pkg/engine"
)

type BindLabel struct {
	UiNode
	value int64
}

func (pself *BindLabel) OnStart() {
	pself.value = 0
}
func (pself *BindLabel) SetValue(value int64) {
	pself.value = value
	pself.refresh()
}
func (pself *BindLabel) AddValue(value int64) {
	pself.value += value
	pself.refresh()
}

func (pself *BindLabel) refresh() {
	txt := fmt.Sprintf("Score: %d", pself.value)
	pself.SetText(txt)
}
