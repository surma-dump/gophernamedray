package gnr

// +gen slice:"Aggregate[InteractionResult],Aggregate[[]InteractionResult],Any"
type Object interface {
	RayInteraction(r Ray) (ir []InteractionResult, didHit bool)
	Contains(p Vector3f) bool
}
