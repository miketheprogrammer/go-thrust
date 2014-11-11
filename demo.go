package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	"github.com/miketheprogrammer/go-thrust/commands"
	. "github.com/miketheprogrammer/go-thrust/common"
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

	spawn.SetBaseDirectory(*baseDir)
	spawn.Run()

	mainSession := session.Session{}
	mainSession.Create()

	path, err := filepath.Abs("./public/index.html")

	if err != nil {
		panic(err)
	}
	thrustWindow := window.Window{
		Url: "file://" + path,
	}
	rootMenu := menu.Menu{}
	fileMenu := menu.Menu{}
	checkList := menu.Menu{}
	radioList := menu.Menu{}
	viewMenu := menu.Menu{}
	// Calls to other methods after create are Queued until Create returns
	thrustWindow.Create(nil)
	thrustWindow.Show()

	rootMenu.Create()
	//rootMenu.AddItem(2, "Root")

	fileMenu.Create()
	fileMenu.AddItem(3, "Open")

	fileMenu.AddItem(4, "Close")
	fileMenu.RegisterEventHandlerByCommandID(4, func(reply commands.CommandResponse, item *menu.MenuItem) {
		fmt.Println("Close button clicked.")
	})
	fileMenu.AddSeparator()

	checkList.Create()
	checkList.AddCheckItem(5, "Check 1")
	checkList.SetChecked(5, true)
	checkList.AddSeparator()
	checkList.AddCheckItem(6, "Check 2")
	checkList.SetChecked(6, true)
	checkList.SetEnabled(6, false)

	radioList.Create()
	radioList.AddRadioItem(7, "Radio 1-1", 1)
	radioList.AddRadioItem(8, "Radio 1-2", 1)
	radioList.AddSeparator()
	radioList.AddRadioItem(9, "Radio 2-1", 2)
	radioList.AddRadioItem(10, "Radio 2-2", 2)
	radioList.AddRadioItem(11, "Radio 2-3", 2)
	radioList.SetVisible(11, false)

	fileMenu.AddSubmenu(11, "CheckList", &checkList)
	fileMenu.AddSubmenu(12, "RadioList", &radioList)

	viewMenu.Create()
	layoutMenu := menu.Menu{}
	layoutMenu.Create()
	layoutStyleMenu := menu.Menu{}
	layoutStyleMenu.Create()
	layoutStyleMenu.AddRadioItem(13, "Horizontal", 3)
	layoutStyleMenu.AddRadioItem(14, "Vertical", 3)

	layoutMenu.AddSubmenu(15, "Styles", &layoutStyleMenu)
	viewMenu.AddSubmenu(16, "Layouts", &layoutMenu)
	rootMenu.AddSubmenu(17, "File", &fileMenu)
	rootMenu.AddSubmenu(18, "View", &viewMenu)

	rootMenu.SetApplicationMenu()
	rootMenu.Popup(&thrustWindow)

	thrustWindow.Maximize()
	thrustWindow.Focus()

	dispatcher.RegisterHandler(rootMenu.DispatchResponse)
	dispatcher.RegisterHandler(mainSession.DispatchResponse)

	// BLOCKING. This is not a thread,
	// It is meant to run on the main thread, and keep the process alive.
	dispatcher.RunLoop()
}
