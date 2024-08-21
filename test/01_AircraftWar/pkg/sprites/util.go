package sprites

import (
	. "godot-ext/gdspx/pkg/engine"
	"reflect"
)

func getPrefabPath(name string) string {
	return "res://Assets/Prefabs/" + name + ".tscn"
}

type ISpriter interface {
	SetId(Object)
	ILife
}

var (
	Id2Sprites = make(map[Object]ISpriter)
)

func CreateSprite[T any]() *T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	name := tType.Name()
	spriteValue := reflect.New(tType).Elem()
	id := SpriteMgr.CreateSprite(getPrefabPath(name))
	println("create sprite ", name, id)
	sprite := spriteValue.Addr().Interface().(ISpriter)
	sprite.SetId(id)
	Id2Sprites[id] = sprite
	sprite.OnStart()
	return spriteValue.Addr().Interface().(*T)
}
