package gnr

// +gen slice:"Aggregate[InteractionResult],Aggregate[[]InteractionResult],Any"
type Object interface {
	RayInteraction(r Ray) []InteractionResult
	Contains(p Vector3f) bool
}
