package gnr

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Objects        []Object
}

func (s Scene) TracePixel(x, y uint64) (bool, Color, float64, Vector3f) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return s.ShootRay(r)
}

type interactionResult struct {
	color          Color
	impact, normal Vector3f
	distance       float64
}

func (s Scene) ShootRay(r Ray) (bool, Color, float64, Vector3f) {
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
		// ir == nil happens if we are in the very first iteration
		if ir == nil || newIr.distance < ir.distance {
			return newIr
		}
		return ir
	})
	if ir == nil {
		return false, ColorBlack, 0, r.Direction
	}
	return true, ir.color, ir.distance, ir.normal
}
