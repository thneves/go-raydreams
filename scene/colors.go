package scene

import "fmt"

// Color is a type alias for Vec3, representing RGB color values.
// Each component (X=R, Y=G, Z=B) should typically be in the range [0, 1].
type Color struct {
	R, G, B float64
}

func newColor(R, G, B float64) Color {
	return Color{
		R: R,
		G: G,
		B: B,
	}
}

func AddColors(u, v Color) Color {
	return newColor(
		u.R+v.R,
		u.G+v.G,
		u.B+v.B,
	)
}

func MulScalarColors(v Color, t float64) Color {
	return newColor(
		t*v.R,
		t*v.G,
		t*v.B,
	)
}

func (color Color) WriteColor() {
	// Translate the [0,1] component values to the byte range [0, 255]
	rbyte := int(255.999 * color.R)
	gbyte := int(255.999 * color.G)
	bbyte := int(255.999 * color.B)

	// Write out the pixel color components
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
