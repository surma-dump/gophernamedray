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
		hit, color, impact, normal := o.RayInteraction(r)
		if !hit {
			return ir
		}
		newIr := &interactionResult{
			color:  color,
			impact: impact,
			normal: normal,
		}
		newIr.distance = VectorDifference(impact, r.Origin).Magnitude()
		return newIr
	})
	if ir == nil {
		return ColorBlack, math.MaxFloat64
	}
	return ir.color, ir.distance
}
