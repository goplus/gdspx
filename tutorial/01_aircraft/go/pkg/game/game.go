package game

import (
	. "gdspx-demo01/pkg/define"
	. "gdspx-demo01/pkg/sprite"
	. "gdspx-demo01/pkg/ui"
	"math/rand"

	. "github.com/realdream-ai/gdspx/pkg/engine"
	. "github.com/realdream-ai/mathf"
)

var (
	timer        = float64(0)
	scoreText    *UiScore
	restartPanel *UiNode
)

func RegisterTypes() {

}
func OnStart() {
	scoreText = CreateUI[UiScore]("Score")
	restartPanel = CreateUI[UiNode]("Restart")
	restartPanel.OnUiClickEvent.Subscribe(func() {
		resetGame()
	})
	resetGame()
}
func resetGame() {
	Score = 0
	timer = 0
	restartPanel.SetVisible(false)
	player := CreateSprite[Aircraft]()
	player.SetPosition(Vec2{0, -WinHeight/2.0 + 200})
	player.OnDieEvent.Subscribe(func() {
		restartPanel.SetVisible(true)
		ClearAllSprites()
	})
}

func OnUpdate(delta float64) {
	if restartPanel.GetVisible() {
		return
	}
	timer += delta
	if timer > 1 {
		timer = 0
		obj := CreateSprite[Enemy]()
		value := (rand.Float64()*2 - 1) * WinWidth / 2.0
		obj.SetPosition(Vec2{value, WinHeight / 2.0})
	}
	scoreText.SetValue(Score)
}

func OnDestroy() {

}
