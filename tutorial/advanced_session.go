package main

import (
	"github.com/miketheprogrammer/go-thrust/commands"
	"github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/session"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

func main() {
	/*
	   use basic setup
	*/
	spawn.SetBaseDirectory("./")
	spawn.Run(true)

	/*
			  Start of Advanced Session Tutorial.
			  We are going to set the Override value to true as ooposed to the false
			  used in basic_session. This will cause ThrustCore to try to invoke methods from us.

		    Look down below func main to find our session class.
	*/
	mysession := session.NewSession(false, true, "cache")

	mysession.SetInvokable(NewSimpleSession())
	/*
	   Modified basic_window, where we provide, a session argument
	   to NewWindow.
	*/
	thrustWindow := window.NewWindow("http://breach.cc/", mysession)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()
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
	common.Log.Debug("InvokeCookiesLoad")
	cookies = make([]session.Cookie, 0)

	return cookies
}

func (ss SimpleSession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, s *session.Session) (cookies []session.Cookie) {
	common.Log.Debug("InvokeCookiesLoadForKey")
	cookies = make([]session.Cookie, 0)

	return cookies
}

func (ss SimpleSession) InvokeCookiesFlush(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Debug("InvokeCookiesFlush")
	return false
}

func (ss SimpleSession) InvokeCookiesAdd(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Debug("InvokeCookiesAdd")
	return false
}

func (ss SimpleSession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Debug("InvokeCookiesUpdateAccessTime")
	return false
}

func (ss SimpleSession) InvokeCookiesDelete(args *commands.CommandResponseArguments, s *session.Session) bool {
	common.Log.Debug("InvokeCookiesDelete")
	return false
}

func (ss SimpleSession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, s *session.Session) {
	common.Log.Debug("InvokeCookieForceKeepSessionState")
}
