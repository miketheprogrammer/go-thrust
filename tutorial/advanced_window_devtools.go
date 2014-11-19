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

	thrustWindow.OpenDevtools()
	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 10)
		thrustWindow.CloseDevtools()
	}()
	// BLOCKING - Dont run before youve excuted all commands you want first
	dispatcher.RunLoop()
}
