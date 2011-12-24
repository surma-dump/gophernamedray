package gnr

type Object interface {
	RayCollision(r *Ray) bool
	RayManipulation(r *Ray) (*Color, []*Ray)
	Normalize()
}
