package window

import (
	"fmt"
	"time"

	"github.com/miketheprogrammer/go-thrust"
	. "github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/session"
)

type Window struct {
	TargetID         uint
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	CommandQueue     []*Command
	Url              string
	Title            string
	Ready            bool
	Displayed        bool
	SendChannel      *connection.In `json:"-"`
}

func NewWindow(url string, sess *session.Session) *Window {
	w := Window{}
	w.Url = url
	if len(w.Url) == 0 {
		w.Url = "http://google.com"
	}
	_, sendChannel := connection.GetCommunicationChannels()

	windowCreate := Command{
		Action:     "create",
		ObjectType: "window",
		Args: CommandArguments{
			RootUrl: w.Url,
			Title:   thrust.ApplicationName,
			Size: SizeHW{
				Width:  1024,
				Height: 768,
			},
		},
	}
	dispatcher.RegisterHandler(w.DispatchResponse)
	if sess == nil {
		w.SetSendChannel(sendChannel)
		w.WaitingResponses = append(w.WaitingResponses, &windowCreate)
		w.Send(&windowCreate)
	} else {
		go func() {
			for {
				if sess.TargetID != 0 {
					fmt.Println("sess", sess.TargetID)
					windowCreate.Args.SessionID = sess.TargetID
					w.SetSendChannel(sendChannel)
					w.WaitingResponses = append(w.WaitingResponses, &windowCreate)
					w.Send(&windowCreate)
					return
				}
				time.Sleep(time.Microsecond * 10)
			}
		}()
	}
	return &w
}

func (w *Window) SetSendChannel(sendChannel *connection.In) {
	w.SendChannel = sendChannel
}

func (w *Window) IsTarget(targetId uint) bool {
	return targetId == w.TargetID
}
func (w *Window) HandleError(reply CommandResponse) {

}

func (w *Window) HandleEvent(reply CommandResponse) {
	//Log.Info(("Window(", w.TargetID, ")::Handling Event::", reply))
}

func (w *Window) HandleReply(reply CommandResponse) {
	for k, v := range w.WaitingResponses {
		if v.ID != reply.ID {
			continue
		}
		Log.Debug("Window(", w.TargetID, ")::Handling Reply::", reply)
		if len(w.WaitingResponses) > 1 {
			// Remove the element at index k
			w.WaitingResponses = w.WaitingResponses[:k+copy(w.WaitingResponses[k:], w.WaitingResponses[k+1:])]
		} else {
			// Just initialize to empty splice literal
			w.WaitingResponses = []*Command{}
		}

		// If we dont already have a TargetID then we accept a create action
		if w.TargetID == 0 && v.Action == "create" {
			if reply.Result.TargetID != 0 {
				w.TargetID = reply.Result.TargetID
				Log.Debug("Received TargetID", "\nSetting Ready State")
				w.Ready = true
			}

			for i, _ := range w.CommandQueue {
				w.CommandQueue[i].TargetID = w.TargetID
				w.Send(w.CommandQueue[i])
			}
			// Reinitialize empty command queue, and allow gc.
			w.CommandQueue = []*Command{}

			return
		}

		if v.Action == "call" && v.Method == "show" {
			w.Displayed = true
		}

	}
}

func (w *Window) DispatchResponse(reply CommandResponse) {

	switch reply.Action {
	case "event":
		w.HandleEvent(reply)
	case "reply":
		w.HandleReply(reply)
	}

}
func (w *Window) Send(command *Command) {

	w.SendChannel.Commands <- command
}

func (w *Window) Call(command *Command) {
	command.Action = "call"
	command.TargetID = w.TargetID
	if w.Ready == false {
		w.CommandQueue = append(w.CommandQueue, command)
		return
	}
	w.Send(command)
}

func (w *Window) CallWhenReady(command *Command) {
	w.WaitingResponses = append(w.WaitingResponses, command)
	go func() {
		for {
			if w.Ready {
				w.Call(command)
				return
			}
			time.Sleep(time.Microsecond * 100)
		}
	}()
}

func (w *Window) CallWhenDisplayed(command *Command) {
	w.WaitingResponses = append(w.WaitingResponses, command)
	go func() {
		for {
			if w.Displayed {
				w.Call(command)
				return
			}
			time.Sleep(time.Microsecond * 100)
		}
	}()
}

func (w *Window) Show() {
	command := Command{
		Method: "show",
	}

	w.CallWhenReady(&command)
}

func (w *Window) Maximize() {
	command := Command{
		Method: "maximize",
	}
	w.CallWhenDisplayed(&command)
}

func (w *Window) UnMaximize() {
	command := Command{
		Method: "unmaximize",
	}
	w.CallWhenDisplayed(&command)
}

func (w *Window) Minimize() {
	command := Command{
		Method: "minmize",
	}
	w.CallWhenDisplayed(&command)
}

func (w *Window) Restore() {
	command := Command{
		Method: "restore",
	}
	w.CallWhenDisplayed(&command)
}

func (w *Window) Focus() {
	command := Command{
		Method: "focus",
		Args: CommandArguments{
			Focus: true,
		},
	}

	w.CallWhenDisplayed(&command)
}

func (w *Window) UnFocus() {
	command := Command{
		Method: "show",
		Args: CommandArguments{
			Focus: false,
		},
	}

	w.CallWhenDisplayed(&command)
}
