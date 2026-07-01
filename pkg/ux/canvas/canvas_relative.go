package canvas

import (
	"image"

	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/fonts"
)

var _ Canvas = (*RelativeCanvas)(nil)

// RelativeCanvas is a canvas that is always 0,0 based inside a parent canvas. ( creates a subcanvas )
type RelativeCanvas struct {
	canvas Canvas
	x      int
	y      int
	w      int
	h      int
}

func (cr *RelativeCanvas) Clear(color color.Color) {
	cr.canvas.Clear(color)
}

func (cr *RelativeCanvas) DrawPoint(x, y, r int, c color.Color) {
	cr.canvas.DrawPoint(x+cr.x, y+cr.y, r, c)
}

func (cr *RelativeCanvas) DrawCircle(x, y, r, w int, c color.Color, fill color.Color) {
	cr.canvas.DrawCircle(x+cr.x, y+cr.y, r, w, c, fill)
}

func (cr *RelativeCanvas) DrawText(x, y int, text string, font *fonts.Font, fg color.Color) {
	cr.canvas.DrawText(x+cr.x, y+cr.y, text, font, fg)
}

func (cr *RelativeCanvas) DrawTextWrapped(x, y, w, s int, text string, font *fonts.Font, fg color.Color) {
	cr.canvas.DrawTextWrapped(x+cr.x, y+cr.y, w, s, text, font, fg)
}

func (cr *RelativeCanvas) DrawTextCentered(x, y int, text string, font *fonts.Font, fg color.Color) {
	cr.canvas.DrawTextCentered(x+cr.x, y+cr.y, text, font, fg)
}

func (cr *RelativeCanvas) DrawArc(x, y, r, w int, start, stop int, color color.Color, fill color.Color) {
	cr.canvas.DrawArc(x+cr.x, y+cr.y, r, w, start, stop, color, fill)
}

func (cr *RelativeCanvas) DrawEllipticalArc(x, y, rx, xy int, start, stop int, color color.Color, fill color.Color) {
	cr.canvas.DrawEllipticalArc(x+cr.x, y+cr.y, rx, xy, start, stop, color, fill)
}

func (cr *RelativeCanvas) DrawImage(x, y int, im image.Image) {
	cr.canvas.DrawImage(x+cr.x, y+cr.y, im)
}

func (cr *RelativeCanvas) ColorPalette() color.Palette {
	return cr.canvas.ColorPalette()
}

func (cr *RelativeCanvas) ClipSet(x, y, w, h int) {
	cr.canvas.ClipSet(x+cr.x, y+cr.y, w, h)
}

func (cr *RelativeCanvas) ClipReset() {
	cr.canvas.ClipReset()
}

func (cr *RelativeCanvas) SetMargin(margin int) {
	cr.x += margin
	cr.y += margin
	cr.w -= margin * 2
	cr.h -= margin * 2
}

func (cr *RelativeCanvas) Width() int {
	return cr.w
}

func (cr *RelativeCanvas) Height() int {
	return cr.h
}

func (cr *RelativeCanvas) DrawTextBlock(x, y int, align uint8, block TextBlock) int {
	x = x + cr.x
	y = y + cr.y
	return cr.canvas.DrawTextBlock(x, y, align, block)
}

func (cr *RelativeCanvas) DrawLine(x1, y1, x2, y2 int, color color.Color) {
	x1 = x1 + cr.x
	y1 = y1 + cr.y
	x2 = x2 + cr.x
	y2 = y2 + cr.y
	cr.canvas.DrawLine(x1, y1, x2, y2, color)
}

func (cr *RelativeCanvas) DrawRect(x, y, w, h int, fill, outline color.Color) {
	x = x + cr.x
	y = y + cr.y
	cr.canvas.DrawRect(x, y, w, h, fill, outline)
}

func (cr *RelativeCanvas) DrawRoundedRect(x, y, w, h, r int, border color.Color, fill color.Color) {
	x = x + cr.x
	y = y + cr.y
	cr.canvas.DrawRoundedRect(x, y, w, h, r, border, fill)
}

func (cr *RelativeCanvas) Child(x, y, w, h int) *RelativeCanvas {
	return &RelativeCanvas{cr.canvas, x + cr.x, y + cr.y, w, h}
}

func NewRelativeCanvas(c Canvas, x, y, w, h int) *RelativeCanvas {
	return &RelativeCanvas{c, x, y, w, h}
}

func (cr *RelativeCanvas) DrawStart() {
	cr.canvas.DrawStart()
	cr.canvas.ClipSet(cr.x, cr.y, cr.w, cr.h)
}

func (cr *RelativeCanvas) DrawEnd() {
	cr.canvas.DrawEnd()
}
