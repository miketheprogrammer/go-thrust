package main

import (
	"fmt"
	"os"

	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

/*
Some docsss
*/
func main() {

	// Initialize the Logger
	InitLogger()

	// Spawn Thrust core and connect it to the connection package
	connection.StdOut, connection.StdIn = spawn.SpawnThrustCore("")

	// Initialize the Connection packages threads
	err := connection.InitializeThreads()

	if err != nil {
		fmt.Println(err)
		os.Exit(2) // Whatever error code you want to throw here.
	}

	// Store the communcation channels
	out, in := connection.GetCommunicationChannels()

	// Create the window struct with default values
	thrustWindow := window.Window{
		Url: "http://breach.cc/",
	}

	// Send a Create call to the Thrust core requesting the window to be created
	// Dont worry about return values, we handle that asynchronously behind the
	// scenes
	thrustWindow.Create(in, nil)

	// Show our new window.
	thrustWindow.Show()

	// Maximize our new window
	thrustWindow.Maximize()

	// Lets focus our window
	thrustWindow.Focus()
	// Register a handler for thrustWindow
	dispatcher.RegisterHandler(thrustWindow.DispatchResponse)
	// Start the main loop
	// Takes a *connection.Out as an argument.
	dispatcher.RunLoop(out)

}
