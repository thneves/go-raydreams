package scene

// What is a Ray
type Ray struct {
	origin    Point3
	direction Vec3
}

func NewRay(origin, direction Vec3) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

// MATH
// P(t) = A + t * b
// P(t) -> The
func (ray Ray) At(t float64) Point3 {
	// Could be
	// return Point3{
	//    X: ray.origin.X + t*ray.direction.X,
	//    Y: ray.origin.Y + t*ray.direction.Y,
	//    Z: ray.origin.Z + t*ray.direction.Z,
	//}

	return Add(ray.origin, MulScalar(ray.direction, t))
}

func (ray Ray) hitSphere(center Point3, radius float64) bool {
	oc := Sub(center, ray.origin)
	a := Dot(ray.direction, ray.direction)
	b := -2.0 * Dot(ray.direction, oc)
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}

func (ray Ray) RayColor() Color {
	point3 := Point3{X: 0, Y: 0, Z: -1}

	if ray.hitSphere(point3, 0.5) {
		return Color{R: 1, G: 0, B: 0}
	}

	unitDir := Unit(ray.direction)
	a := 0.5 * (unitDir.Y + 1.0)
	white := Color{R: 1, G: 1, B: 1}
	sky := Color{R: 0.5, G: 0.7, B: 1.0}
	result := AddColors(MulScalarColors(white, 1-a), MulScalarColors(sky, a))

	return result
}
