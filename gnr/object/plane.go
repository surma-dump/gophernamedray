package object

import (
	"github.com/surma-dump/gophernamedray/gnr"
)

// Plane models an infinite 2D plane in 3D space
type Plane struct {
	// Normal of the plane
	Normal gnr.Vector3f
	// Minimal distance between plane and origin
	Distance float64
}

func (p Plane) Normalize() {
	p.Normal = p.Normal.Normalize()
}

func (p Plane) RayInteraction(r gnr.Ray) (bool, gnr.Color, gnr.Vector3f, gnr.Vector3f) {
	// Ray: {P | P = r.Origin + t * r.Direction}
	// Plane: {P | P * p.Normal + p.Distance = 0}
	// Substitute: (r.Origin + t * r.Direction) * p.Normal + p.Distance = 0
	t := -(p.Distance + gnr.VectorProduct(r.Origin, p.Normal)) / gnr.VectorProduct(r.Direction, p.Normal)
	impact := gnr.VectorSum(r.Direction.Multiply(t), r.Origin)
	return t >= 0, gnr.ColorWhite, impact, p.Normal
}

func (p Plane) DistanceToPoint(pt gnr.Vector3f) float64 {
	return gnr.VectorProduct(p.Normal.Normalize(), pt)
}
