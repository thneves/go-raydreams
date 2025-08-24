package main

import "fmt"

// Color is a type alias for Vec3, representing RGB color values.
// Each component (X=R, Y=G, Z=B) should typically be in the range [0, 1].
type Color struct {
	R, G, B float64
}

func (color Color) WriteColor() {
	// Translate the [0,1] component values to the byte range [0, 255]
	rbyte := int(255.999 * color.R)
	gbyte := int(255.999 * color.G)
	bbyte := int(255.999 * color.B)

	// Write out the pixel color components
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
