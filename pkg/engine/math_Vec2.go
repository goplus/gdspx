package engine

import (
	"math"
)

var (
	Math_PI = float32(3.1415926535897932384626433833)
)

func DegToRad(p_y float32) float32 {
	return p_y * (Math_PI / 180.0)
}

func RadToDeg(p_y float32) float32 {
	return p_y * (180.0 / Math_PI)
}

func (v Vec2) AddY(val float32) Vec2 {
	return Vec2{v.X, v.Y + val}
}
func (v Vec2) AddX(val float32) Vec2 {
	return Vec2{v.X + val, v.Y}
}
func (v Vec2) Subtract(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec2) AngleToPoint(v2 Vec2) float32 {
	return v.Subtract(v2).Angle()
}

func (v Vec2) Angle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}
