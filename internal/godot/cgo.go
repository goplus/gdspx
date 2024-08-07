package godot

import (
	"unsafe"
)

import "C"

var dlsymGD func(string) unsafe.Pointer

var (
	godot = API{}
)
func Link() (API, bool) {
	if dlsymGD == nil {
		return godot, false
	}
	return godot, true
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		return get_proc_address(lookupFunc, s)
	}
	linkCGO(&godot)
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(GDExtensionInitializationLevelScene)
	doInitialization(init)
	return 1
}

//go:linkname main main.main
func main()
