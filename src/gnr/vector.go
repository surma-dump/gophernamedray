package gnr

import (
	"fmt"
	"math"
)

type Vector3f struct {
	X, Y, Z float64
}

func (v *Vector3f) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3f) Normalize() *Vector3f {
	r := *v
	m := v.Magnitude()
	if m == 0 {
		return v
	}
	r.X /= m
	r.Y /= m
	r.Z /= m
	return &r
}

func (v *Vector3f) String() string {
	return fmt.Sprintf("(%0.3f, %0.3f, %0.3f)", v.X, v.Y, v.Z)
}

func (v *Vector3f) Multiply(f float64) *Vector3f {
	r := *v
	r.X *= f
	r.Y *= f
	r.Z *= f
	return &r
}

func VectorSum(v1, v2 *Vector3f) *Vector3f {
	return &Vector3f{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func VectorDifference(v1, v2 *Vector3f) *Vector3f {
	return &Vector3f{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}
