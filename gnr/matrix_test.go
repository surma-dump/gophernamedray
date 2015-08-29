package gnr

import (
	"testing"
)

func TestMatrix33f_VectorProduct(t *testing.T) {
	m1 := &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}
	v1 := &Vector3f{1, 2, 3}
	r := m1.VectorProduct(v1)
	if !VectorEqual(r, &Vector3f{14, 32, 50}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector changed: %s", v1)
	}
	if !MatrixEqual(m1, &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}) {
		t.Fatalf("Matrix changed: %s", m1)
	}
}

func TestMatrix33f_Multiply(t *testing.T) {
	m1 := &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}
	r := m1.Multiply(2)
	if !MatrixEqual(r, &Matrix33f{[9]float64{
		2, 4, 6,
		8, 10, 12,
		14, 16, 18,
	}}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !MatrixEqual(r, m1) {
		t.Fatalf("Matrix changed: %s", m1)
	}
}

func TestMatrix33f_Add(t *testing.T) {
	m1 := &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}
	m2 := &Matrix33f{[9]float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}}
	r := m1.Add(m2)
	if !MatrixEqual(r, &Matrix33f{[9]float64{
		2, 3, 4,
		5, 6, 7,
		8, 9, 10,
	}}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !MatrixEqual(r, m1) {
		t.Fatalf("Matrix m1 changed: %s", m1)
	}
	if !MatrixEqual(m2, &Matrix33f{[9]float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}}) {
		t.Fatalf("Matrix m2 changed: %s", m2)
	}

}

func TestMatrixSum(t *testing.T) {
	m1 := &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}
	m2 := &Matrix33f{[9]float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}}
	r := MatrixSum(m1, m2)
	if !MatrixEqual(r, &Matrix33f{[9]float64{
		2, 3, 4,
		5, 6, 7,
		8, 9, 10,
	}}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !MatrixEqual(m1, &Matrix33f{[9]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}}) {
		t.Fatalf("Matrix m1 changed: %s", m1)
	}
	if !MatrixEqual(m2, &Matrix33f{[9]float64{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}}) {
		t.Fatalf("Matrix m2 changed: %s", m2)
	}
}

func TestRotationMatrix(t *testing.T) {
	axis := &Vector3f{1, 2, 3}
	angle := 40.0
	r := RotationMatrix(axis, angle)
	// TODO: Test me
	_ = r
	if !MatrixEqual(MatrixIdentity, &Matrix33f{[9]float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}}) {
		t.Fatalf("Matrix MatrixIdentity changed: %s", MatrixIdentity)
	}
	if !VectorEqual(axis, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector axis changed: %s", axis)
	}
}
