package main

import (
	"time"

	"github.com/sadasant/go-thrust/dispatcher"
	"github.com/sadasant/go-thrust/spawn"
	"github.com/sadasant/go-thrust/window"
)

func main() {
	spawn.SetBaseDirectory("./")
	spawn.Run()
	thrustWindow := window.NewWindow("http://breach.cc/", nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 5)
		thrustWindow.Close()
	}()
	// BLOCKING - Dont run before youve excuted all commands you want first
	dispatcher.RunLoop()
}
