package main

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

func (ray Ray) At(t float64) Point3 {
	// Could be
	// return Point3{
	//    X: ray.origin.X + t*ray.direction.X,
	//    Y: ray.origin.Y + t*ray.direction.Y,
	//    Z: ray.origin.Z + t*ray.direction.Z,
	//}

	return Add(ray.origin, MulScalar(ray.direction, t))
}
