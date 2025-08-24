package main

import "fmt"

func main() {

	// IMAGE

	image_width := 256
	image_height := 256

	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)

	// RENDER

	// Required Header For a PPM file
	// The P3 means colors are in ASCII, then 255 columns and 255 rows

	for i := 0; i < image_height; i++ {
		for j := 0; j < image_width; j++ {
			r := float64(j) / float64(image_width-1)
			g := float64(i) / float64(image_height-1)
			b := 0.0
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
