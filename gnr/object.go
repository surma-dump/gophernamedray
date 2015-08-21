package gnr

// +gen slice:"Aggregate[InteractionResult],Aggregate[[]InteractionResult]"
type Object interface {
	RayInteraction(r Ray) (ir []InteractionResult, didHit bool)
}
