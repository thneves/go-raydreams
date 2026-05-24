package scene

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// Camera owns view geometry and drives the render loop.
// Public fields are inputs; unexported fields are derived in initialize().
type Camera struct {
	AspectRatio     float64
	ImageWidth      int
	SamplesPerPixel int
	MaxDepth        int
	VFov            float64 // vertical field of view in degrees
	LookFrom        Point3
	LookAt          Point3
	VUp             Vec3
	DefocusAngle    float64 // variation angle of rays through each pixel, degrees
	FocusDist       float64 // distance from LookFrom to plane of perfect focus

	imageHeight       int
	pixelSamplesScale float64
	center            Point3
	pixel00Loc        Point3
	pixelDeltaU       Vec3
	pixelDeltaV       Vec3
	u, v, w           Vec3
	defocusDiskU      Vec3
	defocusDiskV      Vec3
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}
	if c.SamplesPerPixel < 1 {
		c.SamplesPerPixel = 1
	}
	if c.MaxDepth < 1 {
		c.MaxDepth = 1
	}
	if c.FocusDist <= 0 {
		c.FocusDist = 1
	}
	c.pixelSamplesScale = 1.0 / float64(c.SamplesPerPixel)
	c.center = c.LookFrom

	theta := c.VFov * math.Pi / 180
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * c.FocusDist
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	c.w = Unit(Sub(c.LookFrom, c.LookAt))
	c.u = Unit(Cross(c.VUp, c.w))
	c.v = Cross(c.w, c.u)

	viewportU := MulScalar(c.u, viewportWidth)
	viewportV := MulScalar(c.v.Negation(), viewportHeight)

	c.pixelDeltaU = DivScalar(viewportU, float64(c.ImageWidth))
	c.pixelDeltaV = DivScalar(viewportV, float64(c.imageHeight))

	viewportUpperLeft := Sub(
		Sub(
			Sub(c.center, MulScalar(c.w, c.FocusDist)),
			MulScalar(viewportU, 0.5),
		),
		MulScalar(viewportV, 0.5),
	)
	c.pixel00Loc = Add(viewportUpperLeft, MulScalar(Add(c.pixelDeltaU, c.pixelDeltaV), 0.5))

	defocusRadius := c.FocusDist * math.Tan(c.DefocusAngle/2*math.Pi/180)
	c.defocusDiskU = MulScalar(c.u, defocusRadius)
	c.defocusDiskV = MulScalar(c.v, defocusRadius)
}

// Render writes the full PPM image to stdout; progress goes to stderr.
func (c *Camera) Render(world Hittable) {
	c.initialize()
	fmt.Printf("P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight)
	for j := 0; j < c.imageHeight; j++ {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d  ", c.imageHeight-j)
		for i := 0; i < c.ImageWidth; i++ {
			pixelColor := Color{}
			for s := 0; s < c.SamplesPerPixel; s++ {
				r := c.getRay(i, j)
				pixelColor = AddColors(pixelColor, c.rayColor(r, c.MaxDepth, world))
			}
			MulScalarColors(pixelColor, c.pixelSamplesScale).WriteColor()
		}
	}
	fmt.Fprintln(os.Stderr, "\rDone.                       ")
}

func (c *Camera) getRay(i, j int) Ray {
	offset := Vec3{rand.Float64() - 0.5, rand.Float64() - 0.5, 0}
	pixelSample := Add(
		c.pixel00Loc,
		Add(
			MulScalar(c.pixelDeltaU, float64(i)+offset.X),
			MulScalar(c.pixelDeltaV, float64(j)+offset.Y),
		),
	)
	var origin Point3
	if c.DefocusAngle <= 0 {
		origin = c.center
	} else {
		origin = c.defocusDiskSample()
	}
	return NewRay(origin, Sub(pixelSample, origin))
}

func (c *Camera) defocusDiskSample() Point3 {
	p := RandomInUnitDisk()
	return Add(c.center, Add(MulScalar(c.defocusDiskU, p.X), MulScalar(c.defocusDiskV, p.Y)))
}

func (c *Camera) rayColor(r Ray, depth int, world Hittable) Color {
	if depth <= 0 {
		return Color{}
	}
	// tMin=0.001 to avoid shadow acne from self-intersection.
	if ok, rec := world.Hit(r, Interval{Min: 0.001, Max: math.Inf(1)}); ok {
		if scatter, attenuation, scattered := rec.Mat.Scatter(r, rec); scatter {
			return MulColors(attenuation, c.rayColor(scattered, depth-1, world))
		}
		return Color{}
	}
	unitDir := Unit(r.direction)
	a := 0.5 * (unitDir.Y + 1.0)
	return AddColors(
		MulScalarColors(Color{R: 1, G: 1, B: 1}, 1-a),
		MulScalarColors(Color{R: 0.5, G: 0.7, B: 1.0}, a),
	)
}
