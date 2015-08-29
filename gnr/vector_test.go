package gnr

import (
	"testing"
)

func TestVector3f_Magnitude(t *testing.T) {
	v1 := &Vector3f{3, 4, 0}
	if m := v1.Magnitude(); m != 5 {
		t.Fatalf("Unexpected result: %f", m)
	}
	if !VectorEqual(v1, &Vector3f{3, 4, 0}) {
		t.Fatalf("Vector changed: %s", v1)
	}
}

func TestVector3f_Normalize(t *testing.T) {
	v1 := &Vector3f{3, 4, 0}
	v1.Normalize()
	if !VectorEqual(v1, &Vector3f{0.6, 0.8, 0}) {
		t.Fatalf("Unexpected result: %s", v1)
	}
}

func TestVector3f_ScalarMultiply(t *testing.T) {
	v1 := &Vector3f{0, 1, 2}
	r := v1.ScalarMultiply(5)
	if !VectorEqual(r, &Vector3f{0, 5, 10}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, r) {
		t.Fatalf("Vector changed: %s", v1)
	}
}

func TestVector3f_Add(t *testing.T) {
	v1 := &Vector3f{1, 1, 1}
	v2 := &Vector3f{1, 2, 3}
	r := v1.Add(v2)
	if !VectorEqual(r, &Vector3f{2, 3, 4}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, r) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVector3f_Subtract(t *testing.T) {
	v1 := &Vector3f{1, 1, 1}
	v2 := &Vector3f{1, 2, 3}
	r := v1.Subtract(v2)
	if !VectorEqual(r, &Vector3f{0, -1, -2}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, r) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVectorSum(t *testing.T) {
	v1 := &Vector3f{1, 1, 1}
	v2 := &Vector3f{1, 2, 3}
	r := VectorSum(v1, v2)

	if !VectorEqual(r, &Vector3f{2, 3, 4}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 1, 1}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVectorDifference(t *testing.T) {
	v1 := &Vector3f{1, 2, 3}
	v2 := &Vector3f{1, 1, 1}
	r := VectorDifference(v1, v2)

	if !VectorEqual(r, &Vector3f{0, 1, 2}) {
		t.Fatalf("Unexpected result: %s", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 1, 1}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVectorProduct(t *testing.T) {
	v1 := &Vector3f{1, 2, 3}
	v2 := &Vector3f{1, 1, 1}
	r := VectorProduct(v1, v2)

	if r != 6 {
		t.Fatalf("Unexpected result: %f", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 1, 1}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVectorCross(t *testing.T) {
	v1 := &Vector3f{1, 2, 3}
	v2 := &Vector3f{1, 1, 1}
	r := VectorCross(v1, v2)

	if !VectorEqual(r, &Vector3f{-1, 2, -1}) {
		t.Fatalf("Unexpected result: %f", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{1, 1, 1}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}

func TestVector3f_CrossProductMatrix(t *testing.T) {
	v1 := &Vector3f{1, 2, 3}
	r := v1.CrossProductMatrix()

	if !MatrixEqual(r, &Matrix33f{[9]float64{
		0, -3, 2,
		3, 0, -1,
		-2, 1, 0,
	}}) {
		t.Fatalf("Unexpected result: %f", r)
	}
	if !VectorEqual(v1, &Vector3f{1, 2, 3}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
}

func TestVectorTensorProduct(t *testing.T) {
	v1 := &Vector3f{1, -2, 3}
	v2 := &Vector3f{-1, 2, -3}
	r := VectorTensorProduct(v1, v2)

	if !MatrixEqual(r, &Matrix33f{[9]float64{
		-1, 2, -3,
		2, -4, 6,
		-3, 6, -9,
	}}) {
		t.Fatalf("Unexpected result: %f", r)
	}
	if !VectorEqual(v1, &Vector3f{1, -2, 3}) {
		t.Fatalf("Vector v1 changed: %s", v1)
	}
	if !VectorEqual(v2, &Vector3f{-1, 2, -3}) {
		t.Fatalf("Vector v2 changed: %s", v2)
	}
}
