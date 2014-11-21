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

/*
Bindings
*/

/* NewWindow creates a new Window Binding */
func NewWindow(url string, sess *session.Session) *window.Window {
	return window.NewWindow(url, sess)
}

/* NewSession creates a new Session Binding */
func NewSession(incognito, overrideDefaultSession bool, path string) *session.Session {
	return session.NewSession(incognito, overrideDefaultSession, path)
}

/* NewMenu creates a new Menu Binding */
func NewMenu() *menu.Menu {
	return menu.NewMenu()
}

/*
Start spawns the thrust core executable, and begins the dispatcher loop in a go routine
*/
func Start() {
	spawn.Run()
	go dispatcher.RunLoop()
}

/*
SetProvisioner overrides the default Provisioner, the default provisioner downloads
Thrust-Core if Thrust-Core is not found.
It also does some other nifty things to configure your install (on darwin) for the ApplicationName you choose.
*/
func SetProvisioner(p spawn.Provisioner) {
	spawn.SetProvisioner(p)
}

/*
Use LockThread on the main thread in lieue of a webserver or some other service that holds the thread
This is primarily used when Thrust and just Thrust is what you are using, in that case lock the thread.
Otherwise, why dont you start an http server, and expose some websockets.
*/
func LockThread() {
	for {
		runtime.Gosched()
	}
}

/*
Initialize and Enable the internal *log.Logger.
*/
func InitLogger() {
	common.InitLogger("")
}

/*
Disable the internal *log.Logger instance
*/
func DisableLogger() {
	common.InitLogger("none")
}

/*
ALWAYS use this method instead of os.Exit()
This method will handle destroying the child process, and exiting as cleanly as possible.
*/
func Exit() {
	connection.CleanExit()
}

/*
Create a new EventHandler for a give event.
*/
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
