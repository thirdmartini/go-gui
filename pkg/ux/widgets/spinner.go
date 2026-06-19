package widgets

import (
	"math"

	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

type Spinner struct {
	Color color.Color
	x     int
	y     int
	r     int
	step  int
}

func (s *Spinner) Draw(canvas canvas.Canvas) {
	x0 := s.x + int(float64(s.r)*math.Cos(float64(s.step)*0.017453))
	y0 := s.y + int(float64(s.r)*math.Sin(float64(s.step)*0.017453))

	s.step += 10

	canvas.DrawLine(s.x, s.y, x0, y0, s.Color)
}

func NewSpinner(x, y, r int, color color.Color) *Spinner {
	return &Spinner{
		Color: color,
		x:     x,
		y:     y,
		r:     r,
		step:  0,
	}
}
