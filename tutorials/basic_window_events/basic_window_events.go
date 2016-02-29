package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudspace/go-thrust/lib/commands"
	"github.com/cloudspace/go-thrust/lib/connection"
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/provisioner"
)

/*
This tutorial teaches how to handle global events.
You can even use this to track menu/window/session etc. events,
if you store your bindings somewhere and track the ids.
Check package thrust for the acceptable handler definitions.
*/
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

	/*
	  Here we use an EventResult Callback, it provides us a little less data than a ComandResponse cb
	  however, its good for if we know we dont need to worry about targetids. Like understanding which window
	  was focuses, just that there is a window, that was focused.
	*/
	onfocus, err := thrust.NewEventHandler("focus", func(er commands.EventResult) {
		fmt.Println("Focus Event Occured")
	})
	fmt.Println(onfocus)
	if err != nil {
		fmt.Println(err)
		connection.CleanExit()
	}

	/*
	  Note blur does not seem to be triggered by unfocus
	  only by actually clicking the window, and then clicking somewhere else.
	*/
	onblur, err := thrust.NewEventHandler("blur", func(er commands.EventResult) {
		fmt.Println("Blur Event Occured")
	})
	fmt.Println(onblur)
	if err != nil {
		fmt.Println(err)
		connection.CleanExit()
	}

	/*
		Here we use a CommandResponse callback just because we can, it provides us more data
		than the EventResult callback
	*/
	onclose, err := thrust.NewEventHandler("closed", func(cr commands.EventResult) {
		fmt.Println("Close Event Occured")
	})
	fmt.Println(onclose)
	if err != nil {
		fmt.Println(err)
		connection.CleanExit()
	}

	/*
	  Lets say we just want to log all events
	*/
	onanything, err := thrust.NewEventHandler("*", func(cr commands.CommandResponse) {
		cr_marshaled, err := json.Marshal(cr)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("Event(%s) - Signaled by Command (%s)", cr.Type, cr_marshaled))
		}
	})

	fmt.Println(onanything)
	if err != nil {
		fmt.Println(err)
		connection.CleanExit()
	}

	time.AfterFunc(time.Second, func() {
		thrustWindow.Focus()
	})

	time.AfterFunc(time.Second*2, func() {
		thrustWindow.UnFocus()
	})

	time.AfterFunc(time.Second*3, func() {
		thrustWindow.Close()
	})

	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 5)
		connection.CleanExit()
	}()
	// BLOCKING - Dont run before youve excuted all commands you want first
	thrust.LockThread()

}
