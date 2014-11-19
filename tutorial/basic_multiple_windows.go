package main

import (
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

	thrustWindow2 := window.NewWindow("http://google.com/", nil)
	thrustWindow2.Show()
	thrustWindow2.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()
}
