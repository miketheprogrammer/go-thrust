package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
	"github.com/miketheprogrammer/go-thrust/connection"
	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/menu"
	"github.com/miketheprogrammer/go-thrust/session"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/window"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)
	// Parses Flags
	baseDir := flag.String("basedir", usr.HomeDir, "Base Directory for Storing files.")
	InitLogger()

	connection.StdOut, connection.StdIn = spawn.SpawnThrustCore(*baseDir)

	err = connection.InitializeThreads()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	out, in := connection.GetCommunicationChannels()

	mainSession := session.Session{}
	mainSession.Create(in)

	thrustWindow := window.Window{
		Url: "http://breach.cc/",
	}
	rootMenu := menu.Menu{}
	fileMenu := menu.Menu{}
	checkList := menu.Menu{}
	radioList := menu.Menu{}
	viewMenu := menu.Menu{}
	// Calls to other methods after create are Queued until Create returns
	thrustWindow.Create(in, nil)
	thrustWindow.Show()

	rootMenu.Create(in)
	rootMenu.AddItem(2, "Root")

	fileMenu.Create(in)
	fileMenu.AddItem(3, "Open")

	fileMenu.AddItem(4, "Close")
	fileMenu.RegisterEventHandlerByCommandID(4, func(reply commands.CommandResponse, item *menu.MenuItem) {
		fmt.Println("Close button clicked.")
	})
	fileMenu.AddSeparator()

	checkList.Create(in)
	checkList.AddCheckItem(5, "Check 1")
	checkList.SetChecked(5, true)
	checkList.AddSeparator()
	checkList.AddCheckItem(6, "Check 2")
	checkList.SetChecked(6, true)
	checkList.SetEnabled(6, false)

	radioList.Create(in)
	radioList.AddRadioItem(7, "Radio 1-1", 1)
	radioList.AddRadioItem(8, "Radio 1-2", 1)
	radioList.AddSeparator()
	radioList.AddRadioItem(9, "Radio 2-1", 2)
	radioList.AddRadioItem(10, "Radio 2-2", 2)
	radioList.AddRadioItem(11, "Radio 2-3", 2)
	radioList.SetVisible(11, false)

	fileMenu.AddSubmenu(11, "CheckList", &checkList)
	fileMenu.AddSubmenu(12, "RadioList", &radioList)

	viewMenu.Create(in)
	layoutMenu := menu.Menu{}
	layoutMenu.Create(in)
	layoutStyleMenu := menu.Menu{}
	layoutStyleMenu.Create(in)
	layoutStyleMenu.AddRadioItem(13, "Horizontal", 3)
	layoutStyleMenu.AddRadioItem(14, "Vertical", 3)

	layoutMenu.AddSubmenu(15, "Styles", &layoutStyleMenu)
	viewMenu.AddSubmenu(16, "Layouts", &layoutMenu)
	rootMenu.AddSubmenu(17, "File", &fileMenu)
	rootMenu.AddSubmenu(18, "View", &viewMenu)

	rootMenu.SetApplicationMenu()

	thrustWindow.Maximize()

	thrustWindow.Focus()

	dispatcher.RegisterHandler(thrustWindow.DispatchResponse)
	dispatcher.RegisterHandler(rootMenu.DispatchResponse)
	dispatcher.RegisterHandler(mainSession.DispatchResponse)

	// BLOCKING. This is not a thread,
	// It is meant to run on the main thread, and keep the process alive.
	dispatcher.RunLoop(out)
}
