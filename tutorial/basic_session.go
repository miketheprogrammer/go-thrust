package main

import (
	"github.com/sadasant/go-thrust/dispatcher"
	"github.com/sadasant/go-thrust/session"
	"github.com/sadasant/go-thrust/spawn"
	"github.com/sadasant/go-thrust/window"
)

func main() {
	/*
	  use basic setup
	*/
	spawn.SetBaseDirectory("./")
	spawn.Run()

	/*
	   Start of Basic Session Tutorial area
	*/
	// arguments (incognito, useDisk)
	mysession := session.NewSession(false, false, "cache")
	//mysession.SetInvokable(*session.NewDummySession())
	/*
	  Modified basic_window, where we provide, a session argument
	  to NewWindow.
	*/
	thrustWindow := window.NewWindow("http://breach.cc/", mysession)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()
}
