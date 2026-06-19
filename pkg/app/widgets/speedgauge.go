package widgets

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/thirdmartini/gogui/pkg/app/metrics"
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/themes"
)

type SpeedGauge struct {
	Title  string
	Width  int
	Height int

	LeftLabel   string
	LeftMetrics metrics.Gauge
	LeftMax     float64

	RightLabel   string
	RightMetrics metrics.Gauge
	RightMax     float64

	CenterMetric metrics.Gauge
	CenterMax    float64
}

func (p *SpeedGauge) Draw(x, y int, canvas canvas.Canvas) {
	cx := x + (p.Width / 2)
	cy := y + (p.Height / 2)

	radius := p.Height / 2 // radius 195

	canvas.DrawRoundedRect(cx+75, cy-50, 300, 100, 10, themes.ColorTextMuted, themes.ColorBackground)
	canvas.DrawRoundedRect(cx+75, cy-50, 300, 100, 10, themes.ColorBorder, nil)

	canvas.DrawRoundedRect(cx-375, cy-50, 300, 100, 10, themes.ColorTextMuted, themes.ColorBackground)
	canvas.DrawRoundedRect(cx-375, cy-50, 300, 100, 10, themes.ColorBorder, nil)

	// outer circle
	canvas.DrawCircle(cx, cy, radius+5, 5, themes.ColorBackground, themes.ColorBackground)

	// inner circle
	canvas.DrawCircle(cx, cy, radius-50, 5, themes.ColorBorder, nil)

	// draw center metric value
	rmp := int((p.CenterMetric.Value() / p.CenterMax) * 360)
	canvas.DrawArc(cx, cy, radius-70, 25, 90, 90+rmp, themes.ColorGraphAxis[2], nil)

	// Right Metric arc
	rv := 90 - int(180*(p.RightMetrics.Value()/p.RightMax))
	canvas.DrawArc(cx, cy, radius-20, 35, 90, rv, themes.ColorGraphAxis[0], nil)

	// Left Metric Arc
	lv := 90 + int(180*(p.LeftMetrics.Value()/p.LeftMax))
	canvas.DrawArc(cx, cy, radius-20, 35, 90, lv, themes.ColorGraphAxis[1], nil)

	// bar splitters
	for x := float64(0); x < 360; x += 10 {
		x0 := cx + int(float64(radius)*math.Cos(gg.Radians(x)))
		y0 := cy + int(float64(radius)*math.Sin(gg.Radians(x)))
		canvas.DrawLine(cx, cy, x0, y0, themes.ColorBackground)
	}

	canvas.DrawCircle(cx, cy, radius+5, 5, themes.ColorBorder, nil)

	// ticks
	canvas.DrawArc(cx, cy, radius, 10, -1, 1, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius, 10, 44, 46, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius-20, 50, 89, 91, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius, 10, 134, 136, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius, 10, 179, 181, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius, 10, 224, 226, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius-20, 50, 269, 271, themes.ColorBorder, nil)
	canvas.DrawArc(cx, cy, radius, 10, 314, 316, themes.ColorBorder, nil)

	// Text fields
	canvas.DrawRoundedRect(cx-75, cy-20, 150, 50, 10, themes.ColorBorder, themes.ColorBackground)
	canvas.DrawRoundedRect(cx-75, cy-20, 150, 50, 10, themes.ColorBorder, nil)
	canvas.DrawTextCentered(cx, cy, p.CenterMetric.String(), themes.Font(themes.FontLarge), themes.ColorTextPrimary)

	font := themes.Font(themes.FontHeader)
	canvas.DrawTextCentered(cx+275, cy-35, p.RightLabel, font, themes.ColorTextPrimary)
	canvas.DrawTextCentered(cx-275, cy-35, p.LeftLabel, font, themes.ColorTextPrimary)

	font = themes.Font(themes.FontLarge)
	canvas.DrawTextCentered(cx+275, cy, p.RightMetrics.String(), font, themes.ColorTextPrimary)
	canvas.DrawTextCentered(cx-275, cy, p.LeftMetrics.String(), font, themes.ColorTextPrimary)

}
