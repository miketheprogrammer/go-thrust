package thrust

import (
	"errors"

	"github.com/miketheprogrammer/go-thrust/commands"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
)

/*
This is our top level package.
*/

var (
	// ApplicationName only functionally applies to OSX builds, otherwise it is only cosmetic
	ApplicationName = "Go Thrust"
)

/**
First Stab at event handlers
**/

type ThrustHandler interface {
	Handle(cr commands.CommandResponse)
	Register()
	SetHandleFunc(fn interface{})
}

func NewEventHandler(event string, fn interface{}) (ThrustEventHandler, error) {
	h := ThrustEventHandler{}
	h.Event = event
	h.Type = "event"
	err := h.SetHandleFunc(fn)
	dispatcher.RegisterHandler(h)
	return h, err
}

type ThrustEventHandler struct {
	Type    string
	Event   string
	Handler interface{}
}

func (teh ThrustEventHandler) Handle(cr commands.CommandResponse) {
	if cr.Action != "event" {
		return
	}
	if cr.Type != teh.Event && teh.Event != "*" {
		return
	}
	cr.Event.Type = cr.Type
	if fn, ok := teh.Handler.(func(commands.CommandResponse)); ok == true {
		fn(cr)
		return
	}
	if fn, ok := teh.Handler.(func(commands.EventResult)); ok == true {
		fn(cr.Event)
		return
	}
}

func (teh *ThrustEventHandler) SetHandleFunc(fn interface{}) error {
	if fn, ok := fn.(func(commands.CommandResponse)); ok == true {
		teh.Handler = fn
		return nil
	}
	if fn, ok := fn.(func(commands.EventResult)); ok == true {
		teh.Handler = fn
		return nil
	}

	return errors.New("Invalid Handler Definition")
}
