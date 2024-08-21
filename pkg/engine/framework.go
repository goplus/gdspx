package engine
import (
	"reflect"
)

var (
	Id2Sprites = make(map[Object]ISpriter)
	TimeSinceGameStart	 = float32(0)
)

func getPrefabPath(name string) string {
	assetName := name
	return "res://assets/prefabs/" +assetName  + ".tscn"
}

func InitEngine(){
	initKeyCode()
}

func CreateSprite[T any]() *T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	name := tType.Name()
	spriteValue := reflect.New(tType).Elem()
	id := SpriteMgr.CreateSprite(getPrefabPath(name))
	sprite := spriteValue.Addr().Interface().(ISpriter)
	sprite.SetId(id)
	Id2Sprites[id] = sprite
	sprite.OnStart()
	return spriteValue.Addr().Interface().(*T)
}
