package session

import (
	. "github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
)

type Session struct {
	TargetID         uint
	CookieStore      bool
	OffTheRecord     bool
	Ready            bool
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	CommandQueue     []*Command
}

func (session *Session) Create(sendChannel *connection.In) {
	command := commands.Commands{
		Action:     "create",
		ObjectType: "session",
		Args: CommandArguments{
			CookieStore:  session.CookieStore,
			OffTheRecord: session.OffTheRecord,
		},
	}
	session.WaitingResponses = append(session.WaitingResponses, &command)
	session.Send(&command)
}
func (session *Session) HandleError(reply CommandResponse) {

}

func (session *Session) HandleEvent(reply CommandResponse) {
	//Log.Info(("Window(", w.TargetID, ")::Handling Event::", reply))
}

func (session *Session) HandleReply(reply CommandResponse) {
	for k, command := range session.WaitingResponses {
		if command.ID != reply.ID {
			Log.Debug("Window(", session.TargetID, ")::Handling Reply::", reply)
			if len(session.WaitingResponses) > 1 {
				// Remove the element at index k
				session.WaitingResponses = session.WaitingResponses[:k+copy(session.WaitingResponses[k:], session.WaitingResponses[k+1:])]
			} else {
				// Just initialize to empty splice literal
				session.WaitingResponses = []*Command{}
			}

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

func (session *Session) InvokeCookiesLoad(response CommandResponse) {

}

func (session *Session) InvokeCookiesLoadForKey(response CommandResponse) {

}

func (session *Session) InvokeCookiesFlush(response CommandResponse) {

}

func (session *Session) Send(command *Command) {
	session.SendChannel.Commands <- command
}
