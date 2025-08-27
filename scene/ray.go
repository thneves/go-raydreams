package scene

import "math"

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

func (ray Ray) hitSphere(center Point3, radius float64) float64 {
	oc := Sub(center, ray.origin)

	a := LenSquared(ray.direction)
	h := Dot(ray.direction, oc)
	c := LenSquared(oc) - radius*radius
	discrimant := h*h - a*c

	if discrimant < 0 {
		return -1.0
	} else {
		return (h - math.Sqrt(discrimant)/a)
	}

}

func (ray Ray) RayColor() Color {
	point3 := Point3{X: 0, Y: 0, Z: -1}

	t := ray.hitSphere(point3, 0.5)

	if t > 0.0 {
		n := Unit(Sub(ray.At(t), point3))
		return MulScalarColors(Color{n.X + 1, n.Y + 1, n.Z + 1}, 0.5)
	}

	unitDir := Unit(ray.direction)
	a := 0.5 * (unitDir.Y + 1.0)
	white := Color{R: 1, G: 1, B: 1}
	sky := Color{R: 0.5, G: 0.7, B: 1.0}
	result := AddColors(MulScalarColors(white, 1-a), MulScalarColors(sky, a))

	return result
}
