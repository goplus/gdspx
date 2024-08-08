package ffi

/*
#include "gdextension_spx_interface.h"

*/
import "C"

import (
	"unsafe"
	"fmt"
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

type initialization = C.GDExtensionInitialization
type initializationLevel = C.GDExtensionInitializationLevel

func doInitialization(init *initialization) {
	C.initialization(init)
}

//export initialize
func initialize(_ unsafe.Pointer, level initializationLevel) {
	if level == 2 {
		main()
	}
}

//export deinitialize
func deinitialize(_ unsafe.Pointer, level initializationLevel) {
	if level == 0 {
	}
}

func get_proc_address(handle uintptr, name string) unsafe.Pointer {
	name = name + "\000"
	char := C.CString(name)
	defer C.free(unsafe.Pointer(char))
	return C.get_proc_address(C.pointer(handle), char)
}


// linkCGO implements the Godot GDExtension API via CGO.
func linkCGO(API *API) {
	register_engine_callback()
	print_error := dlsymGD("print_error")
	API.PrintError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(  
			C.uintptr_t(uintptr(print_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
}

func register_engine_callback() {
	spx_global_register_callbacks := dlsymGD("spx_global_register_callbacks")
	engine_register_func := func() {
		C.spx_global_register_callbacks(
			C.uintptr_t(uintptr(spx_global_register_callbacks)),
		)
	}
	engine_register_func()
}

//export func_on_engine_start
func func_on_engine_start() {
    // TODO: implement
	fmt.Println("go ~~~ funcOnEngineStart")
}
//export func_on_engine_update
func func_on_engine_update(delta C.GDReal) {
    // TODO: implement
}
//export func_on_engine_destroy
func func_on_engine_destroy() {
    // TODO: implement
	fmt.Println("go ~~~ func_on_engine_destroy")
}
//export func_on_sprite_ready
func func_on_sprite_ready(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_sprite_updated
func func_on_sprite_updated(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_sprite_destroyed
func func_on_sprite_destroyed(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_mouse_pressed
func func_on_mouse_pressed(keyid C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_mouse_released
func func_on_mouse_released(keyid C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_key_pressed
func func_on_key_pressed(keyid C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_key_released
func func_on_key_released(keyid C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_action_pressed
func func_on_action_pressed(actionName C.gdstring) {
    // TODO: implement
}
//export func_on_action_just_pressed
func func_on_action_just_pressed(actionName C.gdstring) {
    // TODO: implement
}
//export func_on_action_just_released
func func_on_action_just_released(actionName C.gdstring) {
    // TODO: implement
}
//export func_on_axis_changed
func func_on_axis_changed(actionName C.gdstring, value C.GDReal) {
    // TODO: implement
}
//export func_on_collision_enter
func func_on_collision_enter(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_collision_stay
func func_on_collision_stay(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_collision_exit
func func_on_collision_exit(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_trigger_enter
func func_on_trigger_enter(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_trigger_stay
func func_on_trigger_stay(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_trigger_exit
func func_on_trigger_exit(selfId, otherId C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_ui_pressed
func func_on_ui_pressed(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_ui_released
func func_on_ui_released(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_ui_hovered
func func_on_ui_hovered(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_ui_clicked
func func_on_ui_clicked(id C.GDExtensionInt) {
    // TODO: implement
}
//export func_on_ui_toggle
func func_on_ui_toggle(id C.GDExtensionInt, isOn C.GDExtensionBool) {
    // TODO: implement
}
//export func_on_ui_text_changed
func func_on_ui_text_changed(id C.GDExtensionInt, text C.gdstring) {
    // TODO: implement
}
