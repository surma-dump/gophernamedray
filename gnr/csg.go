package gnr

type Union struct {
	Objects []Object
}

func NewUnion(o ...Object) Union {
	return Union{o}
}

func (u Union) RayInteraction(r Ray) ([]InteractionResult, bool) {
	didHitSomething := false
	// Check ray interaction with all objects, only return the one closes to the origin
	irs := ObjectSlice(u.Objects).AggregateSliceInteractionResult(func(irs []InteractionResult, o Object) []InteractionResult {
		newIrs, ok := o.RayInteraction(r)
		if !ok {
			return irs
		}
		didHitSomething = true
		return append(irs, newIrs...)
	})
	if !didHitSomething {
		return []InteractionResult{}, false
	}
	return InteractionResultSlice(irs).SortBy(InteractionResultDistance), true
}

func (u Union) Contains(p Vector3f) bool {
	return ObjectSlice(u.Objects).Any(func(o Object) bool {
		return o.Contains(p)
	})
}
