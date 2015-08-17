package gnr

type Scene struct {
	GlobalLighting bool
	Camera         Camera
	Objects        []Object
}

func (s Scene) TracePixel(x, y uint64) (InteractionResult, bool) {
	r := s.Camera.GetRayForPixel(x, y)
	r.Intensity = 1.0
	return s.ShootRay(r)
}

func (s Scene) ShootRay(r Ray) (InteractionResult, bool) {
	didHitSomething := false
	// Check ray interaction with all objects, only return the one closes to the origin
	ir := ObjectSlice(s.Objects).AggregateInteractionResult(func(ir InteractionResult, o Object) InteractionResult {
		newIr, ok := o.RayInteraction(r)
		if !ok {
			return ir
		}
		if !didHitSomething || newIr.Distance < ir.Distance {
			didHitSomething = true
			return newIr
		}
		return ir
	})
	if !didHitSomething {
		return InteractionResult{}, false
	}
	return ir, true
}
