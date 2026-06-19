package kano

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/google/uuid"
	"github.com/tarm/serial"
)

const (
	GestureUp    = "up"
	GestureDown  = "down"
	GestureLeft  = "left"
	GestureRight = "right"
)

type Mode string

const (
	ModeProximity Mode = "proximity"
	ModeGesture   Mode = "gesture"
)

type Motion struct {
	sync.Mutex
	rw  io.ReadWriter
	rpc map[string]*Response

	OnGesture   func(p *Gesture)
	OnProximity func(p *Proximity)
}

type Request struct {
	Type   string              `json:"type"`
	Id     string              `json:"id"`
	Method string              `json:"method"`
	Params []map[string]string `json:"params"`
}

type Message struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Error string `json:"error"`
	Value string `json:"value"`

	wq sync.WaitGroup `json:"_"`
}

func (r *Response) Done() {
	r.wq.Done()
}

func (r *Response) Wait() {
	r.wq.Wait()
}

type Proximity struct {
	Name   string `json:"name"`
	Detail struct {
		Proximity int `json:"proximity,omitempty"`
	} `json:"detail"`
}

type Gesture struct {
	Name   string `json:"name"`
	Detail struct {
		Type string `json:"type"`
	} `json:"detail"`
}

func (m *Motion) Queue(r *Request) *Response {
	data, err := json.Marshal(&r)
	data = append(data, '\r', '\n')

	response := &Response{}
	response.wq.Add(1)

	m.Lock()
	m.rpc[r.Id] = response
	m.Unlock()

	cnt, err := m.rw.Write(data)
	if err != nil {
		m.Lock()
		delete(m.rpc, r.Id)
		response.Error = "could not send"
		response.Done()
		m.Unlock()
	}

	fmt.Println(cnt, string(data))

	return response
}

func (m *Motion) QueueRequest(method string, params map[string]string) *Response {
	arr := make([]map[string]string, 0, 1)
	arr = append(arr, params)

	r := Request{
		Type:   "rpc-request",
		Id:     uuid.New().String(),
		Method: method,
		Params: arr,
	}

	response := m.Queue(&r)
	return response
}

func (m *Motion) SetMode(mode Mode) error {
	var resp *Response

	parms := make(map[string]string)
	parms["mode"] = string(mode)

	switch mode {
	case ModeProximity:
		resp = m.QueueRequest("set-mode", parms)
	case ModeGesture:
		resp = m.QueueRequest("set-mode", parms)

	default:
		return fmt.Errorf("invalid mode")
	}

	fmt.Println("Set mode:", mode)
	resp.Wait()
	fmt.Println("Set mode:", mode, "DONE")

	if resp.Error != "" {
		return fmt.Errorf(resp.Error)
	}
	return nil
}

func (m *Motion) onProximity(p *Proximity) {
	if m.OnProximity != nil {
		m.OnProximity(p)
	}
}

func (m *Motion) onGesture(p *Gesture) {
	if m.OnGesture != nil {
		m.OnGesture(p)
	}
}

func (m *Motion) loop() error {
	msg := Message{}

	scanner := bufio.NewScanner(m.rw)
	for scanner.Scan() {
		data := scanner.Bytes()

		fmt.Println(string(data))

		err := json.Unmarshal(data, &msg)
		if err != nil {
			continue
		}

		switch msg.Type {
		case "event":
			switch msg.Name {
			case "gesture":
				g := &Gesture{}

				err = json.Unmarshal(data, g)
				if err != nil {
					continue
				}
				m.onGesture(g)

			case "proximity":
				p := &Proximity{}
				err = json.Unmarshal(data, p)
				if err != nil {
					continue
				}
				m.onProximity(p)
			}

		case "rpc-response":
			m.Lock()
			k, ok := m.rpc[msg.Id]
			if ok {
				err = json.Unmarshal(data, &k)
				if err != nil {
					k.Error = "unable to unmarshal request"
				}

				delete(m.rpc, msg.Id)
				k.Done()
			}
			m.Unlock()
		}
	}

	data := make([]byte, 10)
	c, err := m.rw.Read(data)

	fmt.Println("exited", scanner.Err(), c, string(data), err)

	return fmt.Errorf("transport exception")
}

func (m *Motion) Run() error {
	return m.loop()
}

func NewMotionSensorKit(rw io.ReadWriter) *Motion {
	m := &Motion{
		rw:  rw,
		rpc: make(map[string]*Response),
	}

	return m
}

func NewMotionSensorKitFromSerial(device string) *Motion {
	config := &serial.Config{
		Name: device,
		Baud: 115200,
		// Timeout has this the sideffect of causing Scanner to think the serial port terminated
		//ReadTimeout: time.Millisecond * 5000,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		return nil
	}

	return NewMotionSensorKit(port)
}
