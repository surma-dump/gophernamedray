package gnr

// +gen slice:"Aggregate[*interactionResult]"
type Object interface {
	RayCollision(r Ray) bool
	RayInteraction(r Ray) (c Color, collision Vector3f, normal Vector3f)
	Normalize()
}
