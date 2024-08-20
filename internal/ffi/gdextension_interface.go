package ffi

/*
#include "gdextension_spx_interface.h"

*/
import "C"

import (
	"godot-ext/gdspx/pkg/engine"
	"unsafe"
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
type GdRect2 C.GdRect2
type GdObj C.GdObj

func ToGdBool(val bool) GdBool {
	if val {
		return GdBool(1)
	} else {
		return GdBool(0)
	}
}
func ToBool(val GdBool) bool {
	return val != 0
}
func ToGdVec2(val engine.Vec2) GdVec2 {
	return GdVec2{C.GdFloat(val.X), C.GdFloat(val.Y)}
}
func ToVec2(val GdVec2) engine.Vec2 {
	return engine.Vec2{float32(val.X), float32(val.Y)}
}
func ToGdColor(val engine.Color) GdColor {
	return GdColor{C.float(val.R), C.float(val.G), C.float(val.B), C.float(val.A)}
}
func ToColor(val GdColor) engine.Color {
	return engine.Color{float32(val.R), float32(val.G), float32(val.B), float32(val.A)}
}
func ToGdRect2(val engine.Rect2) GdRect2 {
	center := ToGdVec2(val.Center)
	size := ToGdVec2(val.Size)
	ret := GdRect2{}
	ret.Center = C.GdVec2(center)
	ret.Size = C.GdVec2(size)
	return ret
}
func ToRect2(val GdRect2) engine.Rect2 {
	ret := engine.Rect2{}
	ret.Center = ToVec2(GdVec2(val.Center))
	ret.Size = ToVec2(GdVec2(val.Size))
	return ret
}
func ToGdObj(val engine.Object) GdObj {
	return GdObj(val)
}
func ToObject(val GdObj) engine.Object {
	return engine.Object(val)
}
func ToGdInt(val int64) GdInt {
	return GdInt(val)
}
func ToInt(val GdInt) int64 {
	return int64(val)
}
func ToInt64(val GdInt) int64 {
	return int64(val)
}
func ToGdFloat(val float32) GdFloat {
	return GdFloat(val)
}
func ToFloat32(val GdFloat) float32 {
	return float32(val)
}
func ToFloat(val GdFloat) float32 {
	return float32(val)
}
func ToString(val GdString) string {
	cstrPtr := (*CString)(unsafe.Pointer(&val))
	return cstrPtr.ToUtf8()
}
func ToGdString(val string) GdString {
	cstr := NewCString(val)
	// TODO (jiepengtan): free cstring
	return cstr.ToGdString()
}

type GDExtensionSpxCallbackInfoPtr C.GDExtensionSpxCallbackInfoPtr
type SpxCallbackInfo C.SpxCallbackInfo

type GDExtensionVariantPtr C.GDExtensionVariantPtr
type GDExtensionConstVariantPtr C.GDExtensionConstVariantPtr
type GDExtensionUninitializedVariantPtr C.GDExtensionUninitializedVariantPtr
type GDExtensionStringNamePtr C.GDExtensionStringNamePtr
type GDExtensionConstStringNamePtr C.GDExtensionConstStringNamePtr
type GDExtensionUninitializedStringNamePtr C.GDExtensionUninitializedStringNamePtr
type GDExtensionStringPtr C.GDExtensionStringPtr
type GDExtensionConstStringPtr C.GDExtensionConstStringPtr
type GDExtensionUninitializedStringPtr C.GDExtensionUninitializedStringPtr
type GDExtensionObjectPtr C.GDExtensionObjectPtr
type GDExtensionConstObjectPtr C.GDExtensionConstObjectPtr
type GDExtensionUninitializedObjectPtr C.GDExtensionUninitializedObjectPtr
type GDExtensionTypePtr C.GDExtensionTypePtr
type GDExtensionConstTypePtr C.GDExtensionConstTypePtr
type GDExtensionUninitializedTypePtr C.GDExtensionUninitializedTypePtr
type GDExtensionMethodBindPtr C.GDExtensionMethodBindPtr
type GDExtensionInt C.GDExtensionInt
type GDExtensionBool C.GDExtensionBool
type GDObjectInstanceID C.GDObjectInstanceID
type GDExtensionRefPtr C.GDExtensionRefPtr
type GDExtensionConstRefPtr C.GDExtensionConstRefPtr

type GDExtensionPtrConstructor C.GDExtensionPtrConstructor
type GDExtensionPtrDestructor C.GDExtensionPtrDestructor
type GDExtensionVariantType C.GDExtensionVariantType

const (
	GDEXTENSION_VARIANT_TYPE_NIL GDExtensionVariantType = iota
	GDEXTENSION_VARIANT_TYPE_BOOL
	GDEXTENSION_VARIANT_TYPE_INT
	GDEXTENSION_VARIANT_TYPE_FLOAT
	GDEXTENSION_VARIANT_TYPE_STRING
	GDEXTENSION_VARIANT_TYPE_VECTOR2
	GDEXTENSION_VARIANT_TYPE_VECTOR2I
	GDEXTENSION_VARIANT_TYPE_RECT2
	GDEXTENSION_VARIANT_TYPE_RECT2I
	GDEXTENSION_VARIANT_TYPE_VECTOR3
	GDEXTENSION_VARIANT_TYPE_VECTOR3I
	GDEXTENSION_VARIANT_TYPE_TRANSFORM2D
	GDEXTENSION_VARIANT_TYPE_VECTOR4
	GDEXTENSION_VARIANT_TYPE_VECTOR4I
	GDEXTENSION_VARIANT_TYPE_PLANE
	GDEXTENSION_VARIANT_TYPE_QUATERNION
	GDEXTENSION_VARIANT_TYPE_AABB
	GDEXTENSION_VARIANT_TYPE_BASIS
	GDEXTENSION_VARIANT_TYPE_TRANSFORM3D
	GDEXTENSION_VARIANT_TYPE_PROJECTION
	GDEXTENSION_VARIANT_TYPE_COLOR
	GDEXTENSION_VARIANT_TYPE_STRING_NAME
	GDEXTENSION_VARIANT_TYPE_NODE_PATH
	GDEXTENSION_VARIANT_TYPE_RID
	GDEXTENSION_VARIANT_TYPE_OBJECT
	GDEXTENSION_VARIANT_TYPE_CALLABLE
	GDEXTENSION_VARIANT_TYPE_SIGNAL
	GDEXTENSION_VARIANT_TYPE_DICTIONARY
	GDEXTENSION_VARIANT_TYPE_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_BYTE_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_INT32_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_INT64_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_FLOAT32_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_FLOAT64_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_STRING_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_VECTOR2_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_VECTOR3_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_COLOR_ARRAY
	GDEXTENSION_VARIANT_TYPE_PACKED_VECTOR4_ARRAY
	GDEXTENSION_VARIANT_TYPE_VARIANT_MAX
)

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
	if callbacks.OnEngineStart != nil {
		callbacks.OnEngineStart()
	}
}

//export func_on_engine_update
func func_on_engine_update(delta C.GDReal) {
	if callbacks.OnEngineUpdate != nil {
		callbacks.OnEngineUpdate(float32(delta))
	}
}

//export func_on_engine_destroy
func func_on_engine_destroy() {
	if callbacks.OnEngineDestroy != nil {
		callbacks.OnEngineDestroy()
	}
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
