package main

import (
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/provisioner"
)

func main() {
	/*
	  use basic setup
	*/
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	/*
	   Start of Basic Session Tutorial area
	*/
	// arguments (incognito, useDisk)
	mysession := thrust.NewSession(false, false, "./cache")
	//mysession.SetInvokable(*session.NewDummySession())
	/*
	  Modified basic_window, where we provide, a session argument
	  to NewWindow.
	*/
	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://breach.cc/",
		Session: mysession,
	})
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	thrust.LockThread()
}
