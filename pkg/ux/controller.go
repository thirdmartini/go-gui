package ux

import (
	"fmt"
	"time"

	"github.com/thirdmartini/gogui/pkg/ux/canvas"
)

type Controller struct {
	ui *GUI

	eventQueue chan *Event
	isSleeping bool
}

func (w *Controller) Render() {
	w.Sleep(false)
	w.ui.Display()
}

func (w *Controller) Sleep(sleep bool) {
	//w.driver.SetBacklight(!sleep)
}

func (w *Controller) PostEvent(ev *Event) {
	w.eventQueue <- ev
}

func (w *Controller) ListenEvents() {
	//app.OnCreate(w.ui)
	w.ui.Display()

	dur := time.Duration(time.Minute * 2)
	timeout := time.NewTimer(dur)
	w.isSleeping = false
	for {
		select {
		// Device interaction timeout. lock the device
		case <-timeout.C:
			fmt.Printf("[T1] timed out waiting for an event %v\n", w.isSleeping)
			//w.trySleepDevice()
			timeout.Reset(dur)

		// Screen update interval
		case <-time.After(time.Millisecond * 333):
			if !w.isSleeping {
				w.ui.Display()
			}

		// Handle UX or application events
		case ev, ok := <-w.eventQueue:
			if !ok {
				return // channel closed on us, exit the application
			}
			switch ev.Type {
			// Device interaction (button) events
			case EventTypeButton, EventTypeKey:
				timeout.Reset(dur)
				// dispatch to ux handler
				id, view := w.ui.ViewController.Current()
				fmt.Printf("Event %+v for %s\n", ev, id)

				dialog, ok := view.(EventHandler)
				if ok {
					if !dialog.OnEvent(ev) {
						w.ui.Display()
						return
					}
				}
				w.ui.Display()

			// Application events
			case EventTypeApplication:
				// dispatch to app event handler
				//w.appHandleEvent(ev)
			}
		}
	}
}

func NewController(canvas canvas.Canvas) *Controller {
	w := &Controller{
		//input:      make(chan int),
		ui:         New(canvas),
		eventQueue: make(chan *Event),
	}
	return w
}
