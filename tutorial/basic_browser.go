package main

import (
	"fmt"
	"os"
	"time"

	"github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

/*
Some docsss
*/
func main() {

	// Parses Flags
	InitLogger()

	connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()

	err := connection.InitializeThreads()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	out, in := connection.GetCommunicationChannels()

	thrustWindow := window.Window{
		Url: "http://breach.cc/",
	}

	// Calls to other methods after create are Queued until Create returns
	thrustWindow.Create(in, nil)
	thrustWindow.Show()

	thrustWindow.Maximize()

	dispatcher.RegisterHandler(func(c commands.CommandResponse) {
		thrustWindow.DispatchResponse(c)
	})

	for {
		select {
		case response := <-out.CommandResponses:
			dispatcher.Dispatch(response)
		default:
			break
		}
		time.Sleep(time.Microsecond * 10)

	}

}
