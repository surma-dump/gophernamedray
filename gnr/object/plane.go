package object

import (
	"math"

	"github.com/surma-dump/gophernamedray/gnr"
)

var (
	// Epsilon is used as a margin of error to compensate for IEEE 754
	// inaccuracies.
	Epsilon = 0.2
)

// Plane models an infinite 2D plane in 3D space
type Plane struct {
	// Normal of the plane
	Normal gnr.Vector3f
	// Minimal distance between plane and origin
	Distance float64
}

func (p Plane) RayInteraction(r gnr.Ray) ([]gnr.InteractionResult, bool) {
	// Ray: {P | P = r.Origin + t * r.Direction}
	// Plane: {P | P * p.Normal + p.Distance = 0}
	// Substitute: (r.Origin + t * r.Direction) * p.Normal + p.Distance = 0
	t := -(p.Distance + gnr.VectorProduct(r.Origin, p.Normal)) / gnr.VectorProduct(r.Direction, p.Normal)
	impact := gnr.VectorSum(r.Direction.Multiply(t), r.Origin)
	return []gnr.InteractionResult{{
		Color:         gnr.ColorWhite,
		PointOfImpact: impact,
		Normal:        p.Normal,
		Distance:      gnr.VectorDifference(impact, r.Origin).Magnitude(),
	}}, t >= 0
}

func (p Plane) DistanceToPoint(pt gnr.Vector3f) float64 {
	return math.Abs(gnr.VectorProduct(p.Normal.Normalize(), pt) + p.Distance)
}

func (p Plane) Contains(pt gnr.Vector3f) bool {
	// TODO: Research if there’s a way to check this without being susceptible
	// to IEEE 754 rounding errors.
	return p.DistanceToPoint(pt) <= Epsilon
}
