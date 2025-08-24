package main

import "math"

// Vec3 represents a 3-dimensional vector with X, Y, and Z components.
// It can be used for positions, directions, colors, or any 3D mathematical operations.
type Vec3 struct {
	X, Y, Z float64
}

// Point3 is a type alias for Vec3, representing a 3D point in space.
type Point3 = Vec3

// newVec3 creates a new Vec3 with the specified X, Y, and Z components.
// This is a constructor function for creating Vec3 instances.
func newVec3(X, Y, Z float64) Vec3 {
	return Vec3{X: X, Y: Y, Z: Z}
}

// Negation returns a new Vec3 with all components negated.
// Result: (-v.X, -v.Y, -v.Z)
// This effectively reverses the direction of the vector.
func (v Vec3) Negation() Vec3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z

	return v
}

// Vec3 UTILITY FUNCTIONS

// Arithmetic Operations

// Add returns the component-wise addition of two vectors u and v.
func Add(u, v Vec3) Vec3 {
	return newVec3(
		u.X+v.X,
		u.Y+v.Y,
		u.Z+v.Z,
	)
}

// Sub returns the component-wise subtraction of vector v from vector u.
func Sub(u, v Vec3) Vec3 {
	return newVec3(
		u.X-v.X,
		u.Y-v.Y,
		u.Z-v.Z,
	)
}

// Mul returns the component-wise multiplication of two vectors u and v.
// Note: This is element-wise multiplication, not dot or cross product.
func Mul(u, v Vec3) Vec3 {
	return newVec3(
		u.X*v.X,
		u.Y*v.Y,
		u.Z*v.Z,
	)
}

// MulScalar returns the scalar multiplication of vector v by scalar t.
// This scales the vector by the given factor.
func MulScalar(v Vec3, t float64) Vec3 {
	return newVec3(
		t*v.X,
		t*v.Y,
		t*v.Z,
	)
}

// DivScalar returns the scalar division of vector v by scalar t.
// Note: No division by zero check is performed.
func DivScalar(v Vec3, t float64) Vec3 {
	return newVec3(v.X/t, v.Y/t, v.Z/t)
}

// Vector Operations

// Dot returns the dot product (scalar product) of vectors u and v.
// The dot product measures the similarity between two vectors.
func Dot(u, v Vec3) float64 {
	return (u.X * v.X) + (u.Y * v.Y) + (u.Z * v.Z)
}

// Cross returns the cross product of vectors u and v.
// The result is a vector perpendicular to both u and v.
// The magnitude equals the area of the parallelogram formed by u and v.
func Cross(u, v Vec3) Vec3 {
	return newVec3(
		(u.Y*v.Z)-(u.Z*v.Y),
		(u.Z*v.X)-(u.X*v.Z),
		(u.X*v.Y)-(u.Y*v.X),
	)
}

// Vector Properties

// Len returns the length (magnitude) of vector v.
// This is the Euclidean norm: sqrt(v.X² + v.Y² + v.Z²)
func Len(v Vec3) float64 {
	return math.Sqrt(LenSquared(v))
}

// LenSquared returns the squared length of vector v.
// Result: v.X² + v.Y² + v.Z²
// This is more efficient than Len when you only need to compare lengths.
func LenSquared(v Vec3) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Unit returns the unit vector (normalized vector) of v.
// The result has the same direction as v but with length 1.
// Note: No zero-length vector check is performed.
func Unit(v Vec3) Vec3 {
	return Vec3{
		v.X / Len(v),
		v.Y / Len(v),
		v.Z / Len(v),
	}
}
