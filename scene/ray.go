package scene

// Ray is a half-line P(t) = origin + t * direction.
type Ray struct {
	origin    Point3
	direction Vec3
}

func NewRay(origin, direction Vec3) Ray {
	return Ray{origin: origin, direction: direction}
}

func (r Ray) Origin() Point3 { return r.origin }

func (r Ray) Direction() Vec3 { return r.direction }

// At evaluates P(t) = origin + t * direction.
func (r Ray) At(t float64) Point3 {
	return Add(r.origin, MulScalar(r.direction, t))
}
