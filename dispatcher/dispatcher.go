package dispatcher

import (
	"runtime"

	"github.com/miketheprogrammer/go-thrust/commands"
	"github.com/miketheprogrammer/go-thrust/connection"
)

type HandleFunc func(commands.CommandResponse)

var registry []HandleFunc

/*
RegisterHandler registers a HandleFunc f to receive a CommandResponse when one is sent to the system.
*/
func RegisterHandler(f HandleFunc) {
	registry = append(registry, f)
}

/*
Dispatch dispatches a CommandResponse to every handler in the registry
*/
func Dispatch(command commands.CommandResponse) {
	for _, f := range registry {
		go f(command)
	}
}

/*
RunLoop starts a loop that receives CommandResponses and dispatches them.
This is a helper method, but you could just implement your own, if you only
need this loop to be the blocking loop.
For Instance, in a HTTP Server setting, you might want to run this as a
goroutine and then let the servers blocking handler keep the process open.
As long as there are commands in the channel, this loop will dispatch as fast
as possible

Also, watch for Signals and handle.
*/
func RunLoop() {
	outChannels := connection.GetOutputChannels()
	defer connection.Clean()

	for {
		Run(outChannels)
		runtime.Gosched()
	}

}

func Run(outChannels *connection.Out) {
	response := <-outChannels.CommandResponses
	Dispatch(response)
}
