package wrap

import (
	"godot-ext/gdspx/internal/ffi"
	. "godot-ext/gdspx/pkg/engine"
)

type EngineStartFunc func()
type EngineUpdateFunc func(delta float64)
type EngineDestroyFunc func()

var (
	mgrs      []IManager
	callbacks CallbackInfo
)

func addManager[T IManager](mgr T) T {
	//typeName := reflect.TypeOf(mgr).Elem().Name()
	mgr.Init(0)
	mgrs = append(mgrs, mgr)
	return mgr
}

func CreateMgrs() []IManager {
	return createMgrs()
}

func RegisterCallbacks(callbacks CallbackInfo) {
	ffi.BindCallback(callbacks)
}
