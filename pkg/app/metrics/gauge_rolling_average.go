package metrics

type RollingAverageGauge struct {
	SimpleGauge
	values []float64
}

func (g *RollingAverageGauge) Set(v float64) {
	if len(g.values) == cap(g.values) {

		sum := float64(0)
		for _, v := range g.values {
			sum += v
		}

		g.SimpleGauge.Set(sum / float64(len(g.values)))

		g.values = make([]float64, 0, cap(g.values))
	}
	g.values = append(g.values, v)
}

func NewRollingAverageGauge(size int) *RollupGauge {
	return &RollupGauge{
		values: make([]float64, 0, size),
	}
}
