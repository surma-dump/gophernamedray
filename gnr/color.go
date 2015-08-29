package gnr

import (
	"image/color"
)

var (
	ColorBlack   = &Vector3f{0, 0, 0}
	ColorWhite   = &Vector3f{1.0, 1.0, 1.0}
	ColorRed     = &Vector3f{1.0, 0.0, 0.0}
	ColorYellow  = &Vector3f{1.0, 1.0, 0.0}
	ColorGreen   = &Vector3f{0.0, 1.0, 0.0}
	ColorCyan    = &Vector3f{0.0, 1.0, 1.0}
	ColorBlue    = &Vector3f{0.0, 0.0, 1.0}
	ColorMagenta = &Vector3f{1.0, 0.0, 1.0}
)

var fColor = LerpCap(0, 1, 0, 255)

func (v *Vector3f) ToColor() color.Color {
	return color.RGBA{
		R: uint8(fColor(v.X)),
		G: uint8(fColor(v.Y)),
		B: uint8(fColor(v.Z)),
		A: 255,
	}
}
