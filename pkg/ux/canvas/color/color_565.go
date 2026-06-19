package color

// Color565 stores the image a 5-6-5 encoded uint16
type Color565 struct {
	value uint16
}

// Return 16bit RGB 5-6-5 encoded color
func (c *Color565) RGB565() uint16 {
	return c.value
}

// Return 0-1 color percentage
func (c *Color565) RGB() (float32, float32, float32) {
	return 0, 0, 0
}

// Return 0-1 color percentage
func (c *Color565) RGB8() (uint8, uint8, uint8) {
	r := ((c.value >> 11) & 0x1F) * 8
	g := ((c.value >> 5) & 0x3F) * 4
	b := (c.value & 0x1F) * 8

	return uint8(r), uint8(g), uint8(b)
}

// Return 0-1 color percentage
func (c *Color565) RGBA() (uint32, uint32, uint32, uint32) {
	r := ((c.value >> 11) & 0x1F) * 8
	g := ((c.value >> 5) & 0x3F) * 4
	b := (c.value & 0x1F) * 8

	return uint32(r), uint32(g), uint32(b), uint32(255)
}

type Palette565 struct {
}

func (m *Palette565) NewRGB8(r uint8, g uint8, b uint8) Color {
	return &Color565{
		value: uint16(r/8)<<11 | uint16(g/4)<<5 | uint16(b/8),
	}
}
