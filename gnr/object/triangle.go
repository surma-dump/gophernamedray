package object

import (
	"github.com/surma-dump/gophernamedray/gnr"
)

// Triangle models an 2D triangle in 3D space
type Triangle struct {
	Points [3]gnr.Vector3f
}

func (t Triangle) Normalize() {
}

func (t Triangle) ToPlane() Plane {
	v1 := gnr.VectorDifference(t.Points[1], t.Points[0])
	v2 := gnr.VectorDifference(t.Points[2], t.Points[0])
	normal := gnr.VectorCross(v2, v1).Normalize()
	return Plane{
		Normal:   normal,
		Distance: -gnr.VectorProduct(normal, t.Points[1]),
	}
}

func (t Triangle) RayInteraction(r gnr.Ray) (gnr.InteractionResult, bool) {
	isInside := true
	ir, ok := t.ToPlane().RayInteraction(r)
	if !ok {
		return gnr.InteractionResult{}, false
	}

	for i := 0; i < 3; i++ {
		p := Triangle{
			Points: [3]gnr.Vector3f{r.Origin, t.Points[(i+1)%3], t.Points[i]},
		}.ToPlane()
		isInside = isInside && p.DistanceToPoint(ir.PointOfImpact) > 0
	}
	return gnr.InteractionResult{
		Color:         gnr.ColorWhite,
		PointOfImpact: ir.PointOfImpact,
		Normal:        ir.Normal,
		Distance:      gnr.VectorDifference(ir.PointOfImpact, r.Origin).Magnitude(),
	}, isInside
}
