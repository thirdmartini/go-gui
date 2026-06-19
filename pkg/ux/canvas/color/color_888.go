package color

// Color888 stores the image a 5-6-5 encoded uint16
type Color888 struct {
	R uint8
	G uint8
	B uint8
}

// RGB565 return 16bit RGB 5-6-5 encoded color
func (c *Color888) RGB565() uint16 {
	return uint16(c.R/8)<<11 | uint16(c.G/4)<<5 | uint16(c.B/8)
}

// RGB return 0-1 color percentage
func (c *Color888) RGB() (float32, float32, float32) {
	return float32(c.R), float32(c.G), float32(c.B)
}

// RGBA satisfies ohe golang image.Color interface
func (c *Color888) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), uint32(255)
}

// RGB8 return 0-1 color percentage
func (c *Color888) RGB8() (uint8, uint8, uint8) {
	return c.R, c.G, c.B
}

type Palette888 struct {
}

func (m *Palette888) NewRGB8(r uint8, g uint8, b uint8) Color {
	return &Color888{
		R: r,
		G: g,
		B: b,
	}
}
