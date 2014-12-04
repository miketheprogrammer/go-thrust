package main

import (
	"github.com/miketheprogrammer/go-thrust"
	"github.com/miketheprogrammer/go-thrust/tutorials/provisioner"
)

func main() {
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	thrustWindow := thrust.NewWindow("http://breach.cc/", nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	thrustWindow2 := thrust.NewWindow("http://google.com/", nil)
	thrustWindow2.Show()
	thrustWindow2.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	thrust.LockThread()
}
