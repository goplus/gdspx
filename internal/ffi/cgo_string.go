package ffi

import (
	"unsafe"
)
/*
#include "gdextension_spx_interface.h"

void cgo_callfn_GDExtensionPtrConstructor(const GDExtensionPtrConstructor fn, GDExtensionUninitializedTypePtr p_base, const GDExtensionConstTypePtr *  p_args) {
    fn(p_base, p_args);
}
void cgo_callfn_GDExtensionPtrDestructor(const GDExtensionPtrDestructor fn, GDExtensionTypePtr p_base) {
    fn(p_base);
}
*/
import "C"

type cstring [8]uint8

func (c *cstring) NativeConstPtr() GDExtensionConstTypePtr {
	return (GDExtensionConstTypePtr)(unsafe.Pointer(c))
}

func (c *cstring) NativePtr() GDExtensionTypePtr {
	return (GDExtensionTypePtr)(unsafe.Pointer(c))
}

type stringMethodBindings struct {
	constructor                      GDExtensionPtrConstructor
	destructor                         GDExtensionPtrDestructor
}

var (
	globalStringMethodBindings stringMethodBindings
	nullptr = unsafe.Pointer(nil)
)

func stringInitConstructorBindings() {
	globalStringMethodBindings.constructor = CallSpxVariantGetPtrConstructor(GDEXTENSION_VARIANT_TYPE_STRING, 0)
	globalStringMethodBindings.destructor = CallSpxVariantGetPtrDestructor(GDEXTENSION_VARIANT_TYPE_STRING)
}

// constructors
func newCString(content string) cstring {
	return newCstringWithUtf8Chars(content)
}

func newCStringEmpty() cstring {
	cx := cstring{}
	ptr := (GDExtensionUninitializedTypePtr)(unsafe.Pointer(cx.NativePtr()))
	CallBuiltinConstructor(globalStringMethodBindings.constructor, ptr)
	return cx
}

func newCstringWithLatin1Chars(content string) cstring {
	cx := cstring{}
	ptr := (GDExtensionUninitializedStringPtr)(unsafe.Pointer(cx.NativePtr()))
	CallSpxStringNewWithLatin1Chars(ptr, content)
	return cx
}

func newCstringWithUtf8Chars(content string) cstring {
	cx := cstring{}
	ptr := (GDExtensionUninitializedStringPtr)(unsafe.Pointer(cx.NativePtr()))
	CallSpxStringNewWithUtf8Chars(ptr, content)
	return cx
}

func (cx *cstring) Destroy() {
	md := (GDExtensionPtrDestructor)(globalStringMethodBindings.destructor)
	bx := cx.NativePtr()
	CallFunc_GDExtensionPtrDestructor(md, bx)
}

func (cx *cstring) String() string {
	return cx.ToUtf8()
}

func (cx *cstring) ToAscii() string {
	size := CallSpxStringToLatin1Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(nullptr), (GdInt)(0))
	cstrSlice := make([]C.char, int(size)+1)
	cstr := unsafe.SliceData(cstrSlice)
	CallSpxStringToLatin1Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(cstr), (GdInt)(size+1))
	ret := C.GoString(cstr)[:]
	return ret
}

func (cx *cstring) ToUtf8() string {
	size := CallSpxStringToUtf8Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(nullptr), (GdInt)(0))
	cstrSlice := make([]C.char, int(size)+1)
	cstr := unsafe.SliceData(cstrSlice)
	CallSpxStringToUtf8Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(cstr), (GdInt)(size+1))
	ret := C.GoString(cstr)[:]
	return ret
}

func CallBuiltinConstructor(constructor GDExtensionPtrConstructor, base GDExtensionUninitializedTypePtr, args ...GDExtensionConstTypePtr) {
	a := (GDExtensionPtrConstructor)(constructor)
	b := (GDExtensionUninitializedTypePtr)(base)
	if a == nil {
		panic("constructor is null")
	}
	c := (*GDExtensionConstTypePtr)(unsafe.SliceData(args))
	CallFunc_GDExtensionPtrConstructor(a, b, c)
}

func CallFunc_GDExtensionPtrConstructor(
	fn GDExtensionPtrConstructor,
	p_base GDExtensionUninitializedTypePtr,
	p_args *GDExtensionConstTypePtr,
) {
	arg0 := (C.GDExtensionPtrConstructor)(fn)
	arg1 := (C.GDExtensionUninitializedTypePtr)(p_base)
	arg2 := (*C.GDExtensionConstTypePtr)(p_args)
	println("called C.cgo_callfn_GDExtensionPtrConstructor")
	C.cgo_callfn_GDExtensionPtrConstructor(arg0, arg1, arg2)
}

func CallFunc_GDExtensionPtrDestructor(
	fn GDExtensionPtrDestructor,
	p_base GDExtensionTypePtr,
) {
	arg0 := (C.GDExtensionPtrDestructor)(fn)
	arg1 := (C.GDExtensionTypePtr)(p_base)
	println("called C.cgo_callfn_GDExtensionPtrDestructor")
	C.cgo_callfn_GDExtensionPtrDestructor(arg0, arg1)
}