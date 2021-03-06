package gnr

// +gen * slice:"SortBy,Where,Select[string],Select[*InteractionResult]"
type InteractionResult struct {
	Color         *Vector3f
	PointOfImpact *Vector3f
	Normal        *Vector3f
	Distance      float64
}

func InteractionResultDistance(ir1, ir2 *InteractionResult) bool {
	return ir1.Distance < ir2.Distance
}
