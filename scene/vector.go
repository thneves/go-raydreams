package scene

import (
	"math"
	"math/rand"
)

// Vec3 represents a 3-dimensional vector with X, Y, and Z components.
// It can be used for positions, directions, or any 3D mathematical operations.
type Vec3 struct {
	X, Y, Z float64
}

// Point3 is a type alias for Vec3, representing a 3D point in space.
type Point3 = Vec3

func newVec3(X, Y, Z float64) Vec3 {
	return Vec3{X: X, Y: Y, Z: Z}
}

// Negation returns a new Vec3 with all components negated.
func (v Vec3) Negation() Vec3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	return v
}

// Arithmetic Operations

// Add returns the component-wise addition of two vectors u and v.
func Add(u, v Vec3) Vec3 {
	return newVec3(u.X+v.X, u.Y+v.Y, u.Z+v.Z)
}

func AddPointer(u, v Point3) Point3 {
	return Point3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

// Sub returns the component-wise subtraction of vector v from vector u.
func Sub(u, v Vec3) Vec3 {
	return newVec3(u.X-v.X, u.Y-v.Y, u.Z-v.Z)
}

func SubPointer(u, v Point3) Point3 {
	return Point3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

// Mul returns the component-wise multiplication of two vectors u and v.
func Mul(u, v Vec3) Vec3 {
	return newVec3(u.X*v.X, u.Y*v.Y, u.Z*v.Z)
}

// MulScalar returns the scalar multiplication of vector v by scalar t.
func MulScalar(v Vec3, t float64) Vec3 {
	return newVec3(t*v.X, t*v.Y, t*v.Z)
}

// DivScalar returns the scalar division of vector v by scalar t.
func DivScalar(v Vec3, t float64) Vec3 {
	return newVec3(v.X/t, v.Y/t, v.Z/t)
}

func DivScalarPointer(p Point3, t float64) Point3 {
	return Point3{p.X / t, p.Y / t, p.Z / t}
}

// Vector Operations

// Dot returns the dot product of vectors u and v.
func Dot(u, v Vec3) float64 {
	return (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z)
}

// Cross returns the cross product of vectors u and v.
func Cross(u, v Vec3) Vec3 {
	return newVec3(
		(u.Y*v.Z)-(u.Z*v.Y),
		(u.Z*v.X)-(u.X*v.Z),
		(u.X*v.Y)-(u.Y*v.X),
	)
}

// Vector Properties

// Len returns the Euclidean length of vector v.
func Len(v Vec3) float64 {
	return math.Sqrt(LenSquared(v))
}

// LenSquared returns the squared length of vector v.
func LenSquared(v Vec3) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Unit returns the normalized vector of v.
func Unit(v Vec3) Vec3 {
	l := Len(v)
	return Vec3{v.X / l, v.Y / l, v.Z / l}
}

// NearZero reports whether v is close to the zero vector along all axes.
func NearZero(v Vec3) bool {
	const s = 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

// Random returns a vector with each component in [0, 1).
func Random() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

// RandomRange returns a vector with each component in [min, max).
func RandomRange(min, max float64) Vec3 {
	span := max - min
	return Vec3{
		min + span*rand.Float64(),
		min + span*rand.Float64(),
		min + span*rand.Float64(),
	}
}

// RandomInUnitSphere samples a point inside the unit sphere via rejection.
func RandomInUnitSphere() Vec3 {
	for {
		p := RandomRange(-1, 1)
		if LenSquared(p) < 1 {
			return p
		}
	}
}

// RandomUnitVector returns a uniformly distributed unit-length vector.
func RandomUnitVector() Vec3 {
	return Unit(RandomInUnitSphere())
}

// RandomOnHemisphere returns a unit vector in the hemisphere centered on normal.
func RandomOnHemisphere(normal Vec3) Vec3 {
	v := RandomUnitVector()
	if Dot(v, normal) > 0 {
		return v
	}
	return v.Negation()
}

// RandomInUnitDisk samples a point inside the unit disk on the XY plane.
func RandomInUnitDisk() Vec3 {
	for {
		p := Vec3{-1 + 2*rand.Float64(), -1 + 2*rand.Float64(), 0}
		if LenSquared(p) < 1 {
			return p
		}
	}
}

// Reflect reflects v about normal n.
func Reflect(v, n Vec3) Vec3 {
	return Sub(v, MulScalar(n, 2*Dot(v, n)))
}

// Refract bends the unit vector uv at a surface with normal n,
// using etaIOverEtaT = eta_incident / eta_transmitted.
func Refract(uv, n Vec3, etaIOverEtaT float64) Vec3 {
	cosTheta := math.Min(Dot(uv.Negation(), n), 1.0)
	rOutPerp := MulScalar(Add(uv, MulScalar(n, cosTheta)), etaIOverEtaT)
	rOutParallel := MulScalar(n, -math.Sqrt(math.Abs(1.0-LenSquared(rOutPerp))))
	return Add(rOutPerp, rOutParallel)
}
