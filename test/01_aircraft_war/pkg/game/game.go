package game

import (
	. "gdspx-demo01/pkg/define"
	. "gdspx-demo01/pkg/ui"
	. "gdspx-demo01/pkg/sprites"
	. "godot-ext/gdspx/pkg/engine"
	"math/rand"
)
var (
	timer = float32(0)
	scoreText *UiScore
)

func OnStart() {
	obj := CreateSprite[Aircraft]()
	obj.SetPosition(Vec2{0, -WinHeight / 2.0 + 200})
	scoreText = CreateUI[UiScore]("Score")
}

func OnUpdate(delta float32) {
	timer+= delta
	if timer > 1 {
		timer = 0
		obj := CreateSprite[Enemy]()
		value := (rand.Float32() *2 -1) * WinWidth / 2.0
		obj.SetPosition(Vec2{value , WinHeight / 2.0})
	}
	scoreText.SetValue(Score)
}
func OnDestroy() {

}