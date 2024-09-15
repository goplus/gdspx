package webffi

import (
	"encoding/json"
	"fmt"
	. "godot-ext/gdspx/pkg/engine"
	"syscall/js"
)

func jsValue2Go(value js.Value) interface{} {
	switch value.Type() {
	case js.TypeObject:
		obj := make(map[string]interface{})
		keys := js.Global().Get("Object").Call("keys", value)
		for i := 0; i < keys.Length(); i++ {
			key := keys.Index(i).String()
			obj[key] = jsValue2Go(value.Get(key)) // 递归处理嵌套对象
		}
		return obj
	case js.TypeString:
		return value.String()
	case js.TypeNumber:
		return value.Float()
	case js.TypeBoolean:
		return value.Bool()
	default:
		return nil
	}
}
func PrintJs(rect js.Value) {
	rectMap := jsValue2Go(rect)
	jsonData, err := json.Marshal(rectMap)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func JsFromGdObj(val Object) js.Value {
	return JsFromGdInt(int64(val))
}

func JsFromGdInt(val int64) js.Value {
	vec2Js := js.Global().Get("Object").New()

	low := uint32(val & 0xFFFFFFFF)
	high := uint32((val >> 32) & 0xFFFFFFFF)
	vec2Js.Set("low", low)
	vec2Js.Set("high", high)
	return vec2Js
}

func JsToGdObject(val js.Value) Object {
	return Object(JsToGdInt(val))
}

func JsToGdObj(val js.Value) int64 {
	return JsToGdInt(val)
}

func JsToGdInt(val js.Value) int64 {
	low := uint32(val.Get("low").Int())
	high := uint32(val.Get("high").Int())

	int64Value := int64(high)<<32 | int64(low)
	return int64Value
}

func JsFromGdString(object string) js.Value {
	return js.ValueOf(object)
}

func JsFromGdVec2(vec Vec2) js.Value {
	vec2Js := js.Global().Get("Object").New()
	vec2Js.Set("x", vec.X)
	vec2Js.Set("y", vec.Y)
	return vec2Js
}

func JsFromGdVec3(vec Vec3) js.Value {
	vec3Js := js.Global().Get("Object").New()
	vec3Js.Set("x", vec.X)
	vec3Js.Set("y", vec.Y)
	vec3Js.Set("z", vec.Z)
	return vec3Js
}

func JsFromGdVec4(vec Vec4) js.Value {
	vec4Js := js.Global().Get("Object").New()
	vec4Js.Set("x", vec.X)
	vec4Js.Set("y", vec.Y)
	vec4Js.Set("z", vec.Z)
	vec4Js.Set("w", vec.W)
	return vec4Js
}

func JsFromGdColor(color Color) js.Value {
	colorJs := js.Global().Get("Object").New()
	colorJs.Set("r", color.R)
	colorJs.Set("g", color.G)
	colorJs.Set("b", color.B)
	colorJs.Set("a", color.A)
	return colorJs
}

func JsFromGdRect2(rect Rect2) js.Value {
	rectJs := js.Global().Get("Object").New()
	rectJs.Set("center", JsFromGdVec2(rect.Center))
	rectJs.Set("size", JsFromGdVec2(rect.Size))
	return rectJs
}

func JsFromGdBool(val bool) js.Value {
	return js.ValueOf(val)
}

func JsFromGdFloat(val float32) js.Value {
	return js.ValueOf(val)
}

func JsToGdString(object js.Value) string {
	return object.String()
}

func JsToGdVec2(vec js.Value) Vec2 {
	return Vec2{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
	}
}

func JsToGdVec3(vec js.Value) Vec3 {
	return Vec3{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
		Z: float32(vec.Get("z").Float()),
	}
}

func JsToGdVec4(vec js.Value) Vec4 {
	return Vec4{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
		Z: float32(vec.Get("z").Float()),
		W: float32(vec.Get("w").Float()),
	}
}

func JsToGdColor(color js.Value) Color {
	return Color{
		R: float32(color.Get("r").Float()),
		G: float32(color.Get("g").Float()),
		B: float32(color.Get("b").Float()),
		A: float32(color.Get("a").Float()),
	}
}

func JsToGdRect2(rect js.Value) Rect2 {
	return Rect2{
		Center: JsToGdVec2(rect.Get("center")),
		Size:   JsToGdVec2(rect.Get("size")),
	}
}

func JsToGdBool(val js.Value) bool {
	return val.Bool()
}

func JsToGdFloat(val js.Value) float32 {
	return float32(val.Float())
}

func JsToGdFloat32(val js.Value) float32 {
	return float32(val.Float())
}

func JsToGdInt64(val js.Value) int64 {
	return int64(val.Int())
}
