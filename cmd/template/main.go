package main

import "C"
import (
	"unsafe"

	"gdspx-demo/pkg/autogen"
	"godot-ext/gdspx/pkg/gdspx"

	. "github.com/godot-go/godot-go/pkg/core"
	. "github.com/godot-go/godot-go/pkg/ffi"
	"github.com/godot-go/godot-go/pkg/log"
)

func main() {
	// log.Trace("this application is meant to be run as a plugin to godot")
}

//export GdExtentionEnterPoint
func GdExtentionEnterPoint(p_get_proc_address unsafe.Pointer, p_library unsafe.Pointer, r_initialization unsafe.Pointer) bool {
	log.Debug("GdExtentionEnterPoint called")
	initObj := NewInitObject(
		(GDExtensionInterfaceGetProcAddress)(p_get_proc_address),
		(GDExtensionClassLibraryPtr)(p_library),
		(*GDExtensionInitialization)(unsafe.Pointer(r_initialization)),
	)

	initObj.RegisterSceneInitializer(func() {
		gdspx.RegisterEngineTypes()
		autogen.RegisterGameTypes()
	})

	initObj.RegisterSceneTerminator(func() {
	})

	return initObj.Init()
}
