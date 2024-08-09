package wrap

import (
	. "godot-ext/gdspx/pkg/engine"
	"reflect"
)

type EngineStartFunc func()
type EngineUpdateFunc func(delta float64)
type EngineDestroyFunc func()

var (
	mgrs      []IManager
	callbacks CallbackInfo
)

func createNode(typeName string) GdNode {
	return 0 // TODO
}

func addManager[T IManager](mgr T) T {
	typeName := reflect.TypeOf(mgr).Elem().Name()
	println("add manager", typeName)
	node := createNode("SpxNode")
	mgr.Init(node)
	mgrs = append(mgrs, mgr)
	return mgr
}

func CreateMgrs() []IManager {
	return createMgrs()
}

func RegisterCallbacks(callback EngineCallbackInfo) {
	callbacks = CallbackInfo{}
	callbacks.OnEngineStart = callback.OnEngineStart
	callbacks.OnEngineUpdate = callback.OnEngineUpdate
	callbacks.OnEngineDestroy = callback.OnEngineDestroy
	// TODO : register callbacks
}
