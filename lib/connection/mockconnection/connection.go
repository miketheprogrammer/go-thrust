package connection

import (
	"io"
	"os"
	"time"

	"github.com/miketheprogrammer/go-thrust/lib/commands"
	. "github.com/miketheprogrammer/go-thrust/lib/common"
)

// Single Connection
//var conn net.Conn
var StdIn io.WriteCloser
var StdOut io.ReadCloser

type In struct {
	Commands chan *commands.Command
	Quit     chan int
}
type Out struct {
	CommandResponses chan commands.CommandResponse
	Errors           chan error
}

var in In
var out Out

/*
Initializes threads with Channel Structs
Opens Connection
*/
func InitializeThreads() error {
	//c, err := net.Dial(proto, address)
	//conn = c

	in = In{
		Commands: make(chan *commands.Command),
		Quit:     make(chan int),
	}

	out = Out{
		CommandResponses: make(chan commands.CommandResponse),
		Errors:           make(chan error),
	}

	go Reader(&out, &in)
	go Writer(&out, &in)

	return nil
}

func GetOutputChannels() *Out {
	return &out
}

func GetInputChannels() *In {
	return &in
}

func GetCommunicationChannels() (*Out, *In) {
	return GetOutputChannels(), GetInputChannels()
}

func Reader(out *Out, in *In) {

	for {
		select {
		case quit := <-in.Quit:
			Log.Errorf("Connection Reader Received a Quit message from somewhere ... Exiting Now")
			os.Exit(quit)
		default:
			time.Sleep(time.Microsecond * 100)
			break
		}
	}

}

func Writer(out *Out, in *In) {
	var targetId uint = 0
	for {
		select {
		case command := <-in.Commands:
			ActionId += 1
			command.ID = ActionId
			response := commands.CommandResponse{}
			switch command.Action {
			case "create":
				targetId += 1
				response.ID = command.ID
				response.Action = "reply"
				response.Result = commands.ReplyResult{
					TargetID: targetId,
				}
				out.CommandResponses <- response
			case "call":
				response.ID = command.ID
				response.Action = "reply"
				out.CommandResponses <- response
			}
		}
		time.Sleep(time.Microsecond * 100)
	}
}
