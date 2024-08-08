package engine

import "fmt"

type Node int64

type Vector2 struct {
	X, Y float64
}

type Vector3 struct {
	X, Y, Z float64
}

type Vector4 struct {
	X, Y, Z, W float64
}
type Color struct {
	// TODO(jiepengtan) support 32bit
	R, G, B, A float64
}

func (color Color) ToString() string {
	return fmt.Sprintf("Color(%f, %f, %f, %f)", color.R, color.G, color.B, color.A)
}

type Rect2 struct {
	Center, Size Vector2
}
