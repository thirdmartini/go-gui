package metrics

import (
	"container/list"
	"fmt"
	"sync"
)

type SimpleGauge struct {
	lock  sync.Mutex
	value float64

	format func(v float64) string

	list    list.List
	history int
}

func (g *SimpleGauge) WithFormat(f FormatFunc) {
	g.format = f
}

func (g *SimpleGauge) WithHistory(count int) {
	g.history = count
}

func (g *SimpleGauge) String() string {
	if g.format != nil {
		return g.format(g.Value())
	}

	return fmt.Sprintf("%v", g.Value())
}

func (g *SimpleGauge) Set(v float64) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.value = v
	if g.history != 0 {
		g.list.PushFront(v)
		if g.list.Len() > g.history {
			e := g.list.Back()
			g.list.Remove(e)
		}
	}
}

func (g *SimpleGauge) Value() float64 {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.value
}

func (g *SimpleGauge) Axis(maxSamples int) ([]float64, float64) {
	g.lock.Lock()
	defer g.lock.Unlock()

	m := float64(0)
	vals := make([]float64, 0, maxSamples)
	for e := g.list.Front(); e != nil; e = e.Next() {
		v := e.Value.(float64)

		if v > m {
			m = v
		}
		vals = append(vals, v)

		if len(vals) >= maxSamples {
			break
		}
	}
	return vals, m
}
