package gnr

type Union struct {
	Objects []Object
}

func NewUnion(o ...Object) Union {
	return Union{o}
}

func (u Union) RayInteraction(r Ray) (InteractionResult, bool) {
	didHitSomething := false
	// Check ray interaction with all objects, only return the one closes to the origin
	ir := ObjectSlice(u.Objects).AggregateInteractionResult(func(ir InteractionResult, o Object) InteractionResult {
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
