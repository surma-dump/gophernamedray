package gnr

// +gen slice:"Aggregate[*interactionResult]"
type Object interface {
	RayInteraction(r Ray) (collision bool, c Color, impact Vector3f, normal Vector3f)
	Normalize()
}
