//go:build platform_web

package wrap

import (
	"godot-ext/gdspx/internal/webffi"
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
func LinkFFI() bool {
	return webffi.Link()
}

func OnLinked() {
	println("OnLinked block forever")
	// wasm need Block forever
	c := make(chan struct{})
	<-c
}

func CreateMgrs() []IManager {
	return createMgrs()
}

func RegisterCallbacks(callbacks CallbackInfo) {
	webffi.BindCallback(callbacks)
}
