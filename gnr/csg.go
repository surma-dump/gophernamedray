package gnr

type Union struct {
	Objects []Object
}

func NewUnion(o ...Object) Union {
	return Union{o}
}

func (u Union) RayInteraction(r Ray) []InteractionResult {
	// Check ray interaction with all objects, only return the one closes to the origin
	irs := ObjectSlice(u.Objects).AggregateSliceInteractionResult(func(irs []InteractionResult, o Object) []InteractionResult {
		newIrs := o.RayInteraction(r)
		return append(irs, newIrs...)
	})
	return InteractionResultSlice(irs).SortBy(InteractionResultDistance)
}

func (u Union) Contains(p Vector3f) bool {
	return ObjectSlice(u.Objects).Any(func(o Object) bool {
		return o.Contains(p)
	})
}
