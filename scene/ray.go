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

var aspectRatio = 16.0 / 9.0
var imageWidth = 400

// Calculate the image height, and ensure that it's at least 1
var imageHeight = int(imageWidth / int(aspectRatio))

func imageHeightCheck(imageHeight int) int {
	if imageHeight < 1 {
		return 1
	} else {
		return imageHeight
	}

}

// Viewport widths less than one are ok since they are real valued.

var viewportHeight = 2.0
var viewportWidth = viewportHeight * float64((imageWidth / imageHeight))
