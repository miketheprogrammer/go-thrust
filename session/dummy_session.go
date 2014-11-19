package session

import (
	"github.com/sadasant/go-thrust/commands"
	"github.com/sadasant/go-thrust/common"
)

type DummySession struct{}

func NewDummySession() (dummy *DummySession) {
	return &DummySession{}
}

/*
For Simplicity type declarations
*/
func (ds DummySession) InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie) {
	common.Log.Debug("InvokeCookiesLoad")
	cookies = make([]Cookie, 0)

	return cookies
}

func (ds DummySession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie) {
	common.Log.Debug("InvokeCookiesLoadForKey")
	cookies = make([]Cookie, 0)

	return cookies
}

func (ds DummySession) InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session) bool {
	common.Log.Debug("InvokeCookiesFlush")
	return false
}

func (ds DummySession) InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session) bool {
	common.Log.Debug("InvokeCookiesAdd")
	return false
}

func (ds DummySession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session) bool {
	common.Log.Debug("InvokeCookiesUpdateAccessTime")
	return false
}

func (ds DummySession) InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session) bool {
	common.Log.Debug("InvokeCookiesDelete")
	return false
}

func (ds DummySession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session) {
	common.Log.Debug("InvokeCookieForceKeepSessionState")
}
