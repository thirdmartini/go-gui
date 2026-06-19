package gauges

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

type SimpleHBar struct {
	Color color.Color
	x     int
	y     int
	w     int
	h     int

	v int
}

func (s *SimpleHBar) Visible(show bool) {
}

func (s *SimpleHBar) SetValue(v int) {
	s.v = v
}

func (s *SimpleHBar) Draw(canvas canvas.Canvas) {
	v := s.v
	if v > s.w {
		v = s.w
	}
	// simple bar graph
	canvas.DrawRect(s.x, s.y, v, s.h, s.Color, s.Color)
}

func NewSimpleHBar(x, y, h int, color color.Color) *SimpleHBar {
	return &SimpleHBar{
		Color: color,
		x:     x,
		y:     y,
		h:     h,
		v:     0,
	}
}
