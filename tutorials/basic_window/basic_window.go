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

	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 5)
		thrustWindow.Close()
		thrust.Exit()
	}()

	// In lieu of something like an http server, we need to lock this thread
	// in order to keep it open, and keep the process running.
	// Dont worry we use runtime.Gosched :)
	thrust.LockThread()
}
