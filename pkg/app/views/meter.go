package views

import (
	"github.com/thirdmartini/gogui/pkg/app/widgets"
	"github.com/thirdmartini/gogui/pkg/ux"
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

type MeterPosition struct {
	x int
	y int
	m widgets.Meter
}

type MeterView struct {
	meters []MeterPosition
}

func (p *MeterView) OnEvent(event *ux.Event) bool {
	return true
}

func (p *MeterView) Visible(show bool) {
}

func (p *MeterView) Add(x, y int, m widgets.Meter) *MeterView {
	p.meters = append(p.meters, MeterPosition{x, y, m})
	return p
}

func (p *MeterView) Draw(canvas canvas.Canvas) {
	for _, mp := range p.meters {
		mp.m.Draw(mp.x, mp.y, canvas)
	}
}

func NewMeterView() *MeterView {
	p := &MeterView{}
	return p
}
