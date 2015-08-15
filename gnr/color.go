package gnr

var (
	ColorBlack = Color{0, 0, 0}
	ColorWhite = Color{1.0, 1.0, 1.0}
)

type Color struct {
	R, G, B float64
}

func (c Color) Normalize() {
	if c.R < 0 {
		c.R = 0
	}
	if c.G < 0 {
		c.G = 0
	}
	if c.B < 0 {
		c.B = 0
	}
	Sum := c.R + c.G + c.G
	if Sum == 0 {
		return
	}
	c.R /= Sum
	c.G /= Sum
	c.B /= Sum
}
