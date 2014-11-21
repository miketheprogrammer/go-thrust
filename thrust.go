package thrust

import (
	"errors"
	"runtime"

	"github.com/miketheprogrammer/go-thrust/lib/bindings/menu"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/session"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/commands"
	"github.com/miketheprogrammer/go-thrust/lib/common"
	"github.com/miketheprogrammer/go-thrust/lib/connection"
	"github.com/miketheprogrammer/go-thrust/lib/dispatcher"
	"github.com/miketheprogrammer/go-thrust/lib/spawn"
)

/*
Begin Generic Access and Binding Management Section.
*/

func NewWindow(url string, sess *session.Session) *window.Window {
	return window.NewWindow(url, sess)
}

func NewSession(incognito, overrideDefaultSession bool, path string) *session.Session {
	return session.NewSession(incognito, overrideDefaultSession, path)
}

func NewMenu() *menu.Menu {
	return menu.NewMenu()
}

func Start() {
	spawn.Run()
	go dispatcher.RunLoop()
}

func SetProvisioner(p spawn.Provisioner) {
	spawn.SetProvisioner(p)
}

func LockThread() {
	for {
		runtime.Gosched()
	}
}

func InitLogger() {
	common.InitLogger("")
}

func DisableLogger() {
	common.InitLogger("none")
}

func Exit() {
	connection.CleanExit()
}

func NewEventHandler(event string, fn interface{}) (ThrustEventHandler, error) {
	h := ThrustEventHandler{}
	h.Event = event
	h.Type = "event"
	err := h.SetHandleFunc(fn)
	dispatcher.RegisterHandler(h)
	return h, err
}

/**
Begin Thrust Handler Code.
**/
type ThrustHandler interface {
	Handle(cr commands.CommandResponse)
	Register()
	SetHandleFunc(fn interface{})
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
