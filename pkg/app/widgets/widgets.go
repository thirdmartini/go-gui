package widgets

import (
	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

type Meter interface {
	Draw(x, y int, c canvas.Canvas)
}
