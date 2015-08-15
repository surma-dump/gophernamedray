package gnr

import (
	"math"
)

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Objects        []Object
}

func (s Scene) TracePixel(x, y uint64) (Color, float64) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return s.ShootRay(r)
}

type interactionResult struct {
	color          Color
	impact, normal Vector3f
	distance       float64
}

func (s Scene) ShootRay(r Ray) (Color, float64) {
	// Check ray interaction with all objects, only return the one closes to the origin
	ir := ObjectSlice(s.Objects).AggregateInteractionResult(func(ir *interactionResult, o Object) *interactionResult {
		if !o.RayCollision(r) {
			return ir
		}
		c, i, n := o.RayInteraction(r)
		newIr := &interactionResult{
			color:  c,
			impact: i,
			normal: n,
		}
		newIr.distance = VectorDifference(i, r.Origin).Magnitude()
		return newIr
	})
	if ir == nil {
		return ColorBlack, math.MaxFloat64
	}
	return ir.color, ir.distance
}
