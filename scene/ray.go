package scene

import "math"

// What is a Ray
type Ray struct {
	origin    Point3
	direction Vec3
}

type HitRecord struct {
	origin Point3
	vector Vec3
	t      float64
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

func (ray Ray) Hit(center Point3, record HitRecord, rayTmin, rayTmax, radius float64) (t float64, ok bool) {
	oc := Sub(center, ray.origin)

	a := LenSquared(ray.direction)
	h := Dot(ray.direction, oc)
	c := LenSquared(oc) - radius*radius
	discrimant := h*h - a*c

	if discrimant < 0 {
		return -1.0, false
	}

	sqrtd := math.Sqrt(discrimant)

	// Find the neares root that lies in the acceptable range.

	root := (h - sqrtd) / a

	if root <= rayTmin || rayTmax <= root {
		root := (h + sqrtd) / a
		if root <= rayTmin || rayTmax <= root {
			return -1.0, false
		}
	}

	record.t = root
	record.origin = ray.At(record.t)
	record.vector = DivScalarPointer(SubPointer(record.origin, center), radius)

	return 1.0, true
}

func (ray Ray) RayColor() Color {
	center := Point3{X: 0, Y: 0, Z: -1}

	t := ray.Hit(center, 0.5)

	if t > 0.0 {
		n := Unit(Sub(ray.At(t), center))
		return MulScalarColors(Color{n.X + 1, n.Y + 1, n.Z + 1}, 0.5)
	}

	unitDir := Unit(ray.direction)
	a := 0.5 * (unitDir.Y + 1.0)
	white := Color{R: 1, G: 1, B: 1}
	sky := Color{R: 0.5, G: 0.7, B: 1.0}
	result := AddColors(MulScalarColors(white, 1-a), MulScalarColors(sky, a))

	return result
}

func (ray Ray) setFaceNormal(outwardNormal Vec3) {}
