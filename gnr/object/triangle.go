package object

import (
	"fmt"
	"math"

	"github.com/surma-dump/gophernamedray/gnr"
)

// Triangle models an 2D triangle in 3D space
type Triangle struct {
	Points [3]*gnr.Vector3f
}

func (t *Triangle) ToPlane() *Plane {
	v1 := gnr.VectorDifference(t.Points[1], t.Points[0])
	v2 := gnr.VectorDifference(t.Points[2], t.Points[0])
	normal := gnr.VectorCross(v1, v2)
	return &Plane{
		Normal:   normal,
		Distance: -gnr.VectorProduct(normal, t.Points[0]),
	}
}

func (t *Triangle) Contains(p *gnr.Vector3f) bool {
	tp := t.ToPlane()
	na := (&Triangle{
		Points: [3]*gnr.Vector3f{t.Points[0], t.Points[1], p},
	}).ToPlane().Normal
	nb := (&Triangle{
		Points: [3]*gnr.Vector3f{t.Points[1], t.Points[2], p},
	}).ToPlane().Normal
	nc := (&Triangle{
		Points: [3]*gnr.Vector3f{t.Points[2], t.Points[0], p},
	}).ToPlane().Normal

	denom := math.Pow(tp.Normal.Magnitude(), 2)
	ba := gnr.VectorProduct(tp.Normal, na) / denom
	bb := gnr.VectorProduct(tp.Normal, nb) / denom
	bc := gnr.VectorProduct(tp.Normal, nc) / denom

	return ba >= 0 && ba <= 1 && bb >= 0 && bb <= 1 && bc >= 0 && bc <= 1
}

func (t *Triangle) Area() float64 {
	v1 := gnr.VectorDifference(t.Points[1], t.Points[0])
	v2 := gnr.VectorDifference(t.Points[2], t.Points[0])
	return gnr.VectorCross(v1, v2).Magnitude() / 2
}

func (t *Triangle) RayInteraction(r *gnr.Ray) []*gnr.InteractionResult {
	ir := t.ToPlane().RayInteraction(r)
	if len(ir) <= 0 || !t.Contains(ir[0].PointOfImpact) {
		return []*gnr.InteractionResult{}
	}

	return []*gnr.InteractionResult{{
		Color: gnr.ColorWhite,
		// TODO: Are copys necessary?
		PointOfImpact: ir[0].PointOfImpact,
		Normal:        ir[0].Normal,
		Distance:      gnr.VectorDifference(ir[0].PointOfImpact, r.Origin).Magnitude(),
	}}
}

func (t *Triangle) String() string {
	return fmt.Sprintf("Triangle[%s, %s, %s]", t.Points[0], t.Points[1], t.Points[2])
}
