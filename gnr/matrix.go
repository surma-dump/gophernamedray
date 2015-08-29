package gnr

import (
	"math"
)

var (
	MatrixIdentity = &Matrix33f{
		Values: [9]float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	}
)

type Matrix33f struct {
	// |0 1 2|
	// |3 4 5|
	// |6 7 8|
	Values [9]float64
}

func (m *Matrix33f) VectorProduct(v *Vector3f) *Vector3f {
	return &Vector3f{
		X: m.Values[0]*v.X + m.Values[1]*v.Y + m.Values[2]*v.Z,
		Y: m.Values[3]*v.X + m.Values[4]*v.Y + m.Values[5]*v.Z,
		Z: m.Values[6]*v.X + m.Values[7]*v.Y + m.Values[8]*v.Z,
	}
}

func (m *Matrix33f) Multiply(f float64) *Matrix33f {
	return &Matrix33f{
		Values: [9]float64{
			m.Values[0] * f, m.Values[1] * f, m.Values[2] * f,
			m.Values[3] * f, m.Values[4] * f, m.Values[5] * f,
			m.Values[6] * f, m.Values[7] * f, m.Values[8] * f,
		},
	}
}

func MatrixSum(m1, m2 *Matrix33f) *Matrix33f {
	return &Matrix33f{
		Values: [9]float64{
			m1.Values[0] + m2.Values[0], m1.Values[1] + m2.Values[1], m1.Values[2] + m2.Values[2],
			m1.Values[3] + m2.Values[3], m1.Values[4] + m2.Values[4], m1.Values[5] + m2.Values[5],
			m1.Values[6] + m2.Values[6], m1.Values[7] + m2.Values[7], m1.Values[8] + m2.Values[8],
		},
	}
}

func MatrixCopy(m *Matrix33f) *Matrix33f {
	r := &Matrix33f{}
	*r = *m
	return r
}

func RotationMatrix(axis *Vector3f, angle float64) *Matrix33f {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return MatrixSum(MatrixSum(MatrixIdentity.Multiply(cos), axis.CrossProductMatrix().Multiply(sin)), VectorTensorProduct(axis, axis).Multiply(1-cos))
}
