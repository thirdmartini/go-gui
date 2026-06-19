package kano

import (
	"fmt"

	"github.com/thirdmartini/gogui/pkg/ux"
)

var (
	ErrDoesNotExist = fmt.Errorf("does not exist")
)

type Kano struct {
	m *Motion
}

func (k *Kano) Listen(OnEvent func(ev *ux.Event)) error {
	k.m.OnGesture = func(g *Gesture) {
		switch g.Detail.Type {
		case GestureUp:
			OnEvent(ux.NewKeyPressEvent(ux.KeyPressUp))
		case GestureDown:
			OnEvent(ux.NewKeyPressEvent(ux.KeyPressDown))
		case GestureLeft:
			OnEvent(ux.NewKeyPressEvent(ux.KeyPressLeft))
		case GestureRight:
			OnEvent(ux.NewKeyPressEvent(ux.KeyPressRight))

		}
	}
	go k.m.SetMode(ModeGesture)
	return k.m.Run()
}

func Open(device string) (*Kano, error) {
	m := NewMotionSensorKitFromSerial(device)
	if m == nil {
		return nil, ErrDoesNotExist
	}
	k := &Kano{
		m: m,
	}

	return k, nil
}
