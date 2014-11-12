package session

import "github.com/miketheprogrammer/go-thrust/commands"

type DummySession struct{}

func NewDummySession() (dummy *DummySession) {
	return &DummySession{}
}

func (d DummySession) InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session) {

}
