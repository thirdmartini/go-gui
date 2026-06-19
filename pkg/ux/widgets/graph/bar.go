package graph

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

type BarGraph struct {
	x int
	y int
	w int
	h int
	v int

	Color      color.Color
	FrameColor color.Color
}

func (s *BarGraph) Visible(show bool) {
}

func (s *BarGraph) PushValue(v float64) {

}

func (s *BarGraph) Draw(canvas canvas.Canvas) {
	v := s.v
	if v > s.w {
		v = s.w
	}
	// simple bar graph
	canvas.DrawRoundedRect(s.x, s.y, s.w, s.h, 5, s.Color, nil)
}

func NewBarGraph(x, y, w, h int, color color.Color) *BarGraph {
	return &BarGraph{
		Color: color,
		x:     x,
		y:     y,
		w:     w,
		h:     h,
		v:     0,
	}
}
