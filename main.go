package main

import (
	"fmt"
	"os"
)

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
