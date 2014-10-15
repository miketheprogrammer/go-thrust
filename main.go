package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var ActionId = 0

const (
	SOCKET_BOUNDARY = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
)

type SizeHW struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type CommandArguments struct {
	RootUrl string `json:"root_url,omitempty"`
	Title   string `json:"title,omitempty"`
	Size    SizeHW `json:"size,omitempty"`
	X       int    `json:"x,omitempty"`
	Y       int    `json:"y,omitempty"`
}
type Command struct {
	ID         int              `json:"_id"`
	Action     string           `json:"_action"`
	ObjectType string           `json:"_type,omitempty"`
	Method     string           `json:"_method,omitempty"`
	TargetID   int              `json:"_target,omitempty"`
	Args       CommandArguments `json:"_args"`
}

func (c Command) Send(conn net.Conn) {

}

//{"_action":"reply","_error":"","_id":1,"_result":{"_target":3}}

type ResponseResult struct {
	TargetID int `json:"_target,omitempty"`
}

type CommandResponse struct {
	Action string         `json:"_action,omitempty"`
	Error  string         `json:"_error,omitempty"`
	ID     int            `json:"_id,omitempty"`
	Result ResponseResult `json:"_result,omitempty"`
}

type ThrustApiObject interface {
	Create(conn net.Conn)
	Call(action string, command *Command, conn net.Conn)
	IsTarget(targetId int) bool
	Handle(reply CommandResponse)
}

type Window struct {
	TargetID         int
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	Url              string
	Title            string
	Conn             net.Conn
	Ready            bool
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
func (w *Window) HandleError(reply CommandResponse) {

}

func (w *Window) HandleEvent(reply CommandResponse) {

}

func (w *Window) HandleReply(reply CommandResponse) {
	fmt.Println("Handling Response", reply)
	for k, v := range w.WaitingResponses {
		if v.ID != reply.ID {
			continue
		}
		if len(w.WaitingResponses) > 1 {
			w.WaitingResponses = w.WaitingResponses[:k+copy(w.WaitingResponses[k:], w.WaitingResponses[k+1:])]
		} else {
			w.WaitingResponses = []*Command{}
		}

		if w.TargetID == 0 {
			//Assume we have a reply to action:create
			if reply.Result.TargetID != 0 {
				w.TargetID = reply.Result.TargetID
				fmt.Println("Received TargetID", "\bSetting Ready State")
				w.Ready = true
			}
		}

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

	w.Send(command, conn)
}

func (w *Window) Show(conn net.Conn) {
	command := Command{
		Method: "show",
	}

	w.Call(&command, conn)
}

func reader(r *bufio.Reader, ch chan CommandResponse) {
	// buf := make([]byte, 2048)
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			panic(err)
		}
		if !strings.Contains(line, SOCKET_BOUNDARY) {
			response := CommandResponse{}
			json.Unmarshal([]byte(line), &response)
			fmt.Println(response)
			ch <- response
		}

		fmt.Println("Got line", line)
	}
	// for {
	// 	n, err := r.Read(buf[:])
	// 	if err != nil {
	// 		return
	// 	}
	// 	println("Client got:", string(buf[0:n]))
	// }
}

func main() {
	addr := flag.String("socket", "", "unix socket where thrust is running")
	flag.Parse()
	fmt.Println(*addr)

	if len(*addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}
	conn, err := net.Dial("unix", *addr)

	defer conn.Close()

	if err != nil {
		os.Exit(2)
	}
	r := bufio.NewReader(conn)
	ch := make(chan CommandResponse)

	go reader(r, ch)

	window := Window{
		Conn: conn,
	}
	window.Create(conn)

	for {
		response := <-ch
		window.HandleReply(response)
		if window.Ready {
			window.Show(conn)
		}
		time.Sleep(time.Millisecond * 10000)

	}

}
