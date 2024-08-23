package engine

import (
	"fmt"
)

type Node int64
type Object int64

type Vec2 struct {
	X, Y float32
}

type Vec3 struct {
	X, Y, Z float32
}

type Vec4 struct {
	X, Y, Z, W float32
}
type Color struct {
	R, G, B, A float32
}

func (color Color) ToString() string {
	return fmt.Sprintf("Color(%f, %f, %f, %f)", color.R, color.G, color.B, color.A)
}

type Rect2 struct {
	Center, Size Vec2
}
