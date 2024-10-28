package webffi

import (
	"syscall/js"

	"github.com/realdream-ai/gdspx/pkg/engine"
)

var (
	callbacks engine.CallbackInfo
)

func Link() bool {
	js.Global().Set("goWasmInit", js.FuncOf(goWasmInit))
	API.loadProcAddresses()
	resiterFuncPtr2Js()
	return true
}

func goWasmInit(this js.Value, args []js.Value) interface{} {
	println("Go wasm init succ! ~")
	return js.ValueOf(nil)
}

func BindCallback(info engine.CallbackInfo) {
	callbacks = info
}

func dlsymGD(funcName string) js.Value {
	val := js.Global().Get(funcName)
	if val.IsUndefined() || val.IsNull() {
		panic("Js Function not found: " + funcName)
	}
	return val
}
