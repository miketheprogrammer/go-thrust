package main

import (
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/advanced_single_binary_distribution/provisioner"
	"github.com/pkg/profile"
)

func main() {

	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(provisioner.NewSingleBinaryThrustProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl:  "http://breach.cc/",
		HasFrame: true,
	})
	thrustWindow.Show()
	thrustWindow.Focus()

	// Lets do a window timeout
	go func() {
		// <-time.After(time.Second * 5)
		// thrustWindow.Close()
		// thrust.Exit()
	}()

	// In lieu of something like an http server, we need to lock this thread
	// in order to keep it open, and keep the process running.
	// Dont worry we use runtime.Gosched :)
	defer profile.Start().Stop()
	thrust.LockThread()
}
