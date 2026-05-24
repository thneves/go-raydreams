package scene

import (
	"math"
	"math/rand"
)

// Material decides how an incoming ray scatters off a surface hit.
// Returns whether the ray scatters, the attenuation color, and the scattered ray.
type Material interface {
	Scatter(rIn Ray, rec HitRecord) (bool, Color, Ray)
}

// Lambertian is a matte (diffuse) surface.
type Lambertian struct {
	Albedo Color
}

func (l Lambertian) Scatter(_ Ray, rec HitRecord) (bool, Color, Ray) {
	scatterDir := Add(rec.Normal, RandomUnitVector())
	if NearZero(scatterDir) {
		scatterDir = rec.Normal
	}
	return true, l.Albedo, NewRay(rec.P, scatterDir)
}

// Metal is a reflective surface with optional fuzz roughness.
type Metal struct {
	Albedo Color
	Fuzz   float64
}

func (m Metal) Scatter(rIn Ray, rec HitRecord) (bool, Color, Ray) {
	reflected := Reflect(rIn.direction, rec.Normal)
	fuzz := m.Fuzz
	if fuzz > 1 {
		fuzz = 1
	}
	reflected = Add(Unit(reflected), MulScalar(RandomUnitVector(), fuzz))
	scattered := NewRay(rec.P, reflected)
	if Dot(scattered.direction, rec.Normal) > 0 {
		return true, m.Albedo, scattered
	}
	return false, Color{}, Ray{}
}

// Dielectric is a clear refractive material (glass, water).
type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) Scatter(rIn Ray, rec HitRecord) (bool, Color, Ray) {
	attenuation := Color{R: 1, G: 1, B: 1}
	ri := d.RefractionIndex
	if rec.FrontFace {
		ri = 1.0 / d.RefractionIndex
	}

	unitDir := Unit(rIn.direction)
	cosTheta := math.Min(Dot(unitDir.Negation(), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0
	var direction Vec3
	if cannotRefract || schlick(cosTheta, ri) > rand.Float64() {
		direction = Reflect(unitDir, rec.Normal)
	} else {
		direction = Refract(unitDir, rec.Normal, ri)
	}
	return true, attenuation, NewRay(rec.P, direction)
}

// schlick is Christophe Schlick's polynomial approximation of Fresnel reflectance.
func schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
