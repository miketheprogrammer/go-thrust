package main

import (
	"time"

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

	thrustWindow.OpenDevtools()
	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 10)
		thrustWindow.CloseDevtools()
	}()
	// BLOCKING - Dont run before youve excuted all commands you want first
	thrust.LockThread()
}
