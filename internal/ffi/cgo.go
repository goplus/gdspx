package ffi

import (
	"godot-ext/gdspx/pkg/engine"
	"unsafe"
)
import "C"

var (
	dlsymGD   func(string) unsafe.Pointer
	callbacks engine.CallbackInfo
)

//go:linkname main main.main
func main()

func Link() bool {
	return dlsymGD != nil
}

func BindCallback(info engine.CallbackInfo) {
	callbacks = info
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		return getProcAddress(lookupFunc, s)
	}

	builtinAPI.loadProcAddresses()
	api.loadProcAddresses()
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(GDExtensionInitializationLevelScene)
	doInitialization(init)
	registerEngineCallback()
	return 1
}
