package scene

import (
	"fmt"
	"math"
	"math/rand"
)

// Color holds RGB components in linear space, typically in [0, 1].
type Color struct {
	R, G, B float64
}

func newColor(R, G, B float64) Color {
	return Color{R: R, G: G, B: B}
}

func AddColors(u, v Color) Color {
	return newColor(u.R+v.R, u.G+v.G, u.B+v.B)
}

func MulColors(u, v Color) Color {
	return newColor(u.R*v.R, u.G*v.G, u.B*v.B)
}

func MulScalarColors(v Color, t float64) Color {
	return newColor(t*v.R, t*v.G, t*v.B)
}

// RandomColor returns a color with each channel in [0, 1).
func RandomColor() Color {
	return Color{rand.Float64(), rand.Float64(), rand.Float64()}
}

// RandomColorRange returns a color with each channel in [min, max).
func RandomColorRange(min, max float64) Color {
	span := max - min
	return Color{
		min + span*rand.Float64(),
		min + span*rand.Float64(),
		min + span*rand.Float64(),
	}
}

// linearToGamma applies a gamma-2 transform (sqrt) for display output.
func linearToGamma(linear float64) float64 {
	if linear > 0 {
		return math.Sqrt(linear)
	}
	return 0
}

var intensityRange = Interval{Min: 0, Max: 0.999}

func (color Color) WriteColor() {
	r := linearToGamma(color.R)
	g := linearToGamma(color.G)
	b := linearToGamma(color.B)
	rbyte := int(256 * intensityRange.Clamp(r))
	gbyte := int(256 * intensityRange.Clamp(g))
	bbyte := int(256 * intensityRange.Clamp(b))
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
