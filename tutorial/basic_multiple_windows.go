package main

import (
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

func main() {
	spawn.SetBaseDirectory("./")
	spawn.Run(true)
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
