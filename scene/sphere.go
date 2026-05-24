package scene

import "math"

type Sphere struct {
	Center Point3
	Radius float64
	Mat    Material
}

func NewSphere(center Point3, radius float64, mat Material) Sphere {
	if radius < 0 {
		radius = 0
	}
	return Sphere{Center: center, Radius: radius, Mat: mat}
}

// Hit uses the simplified quadratic form with h = Dot(d, oc) where
// oc = Center - Origin, so the factor of 2 in the standard quadratic
// is absorbed into h. Roots are (h ± sqrt(h*h - a*c)) / a.
func (s Sphere) Hit(r Ray, rayT Interval) (bool, HitRecord) {
	var rec HitRecord

	oc := Sub(s.Center, r.origin)
	d := r.direction
	a := LenSquared(d)
	h := Dot(d, oc)
	c := LenSquared(oc) - s.Radius*s.Radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false, rec
	}
	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtd) / a
		if !rayT.Surrounds(root) {
			return false, rec
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := MulScalar(Sub(rec.P, s.Center), 1.0/s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = s.Mat
	return true, rec
}
