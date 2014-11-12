package session

import (
	"fmt"

	. "github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
)

type Session struct {
	TargetID                 uint
	CookieStore              bool
	OffTheRecord             bool
	Ready                    bool
	CommandHistory           []*Command
	ResponseHistory          []*CommandResponse
	WaitingResponses         []*Command
	CommandQueue             []*Command
	SendChannel              *connection.In
	SessionOverrideInterface SessionInvokable
}

func NewSession(incognito, overrideDefaultSession bool, saveType string) *Session {
	session := Session{
		CookieStore:  overrideDefaultSession,
		OffTheRecord: incognito,
	}
	if overrideDefaultSession == true {
		session.SetInvokable(*NewDummySession())
	}
	command := Command{
		Action:     "create",
		ObjectType: "session",
		Args: CommandArguments{
			CookieStore:  session.CookieStore,
			OffTheRecord: session.OffTheRecord,
			Path:         saveType,
		},
	}
	session.SendChannel = connection.GetInputChannels()
	session.WaitingResponses = append(session.WaitingResponses, &command)
	session.Send(&command)
	dispatcher.RegisterHandler(session.DispatchResponse)
	return &session
}

func (session *Session) HandleInvoke(reply CommandResponse) {
	if reply.TargetID == session.TargetID {
		switch reply.Method {
		case "cookies_load":
			session.SessionOverrideInterface.InvokeCookiesLoad(&reply.Args, session)
		case "cookies_load_for_key":
			session.SessionOverrideInterface.InvokeCookiesLoadForKey(&reply.Args, session)
		case "cookies_flush":
			session.SessionOverrideInterface.InvokeCookiesFlush(&reply.Args, session)
		case "cookies_add":
			session.SessionOverrideInterface.InvokeCookiesAdd(&reply.Args, session)
		case "cookies_delete":
			session.SessionOverrideInterface.InvokeCookiesDelete(&reply.Args, session)
		case "cookies_update_access_time":
			session.SessionOverrideInterface.InvokeCookiesUpdateAccessTime(&reply.Args, session)
		case "cookies_force_keep_session_state":
			session.SessionOverrideInterface.InvokeCookieForceKeepSessionState(&reply.Args, session)
		}
	}
}

func (session *Session) HandleReply(reply CommandResponse) {
	fmt.Println(reply)
	for k, command := range session.WaitingResponses {
		if command.ID != reply.ID {
			continue
		}
		if command.ID == reply.ID {
			Log.Debug("Window(", session.TargetID, ")::Handling Reply::", reply)
			if len(session.WaitingResponses) > 1 {
				// Remove the element at index k
				session.WaitingResponses = session.WaitingResponses[:k+copy(session.WaitingResponses[k:], session.WaitingResponses[k+1:])]
			} else {
				// Just initialize to empty splice literal
				session.WaitingResponses = []*Command{}
			}
			fmt.Println("session", session.TargetID, command.Action, reply.Result.TargetID)
			if session.TargetID == 0 && command.Action == "create" {
				if reply.Result.TargetID != 0 {
					session.TargetID = reply.Result.TargetID
					Log.Debug("Session:: Received TargetID(", session.TargetID, ") :: Setting Ready State")
					session.Ready = true
				}
			}
		}
	}
}

func (session *Session) DispatchResponse(response CommandResponse) {
	switch response.Action {
	case "invoke":
		session.HandleInvoke(response)
	case "reply":
		session.HandleReply(response)

	}
}

func (session *Session) Send(command *Command) {
	session.SendChannel.Commands <- command
}

func (session *Session) SetInvokable(si SessionInvokable) {
	session.SessionOverrideInterface = si
}

/*
Methods prefixed with Invoke are methods that can be called by ThrustCore, this differs to our
standard call/reply, or event actions, since we are now the responder.
*/
/*
SessionInvokable is an interface designed to allow you to create your own Session Store.
Simple build a structure that supports these methods, and call session.SetInvokable(myInvokable)
*/
type SessionInvokable interface {
	InvokeCookiesLoad(args *CommandResponseArguments, session *Session)
	InvokeCookiesLoadForKey(args *CommandResponseArguments, session *Session)
	InvokeCookiesFlush(args *CommandResponseArguments, session *Session)
	InvokeCookiesAdd(args *CommandResponseArguments, session *Session)
	InvokeCookiesUpdateAccessTime(args *CommandResponseArguments, session *Session)
	InvokeCookiesDelete(args *CommandResponseArguments, session *Session)
	InvokeCookieForceKeepSessionState(args *CommandResponseArguments, session *Session)
}

type DummySession struct{}

func NewDummySession() (dummy *DummySession) {
	return &DummySession{}
}

func (d DummySession) InvokeCookiesLoad(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesLoadForKey(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesFlush(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesAdd(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesUpdateAccessTime(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookiesDelete(args *CommandResponseArguments, session *Session) {

}

func (d DummySession) InvokeCookieForceKeepSessionState(args *CommandResponseArguments, session *Session) {

}
