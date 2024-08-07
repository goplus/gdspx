package gdspx

import (
	"godot-ext/gdspx/internal/godot"
)

func LinkEngine()  {
	_ , ok := godot.Link()
	if(!ok){
		panic("godot not found")
	}
	println("LinkEngine")
}
