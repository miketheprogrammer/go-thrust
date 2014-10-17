package dispatcher

import "github.com/miketheprogrammer/go-thrust/commands"

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
