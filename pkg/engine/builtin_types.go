package engine

import (
	"fmt"
)

type Node int64
type Object int64

type Vec2 struct {
	X, Y float32
}

func (val Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", val.X, val.Y)
}

type Vec3 struct {
	X, Y, Z float32
}

func (val Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", val.X, val.Y, val.Z)
}

type Vec4 struct {
	X, Y, Z, W float32
}

func (val Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", val.X, val.Y, val.Z, val.W)
}

type Color struct {
	R, G, B, A float32
}

func (val Color) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", val.R, val.G, val.B, val.A)
}

type Rect2 struct {
	Position Vec2 // TopLeft point
	Size     Vec2
}

func (val Rect2) String() string {
	return fmt.Sprintf("(position:%s, size:%s)", val.Position.String(), val.Size.String())
}
