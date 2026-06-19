package ux

import (
	"time"
)

type ApplicationController interface {
	OnEvent(ev *Event) bool
	OnRepaint()
}

type Application struct {
	eventQueue chan *Event
	ctrl       ApplicationController
}

func (app *Application) PostEvent(ev *Event) {
	app.eventQueue <- ev
}

func (app *Application) Run(ctrl ApplicationController) {
	app.ctrl = ctrl
	for {
		select {
		// Screen update interval
		case <-time.After(time.Second):
			app.ctrl.OnRepaint()

		// Handle UX or application events
		case ev, ok := <-app.eventQueue:
			if !ok {
				return // channel closed on us, exit the application
			}
			if app.ctrl.OnEvent(ev) {
				app.ctrl.OnRepaint()
			}
		}
	}
}

func NewApplication() *Application {
	return &Application{
		eventQueue: make(chan *Event),
	}
}
