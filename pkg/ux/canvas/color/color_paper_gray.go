package color

// Color565 stores the image a 5-6-5 encoded uint16
type ColorPaperGray struct {
	value uint8
}

// Return 16bit RGB 5-6-5 encoded color
func (c *ColorPaperGray) RGB565() uint16 {
	value := uint16(c.value/8)<<11 | uint16(c.value/4)<<5 | uint16(c.value/8)
	return value
}

// Return 0-1 color percentage
func (c *ColorPaperGray) RGB8() (uint8, uint8, uint8) {
	return c.value, c.value, c.value
}

// RGBA satisfies ohe golang image.Color interface
func (c *ColorPaperGray) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(c.value), uint32(c.value), uint32(c.value), uint32(255)
}

type PalettePaperGray struct {
}

func (m *PalettePaperGray) NewRGB8(r uint8, g uint8, b uint8) Color {
	return &ColorPaperGray{
		value: uint8((uint16(r) + uint16(g) + uint16(b)) / 3),
	}
}
