package scene

import "math"

type HitRecord struct {
	P      Point3
	Normal Vec3    // surface normal at P
	T      float64 // ray param where P = oriin + T * direction
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool, HitRecord)
}

type Sphere struct {
	Center Point3
	Radius float64
}

func NewSphere(center Point3, radius float64) Sphere {
	if radius < 0 {
		radius = 0
	}
	return Sphere{Center: center, Radius: radius}
}

// Hit checks the neares valid intersection in [tMin, tMax]

func (s Sphere) Hit(r Ray, tMin, tMax float64) (bool, HitRecord) {
	var rec HitRecord

	oc := Sub(s.Center, r.origin)
	d := r.direction
	a := Dot(d, d)
	h := Dot(d, oc)
	c := LenSquared(oc) - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false, rec
	}

	sqrtd := math.Sqrt(discriminant)

	// Try the nearer root first
	root := (h - sqrtd) / a
	if root <= tMin || root >= tMax {
		root = (h + sqrtd) / a
		if root <= tMin || root >= tMax {
			return false, rec
		}
	}

	// Fill the hit record
	rec.T = root
	rec.P = r.At(rec.T)

	rec.Normal = MulScalar(Sub(rec.P, s.Center), 1.0/s.Radius)
	return true, rec
}
