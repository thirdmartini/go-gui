package vnc

import (
	"net"
)

const (
	v3    = "RFB 003.003\n"
	v7    = "RFB 003.007\n"
	v8    = "RFB 003.008\n"
	v8osx = "RFB 003.889\n"

	authError = 0
	authNone  = 1
	authVNC   = 2

	statusOK     = 0
	statusFailed = 1

	encodingRaw      = 0
	encodingCopyRect = 1

	// Client -> Server
	cmdSetPixelFormat           = 0
	cmdSetEncodings             = 2
	cmdFramebufferUpdateRequest = 3
	cmdKeyEvent                 = 4
	cmdPointerEvent             = 5
	cmdClientCutText            = 6

	// Server -> Client
	cmdFramebufferUpdate = 0
)

func NewServer(width, height int) *Server {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}
	conns := make(chan *Conn, 16)
	return &Server{
		width:  width,
		height: height,
		Conns:  conns,
		conns:  conns,
	}
}

type Server struct {
	width, height int
	conns         chan *Conn // read/write version of Conns

	// Conns is a channel of incoming connections.
	Conns <-chan *Conn
}

func (s *Server) Serve(ln net.Listener) error {
	for {
		c, err := ln.Accept()
		if err != nil {
			return err
		}
		conn := newConn(c, uint16(s.width), uint16(s.height))
		select {
		case s.conns <- conn:
		default:
			// client is behind; doesn't get this updated.
		}
		go conn.serve()
	}
	panic("unreachable")
}
