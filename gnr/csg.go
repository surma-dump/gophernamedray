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

type Difference struct {
	Minuend, Subtrahend Object
}

func (d Difference) RayInteraction(r Ray) []InteractionResult {
	mIrs := d.Minuend.RayInteraction(r)
	sIrs := d.Subtrahend.RayInteraction(r)
	mIrs = InteractionResultSlice(mIrs).Where(func(ir InteractionResult) bool {
		return !d.Subtrahend.Contains(ir.PointOfImpact)
	})
	sIrs = InteractionResultSlice(sIrs).Where(func(ir InteractionResult) bool {
		return d.Minuend.Contains(ir.PointOfImpact)
	})
	return append(mIrs, sIrs...)
}

func (d Difference) Contains(p Vector3f) bool {
	return d.Minuend.Contains(p) && !d.Subtrahend.Contains(p)
}
