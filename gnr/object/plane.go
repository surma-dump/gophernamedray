package object

import (
	"math"

	"github.com/surma-dump/gophernamedray/gnr"
)

var (
	// Epsilon is used as a margin of error to compensate for IEEE 754
	// inaccuracies.
	Epsilon = 0.001
)

// Plane models an infinite 2D plane in 3D space
type Plane struct {
	// Normal of the plane
	Normal *gnr.Vector3f
	// Minimal distance between plane and origin
	Distance float64
}

func (p *Plane) RayInteraction(r *gnr.Ray) []*gnr.InteractionResult {
	// Ray: {P | P = r.Origin + t * r.Direction}
	// Plane: {P | P * p.Normal + p.Distance = 0}
	// Substitute: (r.Origin + t * r.Direction) * p.Normal + p.Distance = 0
	t := -(p.Distance + gnr.VectorProduct(r.Origin, p.Normal)) / gnr.VectorProduct(r.Direction, p.Normal)
	if t < 0 {
		return []*gnr.InteractionResult{}
	}
	impact := gnr.CopyVector3f(r.Direction).ScalarMultiply(t).Add(r.Origin)
	return []*gnr.InteractionResult{{
		Color:         gnr.ColorWhite,
		PointOfImpact: impact,
		Normal:        gnr.CopyVector3f(p.Normal),
		Distance:      gnr.VectorDifference(impact, r.Origin).Magnitude(),
	}}
}

func (p *Plane) DistanceToPoint(pt *gnr.Vector3f) float64 {
	n := gnr.CopyVector3f(p.Normal).Normalize()
	return math.Abs(gnr.VectorProduct(n, pt) + p.Distance)
}

func (p *Plane) Contains(pt *gnr.Vector3f) bool {
	// TODO: Research if there’s a way to check this without being susceptible
	// to IEEE 754 rounding errors.
	return p.DistanceToPoint(pt) <= Epsilon
}
