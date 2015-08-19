package gnr

// +gen slice:"Aggregate[InteractionResult],Select[InteractionResult]"
type Object interface {
	RayInteraction(r Ray) (ir InteractionResult, didHit bool)
}
