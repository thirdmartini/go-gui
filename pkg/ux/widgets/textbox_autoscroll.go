package widgets

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/fonts"
)

type ScrollingTextBox struct {
	*TextBox
	text   string
	offset uint64
}

func (t *ScrollingTextBox) SetText(text string) {
	t.text = text
	t.offset = 0
}

func (t *ScrollingTextBox) Draw(canvas canvas.Canvas) {
	if len(t.text) == 0 {
		t.TextBox.Draw(canvas)
		return
	}

	idx := t.offset % uint64(len(t.text))

	count := uint64(t.TextBox.Width()) / uint64(t.TextBox.Font.Width)

	t.TextBox.SetText(t.text[idx:])
	t.TextBox.Draw(canvas)
	t.offset++

	if idx+count > uint64(len(t.text)) {
		t.offset = 0
	}
}

func NewScrollingTextBox(x, y, w, h int, align uint8, text string, font *fonts.Font, color color.Color) *ScrollingTextBox {
	return &ScrollingTextBox{
		TextBox: NewTextBox(x, y, w, h, align, text, font, color),
		text:    text,
	}
}
