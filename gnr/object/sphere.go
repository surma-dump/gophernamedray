package object

import (
	"math"

	"go.surma.link/gophernamedray/gnr"
)

type Sphere struct {
	Center *gnr.Vector3f
	Radius float64
}

func (s *Sphere) RayInteraction(r *gnr.Ray) []*gnr.InteractionResult {
	// Ray: {P | P = r.Origin + t * r.Direction}
	// Sphere: {P | mag(P - s.Center)^2 - s.Radius^2 = 0}
	// subsitution
	// mag(r.Origin + t * r.Direction)^2 - s.Radius^2 = 0
	d := gnr.VectorDifference(r.Origin, s.Center)
	p := 2 * gnr.VectorProduct(r.Direction, d)
	q := math.Pow(d.Magnitude(), 2) - math.Pow(s.Radius, 2)
	p2 := math.Pow(p, 2)
	inSqrt := p2/4 - q
	if inSqrt < 0 {
		return []*gnr.InteractionResult{}
	}
	sqrt := math.Sqrt(inSqrt)

	t1 := -p/2 + sqrt
	t2 := -p/2 - sqrt
	P1 := gnr.VectorSum(r.Origin, gnr.CopyVector3f(r.Direction).ScalarMultiply(t1))
	P2 := gnr.VectorSum(r.Origin, gnr.CopyVector3f(r.Direction).ScalarMultiply(t2))
	irs := []*gnr.InteractionResult{
		{Color: gnr.ColorWhite},
		{Color: gnr.ColorWhite},
	}

	irs[0].PointOfImpact = P1
	irs[0].Normal = gnr.VectorDifference(P1, s.Center)
	irs[0].Distance = gnr.VectorDifference(P1, r.Origin).Magnitude()
	if inSqrt != 0 {
		irs[1].PointOfImpact = P2
		irs[1].Normal = gnr.VectorDifference(P2, s.Center)
		irs[1].Distance = gnr.VectorDifference(P2, r.Origin).Magnitude()
	} else {
		irs = irs[0:1]
	}
	return irs
}

func (s *Sphere) Contains(p *gnr.Vector3f) bool {
	// TODO: Research if there’s a way to check this without being susceptible
	// to IEEE 754 rounding errors.
	return gnr.VectorDifference(p, s.Center).Magnitude()-s.Radius < Epsilon
}
