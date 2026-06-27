package canvas

import (
	"image"

	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"

	"github.com/thirdmartini/gogui/pkg/ux/canvas/fonts"
)

const (
	TextAlignTop    = 0x0
	TextAlignBottom = 0x2
)

type TextBlock struct {
	Content string
	Color   color.Color
	Font    *fonts.Font
}

type Surface interface {
	Set(x, y int, c color.Color)
	Width() int
	Height() int
	ColorPalette() color.Palette
	Show()
}

type Canvas interface {
	//SetColor(c color.Color)
	//SetFillColor(c color.Color)
	//SetFont(font *fonts.Font)

	Clear(color color.Color)

	DrawPoint(x, y, r int, c color.Color)
	DrawLine(x1, y1, x2, y2 int, c color.Color)

	DrawCircle(x, y, r, w int, c color.Color, fill color.Color)
	DrawRect(x, y, w, h int, c color.Color, fill color.Color)
	DrawRoundedRect(x, y, w, h, r int, c color.Color, fill color.Color)
	DrawText(x, y int, text string, font *fonts.Font, fg color.Color)
	DrawTextWrapped(x, y, w, s int, text string, font *fonts.Font, fg color.Color)
	DrawTextCentered(x, y int, text string, font *fonts.Font, fg color.Color)

	DrawTextBlock(x, y int, valign uint8, text TextBlock) int

	DrawArc(x, y, r, w int, start, stop int, color color.Color, fill color.Color)
	DrawEllipticalArc(x, y, rx, xy int, start, stop int, color color.Color, fill color.Color)

	DrawImage(x, y int, im image.Image)

	//Show()
	ColorPalette() color.Palette
	//Invalidate()

	ClipSet(x, y, w, h int)
	ClipReset()
}
