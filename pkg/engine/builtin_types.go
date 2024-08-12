package engine

import "fmt"

type Node int64
type Object int64

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	X, Y, Z float64
}

type Vec4 struct {
	X, Y, Z, W float64
}
type Color struct {
	// TODO(jiepengtan) support 32bit
	R, G, B, A float32
}

func (color Color) ToString() string {
	return fmt.Sprintf("Color(%f, %f, %f, %f)", color.R, color.G, color.B, color.A)
}

type Rect2 struct {
	Center, Size Vec2
}
