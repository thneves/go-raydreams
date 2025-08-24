package main

import (
	"fmt"
	"os"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func newVector(X, Y, Z float64) Vector {
	return Vector{X: X, Y: Y, Z: Z}
}

// Returns the Opposite Vector with Negative Values
func (v Vector) Negation() Vector {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z

	return v
}

// VECTOR UTILITY FUNCTIONS

// Artithmetics
func Add(u, v Vector) Vector {
	return newVector(
		u.X+v.X,
		u.Y+v.Y,
		u.Z+v.Z,
	)
}

func Sub(u, v Vector) Vector {
	return newVector(
		u.X-v.X,
		u.Y-v.Y,
		u.Z-v.Z,
	)
}

func Mul(u, v Vector) Vector {
	return newVector(
		u.X*v.X,
		u.Y*v.Y,
		u.Z*v.Z,
	)
}

func MultiplyScalar(t float64, v Vector) Vector {
	return newVector(
		t*v.X,
		t*v.Y,
		t*v.Z,
	)
}

func Dot(u, v Vector) float64 {
	return (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z)
}

func Cross(u, v Vector) Vector {
	return newVector(
		(u.Y*v.Z)-(u.Z*v.Y),
		(u.Z*v.X)-(u.X*v.Z),
		(u.X*v.Y)-(u.Y*v.X),
	)
}

func main() {

	// IMAGE

	imageWidth := 256
	imageHeight := 256

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	// RENDER

	// Required Header For a PPM file
	// The P3 means colors are in ASCII, then 256 columns and 256 rows

	for i := 0; i < imageHeight; i++ {

		remaining := imageHeight - i
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", remaining)

		for j := 0; j < imageWidth; j++ {
			r := float64(j) / float64(imageWidth-1)
			g := float64(i) / float64(imageHeight-1)
			b := 0.0
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprintln(os.Stderr)
}
