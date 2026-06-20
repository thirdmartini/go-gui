package ux

import (
	"image"

	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

// Component is a shared class that contains common properties used by the draw code
type Component struct {
	R         image.Rectangle
	IsVisible bool
}

func (c *Component) X() int {
	return c.R.Min.X
}

func (c *Component) Y() int {
	return c.R.Min.Y
}

func (c *Component) W() int {
	return c.R.Dx()
}

func (c *Component) H() int {
	return c.R.Dy()
}

func (c *Component) Visible(show bool) {
	c.IsVisible = show
}

func (c *Component) Draw(_ canvas.Canvas) {
}

func (c *Component) OnEvent(_ *Event) bool {
	return false
}

func NewComponent(r image.Rectangle) *Component {
	return &Component{
		R: r,
	}
}

func NewComponentD(x, y, w, h int) *Component {
	return &Component{
		R: image.Rect(x, y, x+w, y+h),
	}
}
