package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
)

// Single Connection
//var conn net.Conn
var Stdin io.WriteCloser
var Stdout io.ReadCloser

type In struct {
	Commands         chan *commands.Command
	CommandResponses chan *commands.CommandResponse
	Quit             chan int
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
func InitializeThreads() {
	//c, err := net.Dial(proto, address)
	//conn = c

	in = In{
		Commands:         make(chan *commands.Command),
		CommandResponses: make(chan *commands.CommandResponse),
		Quit:             make(chan int),
	}

	out = Out{
		CommandResponses: make(chan commands.CommandResponse),
		Errors:           make(chan error),
	}

	go Reader(&out, &in)
	go Writer(&out, &in)

	return
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

	r := bufio.NewReader(Stdout)
	defer Stdin.Close()
	for {
		select {
		case quit := <-in.Quit:
			Log.Errorf("Connection Reader Received a Quit message from somewhere ... Exiting Now")
			os.Exit(quit)
		default:
			//a := <-in.Quit
			//fmt.Println(a)
			line, err := r.ReadString(byte('\n'))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			Log.Debug("SOCKET::Line", line)
			if !strings.Contains(line, SOCKET_BOUNDARY) {
				response := commands.CommandResponse{}
				json.Unmarshal([]byte(line), &response)
				//Log.Debug(response)
				out.CommandResponses <- response
			}

		}
		time.Sleep(time.Microsecond * 100)

	}

}

func Writer(out *Out, in *In) {
	for {
		select {
		case response := <-in.CommandResponses:
			Log.Info("CommandResponse Marhshaling.")
			cmd, _ := json.Marshal(response)
			Log.Debug("Writing RESPONSE", string(cmd), "\n", SOCKET_BOUNDARY)

			Stdin.Write(cmd)
			Stdin.Write([]byte("\n"))
			Stdin.Write([]byte(SOCKET_BOUNDARY))
			Stdin.Write([]byte("\n"))
		case command := <-in.Commands:
			ActionId += 1
			command.ID = ActionId

			//fmt.Println(command)
			cmd, _ := json.Marshal(command)
			Log.Debug("Writing", string(cmd), "\n", SOCKET_BOUNDARY)

			Stdin.Write(cmd)
			Stdin.Write([]byte("\n"))
			Stdin.Write([]byte(SOCKET_BOUNDARY))
			Stdin.Write([]byte("\n"))
		}

		time.Sleep(time.Microsecond * 100)
	}
}
