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

type CString [8]uint8

func (c *CString) ToGdString() GdString {
	return (GdString)(unsafe.Pointer(c))
}
func (c *CString) NativeConstPtr() GDExtensionConstTypePtr {
	return (GDExtensionConstTypePtr)(unsafe.Pointer(c))
}

func (c *CString) NativePtr() GDExtensionTypePtr {
	return (GDExtensionTypePtr)(unsafe.Pointer(c))
}

func FromGdString(gdstr GdString) *CString {
	return (*CString)(unsafe.Pointer(&gdstr))
}

type stringMethodBindings struct {
	constructor GDExtensionPtrConstructor
	destructor  GDExtensionPtrDestructor
}

var (
	globalStringMethodBindings stringMethodBindings
	nullptr                    = unsafe.Pointer(nil)
)

func stringInitConstructorBindings() {
	globalStringMethodBindings.constructor = CallVariantGetPtrConstructor(GDEXTENSION_VARIANT_TYPE_STRING, 0)
	globalStringMethodBindings.destructor = CallVariantGetPtrDestructor(GDEXTENSION_VARIANT_TYPE_STRING)
}

// constructors
func NewCString(content string) CString {
	return NewCStringWithUtf8Chars(content)
}

func NewCStringEmpty() CString {
	cx := CString{}
	ptr := (GDExtensionUninitializedTypePtr)(unsafe.Pointer(cx.NativePtr()))
	CallBuiltinConstructor(globalStringMethodBindings.constructor, ptr)
	return cx
}

func NewCStringWithLatin1Chars(content string) CString {
	cx := CString{}
	ptr := (GDExtensionUninitializedStringPtr)(unsafe.Pointer(cx.NativePtr()))
	CallStringNewWithLatin1Chars(ptr, content)
	return cx
}

func NewCStringWithUtf8Chars(content string) CString {
	cx := CString{}
	ptr := (GDExtensionUninitializedStringPtr)(unsafe.Pointer(cx.NativePtr()))
	CallStringNewWithUtf8Chars(ptr, content)
	return cx
}

func (cx *CString) Destroy() {
	md := (GDExtensionPtrDestructor)(globalStringMethodBindings.destructor)
	bx := cx.NativePtr()
	CallFunc_GDExtensionPtrDestructor(md, bx)
}

func (cx *CString) String() string {
	return cx.ToUtf8()
}

func (cx *CString) ToAscii() string {
	size := CallStringToLatin1Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(nullptr), (GdInt)(0))
	cstrSlice := make([]C.char, int(size)+1)
	cstr := unsafe.SliceData(cstrSlice)
	CallStringToLatin1Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(cstr), (GdInt)(size+1))
	ret := C.GoString(cstr)[:]
	return ret
}

func (cx *CString) ToUtf8() string {
	size := CallStringToUtf8Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(nullptr), (GdInt)(0))
	cstrSlice := make([]C.char, int(size)+1)
	cstr := unsafe.SliceData(cstrSlice)
	CallStringToUtf8Chars((GDExtensionConstStringPtr)(cx.NativeConstPtr()), (*Char)(cstr), (GdInt)(size+1))
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
	C.cgo_callfn_GDExtensionPtrConstructor(arg0, arg1, arg2)
}

func CallFunc_GDExtensionPtrDestructor(
	fn GDExtensionPtrDestructor,
	p_base GDExtensionTypePtr,
) {
	arg0 := (C.GDExtensionPtrDestructor)(fn)
	arg1 := (C.GDExtensionTypePtr)(p_base)
	C.cgo_callfn_GDExtensionPtrDestructor(arg0, arg1)
}
