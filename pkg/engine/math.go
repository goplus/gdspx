package engine


import (
	"math"
)
func Abs(x float32) float32 {
	return float32( math.Abs(float64(x)))
}
func Sign(x float32) float32 {
	if math.Signbit(float64(x)) {
		return -1.0
	}
	return 1.0
}
func MoveToward(p_from, p_to, p_delta float32) float32 {
	diff := p_to - p_from
	if Abs(diff) <= p_delta {
		return p_to
	}
	return p_from + Sign(diff) * p_delta
}