package object

import (
	"math"

	"github.com/surma-dump/gophernamedray/gnr"
)

type Sphere struct {
	Center gnr.Vector3f
	Radius float64
}

func (s Sphere) Normalize() {
}

func (s Sphere) RayInteraction(r gnr.Ray) (gnr.InteractionResult, bool) {
	// Ray: {P | P = r.Origin + t * r.Direction}
	// Sphere: {P | mag(P - s.Center)^2 - s.Radius^2 = 0}
	// subsitution
	// mag(r.Origin + t * r.Direction)^2 - s.Radius^2 = 0
	p := 2 * gnr.VectorProduct(r.Direction, gnr.VectorDifference(r.Origin, s.Center))
	q := math.Pow(gnr.VectorDifference(r.Origin, s.Center).Magnitude(), 2) - math.Pow(s.Radius, 2)
	p2 := math.Pow(p, 2)
	inSqrt := p2/4 - q
	if inSqrt < 0 {
		return gnr.InteractionResult{}, false
	}
	sqrt := math.Sqrt(inSqrt)

	t1 := -p/2 + sqrt
	t2 := -p/2 - sqrt
	P1 := gnr.VectorSum(r.Origin, r.Direction.Multiply(t1))
	P2 := gnr.VectorSum(r.Origin, r.Direction.Multiply(t2))
	d1 := gnr.VectorDifference(r.Origin, P1).Magnitude()
	d2 := gnr.VectorDifference(r.Origin, P2).Magnitude()
	ir := gnr.InteractionResult{
		Color: gnr.ColorWhite,
	}
	if d1 <= d2 {
		ir.PointOfImpact = P1
		ir.Normal = gnr.VectorDifference(P1, s.Center).Normalize()
	} else {
		ir.PointOfImpact = P2
		ir.Normal = gnr.VectorDifference(P2, s.Center).Normalize()
	}
	ir.Distance = gnr.VectorDifference(ir.PointOfImpact, r.Origin).Magnitude()
	return ir, true
}
