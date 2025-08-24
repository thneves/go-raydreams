package main

import (
	"fmt"
	"os"

	"github.com/thneves/go-raydreams/scene"
)

func main() {

	// IMAGE

	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	if imageHeight < 1 {
		imageHeight = 1
	}

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	// CAMERA & VIEWPORT

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := scene.Point3{X: 0, Y: 0, Z: 0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.

	viewportU := scene.Vec3{X: viewportWidth, Y: 0, Z: 0}
	viewportV := scene.Vec3{X: 0, Y: -viewportHeight, Z: 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.

	pixelDeltaU := scene.DivScalar(viewportU, float64(imageWidth))
	pixelDeltaV := scene.DivScalar(viewportV, float64(imageHeight))

	// Calculate the location of the upper left pixel

	viewPortCenter := scene.Sub(cameraCenter, scene.Vec3{X: 0, Y: 0, Z: focalLength})
	viewPortUpperLeft := scene.Sub(scene.Sub(viewPortCenter, scene.MulScalar(viewportU, 0.5)), scene.MulScalar(viewportV, 0.5))
	pixel00Loc := scene.Add(viewPortUpperLeft, scene.MulScalar(scene.Add(pixelDeltaU, pixelDeltaV), 0.5))

	// RENDER

	for i := 0; i < imageHeight; i++ {

		remaining := imageHeight - i
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", remaining)

		for j := 0; j < imageWidth; j++ {

			pixelCenter := scene.Add(pixel00Loc, scene.Add(scene.MulScalar(pixelDeltaU, float64(j)), scene.MulScalar(pixelDeltaV, float64(i))))
			rayDirection := scene.Sub(pixelCenter, cameraCenter)

			ray := scene.NewRay(cameraCenter, rayDirection)

			pixelColor := ray.RayColor()
			pixelColor.WriteColor()
		}
	}
	fmt.Fprintln(os.Stderr)
}
