package window

import (
	"encoding/json"
	"fmt"
	"net"

	. "github.com/miketheprogrammer/thrust-go/commands"
	. "github.com/miketheprogrammer/thrust-go/common"
)

type Window struct {
	TargetID         int
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	CommandQueue     []*Command
	Url              string
	Title            string
	Conn             net.Conn
	Ready            bool
	Displayed        bool
}

func (w *Window) Create(conn net.Conn) {
	url := w.Url
	if len(url) == 0 {
		url = "http://google.com"
	}
	windowCreate := Command{
		Action:     "create",
		ObjectType: "window",
		Args: CommandArguments{
			RootUrl: url,
			Title:   "helloworld",
			Size: SizeHW{
				Width:  1024,
				Height: 768,
			},
		},
	}
	w.Send(&windowCreate, conn)
}

func (w *Window) IsTarget(targetId int) bool {
	return targetId == w.TargetID
}
func (w *Window) HandleError(reply CommandResponse, conn net.Conn) {

}

func (w *Window) HandleEvent(reply CommandResponse, conn net.Conn) {

}

func (w *Window) HandleReply(reply CommandResponse, conn net.Conn) {
	fmt.Println("Handling Response", reply)
	for k, v := range w.WaitingResponses {
		if v.ID != reply.ID {
			continue
		}

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
				fmt.Println("Received TargetID", "\nSetting Ready State")
				w.Ready = true
			}

			for i, _ := range w.CommandQueue {
				w.CommandQueue[i].TargetID = w.TargetID
				w.Send(w.CommandQueue[i], conn)
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

func (w *Window) DispatchResponse(reply CommandResponse, conn net.Conn) {
	switch reply.Action {
	case "event":
		w.HandleEvent(reply, conn)
	case "reply":
		w.HandleReply(reply, conn)
	}

}
func (w *Window) Send(command *Command, conn net.Conn) {
	ActionId += 1

	command.ID = ActionId

	fmt.Println(command)
	cmd, _ := json.Marshal(&command)
	fmt.Println("Writing", string(cmd), "\n", SOCKET_BOUNDARY)

	w.WaitingResponses = append(w.WaitingResponses, command)

	conn.Write(cmd)
	conn.Write([]byte("\n"))
	conn.Write([]byte(SOCKET_BOUNDARY))
}

func (w *Window) Call(command *Command, conn net.Conn) {
	command.Action = "call"
	command.TargetID = w.TargetID

	if w.Ready == false {
		w.CommandQueue = append(w.CommandQueue, command)
		return
	}
	w.Send(command, conn)
}

func (w *Window) Show(conn net.Conn) {
	command := Command{
		Method: "show",
	}

	w.Call(&command, conn)
}
