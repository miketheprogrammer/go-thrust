package main

import (
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/provisioner"
)

func main() {
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://breach.cc/",
	})
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	thrustWindow2 := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://google.com/",
	})
	thrustWindow2.Show()
	thrustWindow2.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	thrust.LockThread()
}
