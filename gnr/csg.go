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
		return append(irs, o.RayInteraction(r)...)
	})
	return InteractionResultSlice(irs).SortBy(InteractionResultDistance)
}

func (u Union) Contains(p Vector3f) bool {
	return ObjectSlice(u.Objects).Any(func(o Object) bool {
		return o.Contains(p)
	})
}

type Intersection struct {
	Objects []Object
}

func NewIntersection(o ...Object) Intersection {
	return Intersection{o}
}

func (is Intersection) RayInteraction(r Ray) []InteractionResult {
	irs := ObjectSlice(is.Objects).AggregateSliceInteractionResult(func(irs []InteractionResult, o Object) []InteractionResult {
		return append(irs, o.RayInteraction(r)...)
	})
	irs = InteractionResultSlice(irs).Where(func(ir InteractionResult) bool {
		return is.Contains(ir.PointOfImpact)
	})
	return irs
}

func (is Intersection) Contains(p Vector3f) bool {
	return ObjectSlice(is.Objects).All(func(o Object) bool {
		return o.Contains(p)
	})
}
