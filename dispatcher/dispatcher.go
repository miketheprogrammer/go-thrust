package dispatcher

import (
	"time"

	"github.com/miketheprogrammer/go-thrust/commands"
	"github.com/miketheprogrammer/go-thrust/connection"
)

type DispatcherHandleFunc func(commands.CommandResponse)

var registry []DispatcherHandleFunc

func RegisterHandler(f DispatcherHandleFunc) {
	registry = append(registry, f)
}

func Dispatch(command commands.CommandResponse) {
	for _, f := range registry {
		go f(command)
	}
}

func RunLoop(outChannels *connection.Out) {
	for {
		select {
		case response := <-outChannels.CommandResponses:
			Dispatch(response)
		default:
			break
		}
		time.Sleep(time.Microsecond * 10)
	}
}
