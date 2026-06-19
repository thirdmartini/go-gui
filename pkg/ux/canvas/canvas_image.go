package canvas

import (
	"image"
	color2 "image/color"

	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

type RGBASurface struct {
	image   *image.RGBA
	palette color.Palette
}

func (s *RGBASurface) Width() int {
	return s.image.Bounds().Size().X
}

func (s *RGBASurface) Height() int {
	return s.image.Bounds().Size().Y
}

func (s *RGBASurface) Show() {
}

func (s *RGBASurface) ColorPalette() color.Palette {
	return s.palette
}

func (s *RGBASurface) Set(x, y int, color color.Color) {
	r, g, b := color.RGB8()
	s.image.Set(int(x), int(y), color2.RGBA{r, g, b, 255})
}

func NewImageCanvas(im *image.RGBA) *GenericCanvas {
	surface := &RGBASurface{
		image:   im,
		palette: &color.Palette888{},
	}

	return NewGenericCanvas(surface)
}
