package gauges

import (
	"github.com/thirdmartini/gogui/pkg/ux"
)

type Gauge interface {
	ux.Drawable
	SetValue(v int)
}
