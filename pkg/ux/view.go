package ux

import (
	"fmt"

	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

// View is a single window that consumes the entire canvas of a device
//  it's used for small screens to present a single UI view
type View struct {
	visible    bool
	components map[string]Drawable
}

func (v *View) Get(id string) Drawable {
	widget, ok := v.components[id]
	if !ok {
		return nil
	}
	return widget
}

func (v *View) Add(id string, widget Drawable) {
	v.components[id] = widget
}

func (v *View) Remove(id string) {
	delete(v.components, id)
}

func (v *View) Draw(canvas canvas.Canvas) {
	for n, drawable := range v.components {
		fmt.Printf(" + View:Draw(%s)\n", n)
		drawable.Draw(canvas)
	}
}

func (v *View) Visible(show bool) {
	v.visible = show
}

func NewView() *View {
	view := &View{
		components: make(map[string]Drawable),
	}
	return view
}
