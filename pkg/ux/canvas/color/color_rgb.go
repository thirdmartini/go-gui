package color

/*
type ColorRGB struct {
	r float32
	g float32
	b float32
}

func (c *ColorRGB) fold16(max uint16, scale float32) uint16 {
	v := uint16(float32(max) * scale)
	if v > max {
		v = max
	}
	return v
}

func (c *ColorRGB) RGB565() uint16 {
	var v uint16
	v = c.fold16(31, c.r)<<11 | c.fold16(63, c.g)<<5 | c.fold16(31, c.b)
	return v
}

func (c *ColorRGB) RGB() (float32, float32, float32) {
	return c.r, c.g, c.b
}

func (c *ColorRGB) RGB8() (uint8, uint8, uint8) {
	return uint8(256 * c.r), uint8(256 * c.g), uint8(256 * c.b)
}



func (c *ColorRGB) Set(r, g, b float32) *ColorRGB {
	c.r = r
	c.g = g
	c.b = b
	return c
}

func NewRGB(r, g, b float32) *ColorRGB {
	c := &ColorRGB{
		r: r,
		g: g,
		b: b,
	}
	return c
}
*/
