package gnr

import (
	"math"
)

type FalloffFunc func(ir InteractionResult) InteractionResult
type FlatShader struct {
	Object
	FalloffFunc
}

func (fs FlatShader) RayInteraction(r Ray) []InteractionResult {
	irs := fs.Object.RayInteraction(r)
	irs = InteractionResultSlice(irs).SelectInteractionResult(fs.FalloffFunc)
	return irs
}

func NewLinearFalloffFunc(lightDir Vector3f) FalloffFunc {
	return func(ir InteractionResult) InteractionResult {
		cosAngle := VectorProduct(lightDir, ir.Normal) / lightDir.Magnitude() / ir.Normal.Magnitude()
		angle := math.Acos(cosAngle)
		ir.Color = VLerpCap(0, math.Pi, ColorBlack, ir.Color)(angle)
		return ir
	}
}
