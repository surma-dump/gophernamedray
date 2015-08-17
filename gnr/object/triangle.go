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
	normal := gnr.VectorCross(v1, v2).Normalize()
	return Plane{
		Normal:   normal,
		Distance: -gnr.VectorProduct(normal, t.Points[1]),
	}
}

func (t Triangle) RayInteraction(r gnr.Ray) (bool, gnr.Color, gnr.Vector3f, gnr.Vector3f) {
	isInside := true
	ok, color, impact, normal := t.ToPlane().RayInteraction(r)
	if !ok {
		return false, color, impact, normal
	}

	for i := 0; i < 3; i++ {
		p := Triangle{
			Points: [3]gnr.Vector3f{r.Origin, t.Points[i], t.Points[(i+1)%3]},
		}.ToPlane()
		isInside = isInside && p.DistanceToPoint(impact) > 0
	}
	return isInside, color, impact, normal
}
