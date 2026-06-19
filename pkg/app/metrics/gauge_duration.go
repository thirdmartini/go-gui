package metrics

import (
	"time"
)

type DurationGauge struct {
	Format string
	value  time.Duration
}

func (g *DurationGauge) String() string {
	return g.value.String()
}

func (g *DurationGauge) Value() float64 {
	return g.value.Seconds()
}

func (g *DurationGauge) Set(d time.Duration) {
	g.value = d
}

func (g *DurationGauge) Axis(maxSamples int) ([]float64, float64) {
	return nil, 0
}
