package engine

import (
	"math"
)

func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}
func Sign(x float32) int64 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
func Signf(x float32) float32 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
func MoveToward(p_from, p_to, p_delta float32) float32 {
	diff := p_to - p_from
	if Abs(diff) <= p_delta {
		return p_to
	}
	return p_from + Signf(diff)*p_delta
}

func Lerpf(a, b, t float32) float32 {
	return a + (b-a)*t
}

func Clampf(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
func Clamp01f(v float32) float32 {
	return Clampf(v, 0, 1)
}

func LerpVec2(a, b Vec2, t float32) Vec2 {
	return Vec2{Lerpf(a.X, b.X, t), Lerpf(a.Y, b.Y, t)}
}
func ClampVec2(v, min, max Vec2) Vec2 {
	return Vec2{Clampf(v.X, min.X, max.X), Clampf(v.Y, min.Y, max.Y)}
}
func Clamp01Vec2(v Vec2) Vec2 {
	return Vec2{Clamp01f(v.X), Clamp01f(v.Y)}
}
