package gnr

type InteractionResult struct {
	Color                 Color
	PointOfImpact, Normal Vector3f
	Distance              float64
}
