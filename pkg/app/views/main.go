package views

import (
	"github.com/thirdmartini/gogui/pkg/ux"
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/fonts"
)

type MainView struct {
	Font   *fonts.Font
	Font18 *fonts.Font
}

func (p *MainView) OnEvent(event *ux.Event) bool {
	switch event.Type {
	case ux.EventTypeButton:
		return true
	}
	return true
}

func (p *MainView) Draw(canvas canvas.Canvas) {
	canvas.DrawText(100, 100, "DEMO SCREEN", p.Font, canvas.ColorPalette().NewRGB8(0, 0, 0))
	canvas.DrawText(100, 140, "Up/Down Keyboard Arrow to change screens", p.Font18, canvas.ColorPalette().NewRGB8(0, 0, 0))
}

func (p *MainView) Visible(show bool) {
}

func NewMainView() *MainView {
	font, err := fonts.Load("assets/light/default.ttf", 30)
	if err != nil {
		panic(err)
	}

	font18, err := fonts.Load("assets/light/default.ttf", 18)
	if err != nil {
		panic(err)
	}

	p := &MainView{
		Font:   font,
		Font18: font18,
	}
	return p
}
