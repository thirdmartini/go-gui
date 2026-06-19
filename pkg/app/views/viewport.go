package views

import (
	"image"

	"github.com/thirdmartini/gogui/pkg/drivers/display"
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

type DisplayCanvas struct {
	canvas.Canvas
	im      *image.RGBA
	display display.Display
}

func (dc *DisplayCanvas) Show() {
	dc.display.Draw(dc.im)
}

func NewDisplayCanvas(d display.Display) *DisplayCanvas {
	pt := d.Size()

	im := image.NewRGBA(image.Rect(0, 0, pt.X, pt.Y))

	dc := &DisplayCanvas{
		Canvas:  canvas.NewGGCanvas(im),
		im:      im,
		display: d,
	}
	return dc
}
