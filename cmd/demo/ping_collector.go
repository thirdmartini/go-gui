package main

import (
	"runtime"
	"time"

	"github.com/go-ping/ping"
)

type PingCollector struct {
	label  string
	pinger *ping.Pinger
}

func (c PingCollector) Collect(label string, Collect func(label string, rtt time.Duration)) {
	c.pinger.OnRecv = func(pkt *ping.Packet) {
		Collect(label, pkt.Rtt)
	}
	c.pinger.Run()
}

func NewPingCollector(address string) (*PingCollector, error) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return nil, err
	}

	if runtime.GOOS == "linux" {
		pinger.SetPrivileged(true)
	}

	return &PingCollector{
		pinger: pinger,
	}, nil
}
