package ffi

/*
#include "gdextension_spx_interface.h"

*/
import "C"

import (
	"unsafe"
	"fmt"
)

type Uint64T C.uint64_t
type Uint32T C.uint32_t
type Uint16T C.uint16_t
type Uint8T C.uint8_t
type Int32T C.int32_t
type Int16T C.int16_t
type Int8T C.int8_t
type Char C.char
type WcharT C.wchar_t
type Char32T C.char32_t
type Char16T C.char16_t

type GdString C.GdString
type GdInt C.GdInt
type GdBool C.GdBool
type GdFloat C.GdFloat
type GdVec4 C.GdVec4
type GdVec3 C.GdVec3
type GdVec2 C.GdVec2
type GdColor C.GdColor
type GdRect C.GdRect

type GDExtensionSpxCallbackInfoPtr C.GDExtensionSpxCallbackInfoPtr
type SpxCallbackInfo C.SpxCallbackInfo

type GDExtensionInitializationLevel int64

const (
	GDExtensionInitializationLevelCore    GDExtensionInitializationLevel = 0
	GDExtensionInitializationLevelServers GDExtensionInitializationLevel = 1
	GDExtensionInitializationLevelScene   GDExtensionInitializationLevel = 2
	GDExtensionInitializationLevelEditor  GDExtensionInitializationLevel = 3
)


type initialization = C.GDExtensionInitialization
type initializationLevel = C.GDExtensionInitializationLevel

func doInitialization(init *initialization) {
	C.initialization(init)
}
func getProcAddress(handle uintptr, name string) unsafe.Pointer {
	name = name + "\000"
	char := C.CString(name)
	defer C.free(unsafe.Pointer(char))
	return C.get_proc_address(C.pointer(handle), char)
}

func registerEngineCallback() {
	spx_global_register_callbacks := dlsymGD("spx_global_register_callbacks")
	C.spx_global_register_callbacks(
		C.pointer(uintptr(spx_global_register_callbacks)),
	)
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
func func_on_action_pressed(actionName C.GdString) {
    // TODO: implement
}
//export func_on_action_just_pressed
func func_on_action_just_pressed(actionName C.GdString) {
    // TODO: implement
}
//export func_on_action_just_released
func func_on_action_just_released(actionName C.GdString) {
    // TODO: implement
}
//export func_on_axis_changed
func func_on_axis_changed(actionName C.GdString, value C.GDReal) {
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
func func_on_ui_text_changed(id C.GDExtensionInt, text C.GdString) {
    // TODO: implement
}
