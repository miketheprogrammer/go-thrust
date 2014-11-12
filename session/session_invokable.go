package session

import "github.com/miketheprogrammer/go-thrust/commands"

/*
Methods prefixed with Invoke are methods that can be called by ThrustCore, this differs to our
standard call/reply, or event actions, since we are now the responder.
*/
/*
SessionInvokable is an interface designed to allow you to create your own Session Store.
Simple build a structure that supports these methods, and call session.SetInvokable(myInvokable)
*/
type SessionInvokable interface {
	InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session)
	InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session)
}
