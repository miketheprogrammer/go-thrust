package main

import (
	"flag"
	"fmt"

	"github.com/sadasant/go-thrust/dispatcher"
	"github.com/sadasant/go-thrust/spawn"
	"github.com/sadasant/go-thrust/window"
)

var (
	host = flag.String("host", "0.0.0.0", "IP address to bind to")
	port = flag.Int("port", 8000, "TCP port to listen on")
)

func main() {
	flag.Parse()
	spawn.SetBaseDirectory("./")
	spawn.Run()
	thrustWindow := window.NewWindow(fmt.Sprintf("http://127.0.0.1:%d", *port), nil)
	thrustWindow.Show()
	thrustWindow.Focus()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()

}
