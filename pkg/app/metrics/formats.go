package metrics

import (
	"fmt"
)

type Axis struct {
	Label string
	Value float64
}

type Gauge interface {
	String() string
	Value() float64
	Axis(maxSamples int) ([]float64, float64)
}

type FormatFunc func(v float64) string

func FormatWithString(sfmt string) FormatFunc {
	return func(v float64) string {
		return fmt.Sprintf(sfmt, v)
	}
}
