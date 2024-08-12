package wrap

import (
	"godot-ext/gdspx/internal/ffi"
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

func createNode(typeName string) Node {
	return 0 // TODO
}

func addManager[T IManager](mgr T) T {
	typeName := reflect.TypeOf(mgr).Elem().Name()
	println("add manager", typeName)
	mgr.Init(0)
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
	ffi.BindCallback(callbacks)
}
