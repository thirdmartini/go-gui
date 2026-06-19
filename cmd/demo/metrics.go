package main

import (
	"fmt"

	"github.com/dustin/go-humanize"

	"github.com/thirdmartini/gogui/pkg/app/metrics"
)

type Metrics struct {
	EdgePingAvg10s   *metrics.RollupGauge
	EdgeBandwidthIn  metrics.SimpleGauge
	EdgeBandwidthOut metrics.SimpleGauge
}

func NewMetrics() *Metrics {
	m := &Metrics{}

	//m.EdgeBandwidthIn.FormatFunc = func(v float64) string { return fmt.Sprintf("%s/sec", humanize.Bytes(uint64(v))) }
	m.EdgeBandwidthIn.WithHistory(800)
	m.EdgeBandwidthIn.WithFormat(func(v float64) string {
		return fmt.Sprintf("%s/sec", humanize.Bytes(uint64(v)))
	})

	m.EdgeBandwidthOut.WithHistory(800)
	m.EdgeBandwidthOut.WithFormat(func(v float64) string {
		return fmt.Sprintf("%s/sec", humanize.Bytes(uint64(v)))
	})

	m.EdgePingAvg10s = metrics.NewRollupGauge(10)
	m.EdgePingAvg10s.WithHistory(800)
	m.EdgePingAvg10s.WithFormat(func(v float64) string {
		return fmt.Sprintf("%7.3f ms", v)
	})

	return m
}
