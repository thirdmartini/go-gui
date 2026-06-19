package ux

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

// ViewController contains any number of sub-widgets that can be displayed one at a time
//   typically used with a View
type ViewController struct {
	current string
	views   map[string]Drawable
	bg      color.Color
}

func (v *ViewController) Add(id string, widget Drawable) {
	v.views[id] = widget
}

func (v *ViewController) Remove(id string) {
	delete(v.views, id)
}

func (v *ViewController) Get(id string) Drawable {
	view, ok := v.views[id]
	if !ok {
		return nil
	}
	return view
}

func (v *ViewController) Show(id string) {
	view, ok := v.views[id]
	if !ok {
		return
	}
	view.Visible(true)
	v.current = id
}

func (v *ViewController) Current() (string, Drawable) {
	view, ok := v.views[v.current]
	if !ok {
		return "", nil
	}
	return v.current, view
}

func (v *ViewController) Draw(canvas canvas.Canvas) {
	if v.bg != nil {
		canvas.Clear(v.bg)
	}

	if view, ok := v.views[v.current]; ok {
		view.Draw(canvas)
	}
}

func (v *ViewController) SetBackground(color color.Color) {
	v.bg = color
}

func NewViewController() *ViewController {
	vc := &ViewController{
		views: make(map[string]Drawable),
	}
	return vc
}
