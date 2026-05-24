package main

import (
	"math/rand"

	"github.com/thneves/go-raydreams/scene"
)

func main() {
	world := &scene.HittableList{}

	groundMat := scene.Lambertian{Albedo: scene.Color{R: 0.5, G: 0.5, B: 0.5}}
	world.Add(scene.NewSphere(scene.Point3{X: 0, Y: -1000, Z: 0}, 1000, groundMat))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := scene.Point3{
				X: float64(a) + 0.9*rand.Float64(),
				Y: 0.2,
				Z: float64(b) + 0.9*rand.Float64(),
			}
			if scene.Len(scene.Sub(center, scene.Point3{X: 4, Y: 0.2, Z: 0})) <= 0.9 {
				continue
			}
			var mat scene.Material
			switch {
			case chooseMat < 0.8:
				albedo := scene.MulColors(scene.RandomColor(), scene.RandomColor())
				mat = scene.Lambertian{Albedo: albedo}
			case chooseMat < 0.95:
				albedo := scene.RandomColorRange(0.5, 1)
				fuzz := 0.5 * rand.Float64()
				mat = scene.Metal{Albedo: albedo, Fuzz: fuzz}
			default:
				mat = scene.Dielectric{RefractionIndex: 1.5}
			}
			world.Add(scene.NewSphere(center, 0.2, mat))
		}
	}

	world.Add(scene.NewSphere(scene.Point3{X: 0, Y: 1, Z: 0}, 1.0,
		scene.Dielectric{RefractionIndex: 1.5}))
	world.Add(scene.NewSphere(scene.Point3{X: -4, Y: 1, Z: 0}, 1.0,
		scene.Lambertian{Albedo: scene.Color{R: 0.4, G: 0.2, B: 0.1}}))
	world.Add(scene.NewSphere(scene.Point3{X: 4, Y: 1, Z: 0}, 1.0,
		scene.Metal{Albedo: scene.Color{R: 0.7, G: 0.6, B: 0.5}, Fuzz: 0.0}))

	cam := &scene.Camera{
		AspectRatio:     16.0 / 9.0,
		ImageWidth:      400,
		SamplesPerPixel: 50,
		MaxDepth:        20,
		VFov:            20,
		LookFrom:        scene.Point3{X: 13, Y: 2, Z: 3},
		LookAt:          scene.Point3{X: 0, Y: 0, Z: 0},
		VUp:             scene.Vec3{X: 0, Y: 1, Z: 0},
		DefocusAngle:    0.6,
		FocusDist:       10.0,
	}
	cam.Render(world)
}
