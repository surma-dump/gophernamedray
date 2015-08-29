package gnr

import (
	"fmt"
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

func (m *Matrix33f) String() string {
	return fmt.Sprintf("[%0.3f, %0.3f, %0.3f; %0.3f, %0.3f, %0.3f; %0.3f, %0.3f, %0.3f]", m.Values[0], m.Values[1], m.Values[2], m.Values[3], m.Values[4], m.Values[5], m.Values[6], m.Values[7], m.Values[8])
}

func (m *Matrix33f) VectorProduct(v *Vector3f) *Vector3f {
	return &Vector3f{
		X: m.Values[0]*v.X + m.Values[1]*v.Y + m.Values[2]*v.Z,
		Y: m.Values[3]*v.X + m.Values[4]*v.Y + m.Values[5]*v.Z,
		Z: m.Values[6]*v.X + m.Values[7]*v.Y + m.Values[8]*v.Z,
	}
}

func (m *Matrix33f) Multiply(f float64) *Matrix33f {
	for i := range m.Values {
		m.Values[i] *= f
	}
	return m
}

func (m *Matrix33f) Add(mo *Matrix33f) *Matrix33f {
	for i := range m.Values {
		m.Values[i] += mo.Values[i]
	}
	return m
}

func MatrixSum(m1, m2 *Matrix33f) *Matrix33f {
	r := &Matrix33f{}
	*r = *m1
	r.Add(m2)
	return r
}

func CopyMatrix33f(m *Matrix33f) *Matrix33f {
	r := &Matrix33f{}
	*r = *m
	return r
}

func RotationMatrix(axis *Vector3f, angle float64) *Matrix33f {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return CopyMatrix33f(MatrixIdentity).Multiply(cos).
		Add(axis.CrossProductMatrix().Multiply(sin)).
		Add(VectorTensorProduct(axis, axis).Multiply(1 - cos))
}

func MatrixEqual(m1, m2 *Matrix33f) bool {
	r := true
	for i := range m1.Values {
		r = r && m1.Values[i] == m2.Values[i]
	}
	return r
}
