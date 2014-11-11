package main

import (
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

func main() {
	spawn.SetBaseDirectory("./vendor")
	spawn.Run()
	thrustWindow := window.Window{Url: "http://breach.cc/"}
	thrustWindow.Create(nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()
}
