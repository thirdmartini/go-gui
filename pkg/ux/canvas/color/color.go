package color

type Color interface {
	RGB565() uint16
	//	RGB() (float32, float32, float32)
	RGB8() (uint8, uint8, uint8)
	// RGBA Satisfy the golang color.color interface
	RGBA() (uint32, uint32, uint32, uint32)
}

// Interface to deal with canvass using different color formats
//   which are expensive to convert to/from on each pixel write
type Palette interface {
	NewRGB8(uint8, uint8, uint8) Color
}
