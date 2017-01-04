package main

import (
	"github.com/cloudspace/go-thrust/lib/bindings/session"
	"github.com/cloudspace/go-thrust/lib/commands"
	"github.com/cloudspace/go-thrust/lib/common"
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/provisioner"
)

func main() {
	/*
	   use basic setup
	*/
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	/*
			  Start of Advanced Session Tutorial.
			  We are going to set the Override value to true as ooposed to the false
			  used in basic_session. This will cause ThrustCore to try to invoke methods from us.

		    Look down below func main to find our session class.
	*/
	mysession := thrust.NewSession(false, true, "session_cache")

	mysession.SetInvokable(NewSimpleSession())
	/*
	   Modified basic_window, where we provide, a session argument
	   to NewWindow.
	*/
	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://breach.cc/",
		Session: mysession,
	})
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// In lieu of something like an http server, we need to lock this thread
	// in order to keep it open, and keep the process running.
	// Dont worry we use runtime.Gosched :)
	thrust.LockThread()
}

// SimpleSession must subscribe to the contract set forth by SessionInvokable interface

type SimpleSession struct {
	cookieStore map[string][]session.Cookie
}

func NewSimpleSession() (ss SimpleSession) {
	return ss
}

/*
For Simplicity type declarations
*/
func (ss SimpleSession) InvokeCookiesLoad(args *commands.CommandResponseArguments, s *session.Session) (cookies []session.Cookie) {
	common.Log.Print("InvokeCookiesLoad")
	cookies = make([]session.Cookie, 0)

	return cookies
}

func (ss SimpleSession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, s *session.Session) (cookies []session.Cookie) {
	common.Log.Print("InvokeCookiesLoadForKey")
	cookies = make([]session.Cookie, 0)

	return cookies
}

func (ss SimpleSession) InvokeCookiesFlush(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Print("InvokeCookiesFlush")
	return false
}

func (ss SimpleSession) InvokeCookiesAdd(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Print("InvokeCookiesAdd")
	return false
}

func (ss SimpleSession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Print("InvokeCookiesUpdateAccessTime")
	return false
}

func (ss SimpleSession) InvokeCookiesDelete(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Print("InvokeCookiesDelete")
	return false
}

func (ss SimpleSession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, s *session.Session) {
	common.Log.Print("InvokeCookieForceKeepSessionState")
}
