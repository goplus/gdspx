package ffi

import (
	"unsafe"
)
import "C"

var dlsymGD func(string) unsafe.Pointer

//go:linkname main main.main
func main()


func BindAPI() (bool) {
	return dlsymGD != nil
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		return getProcAddress(lookupFunc, s)
	}
	FFI.loadProcAddresses()
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(GDExtensionInitializationLevelScene)
	doInitialization(init)
	registerEngineCallback()
	return 1
}
