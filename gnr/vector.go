package gnr

import (
	"fmt"
	"math"
)

type Vector3f struct {
	X, Y, Z float64
}

func (v Vector3f) Magnitude() float64 {
	return math.Sqrt(VectorProduct(v, v))
}

func (v Vector3f) Normalize() Vector3f {
	m := v.Magnitude()
	if m == 0 {
		return v
	}
	return Vector3f{
		X: v.X / m,
		Y: v.Y / m,
		Z: v.Z / m,
	}
}

func (v Vector3f) String() string {
	return fmt.Sprintf("(%0.3f, %0.3f, %0.3f)", v.X, v.Y, v.Z)
}

func (v Vector3f) Multiply(f float64) Vector3f {
	return Vector3f{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

func VectorSum(v1, v2 Vector3f) Vector3f {
	return Vector3f{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func VectorDifference(v1, v2 Vector3f) Vector3f {
	return Vector3f{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func VectorProduct(v1, v2 Vector3f) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func VectorAngle(v1, v2 Vector3f) float64 {
	return math.Acos(VectorProduct(v1, v2))
}

func VectorCross(v1, v2 Vector3f) Vector3f {
	return Vector3f{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: -v1.X*v2.Z + v1.Z*v2.X,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}
