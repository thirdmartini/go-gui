package ux

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

const (
	AlignDefault = 0x0
	AlignLeft    = 0x0
	AlignRight   = 0x1
	AlignTop     = 0x0
	AlignBottom  = 0x2
)

type Widget interface {
	EventHandler
	Drawable
}

type Drawable interface {
	//	Show(canvas canvas.Canvas)
	Draw(canvas canvas.Canvas)
	Visible(show bool)
}

type GUI struct {
	*ViewController
	overlay Drawable

	canvas canvas.Canvas
}

func (g *GUI) Display() {
	g.ViewController.Draw(g.canvas)
	if g.overlay != nil {
		g.overlay.Draw(g.canvas)
	}
	//g.canvas.Show()
}

func (g *GUI) ColorPalette() color.Palette {
	return g.canvas.ColorPalette()
}

func (g *GUI) SetOverlay(overlay Drawable) {
	g.overlay = overlay
}

func New(canvas canvas.Canvas) *GUI {
	g := &GUI{
		ViewController: NewViewController(),
		canvas:         canvas,
	}
	return g
}
