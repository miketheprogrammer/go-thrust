package main

import (
	"fmt"

	"github.com/sadasant/go-thrust/commands"
	"github.com/sadasant/go-thrust/dispatcher"
	"github.com/sadasant/go-thrust/menu"
	"github.com/sadasant/go-thrust/spawn"
	"github.com/sadasant/go-thrust/window"
)

func main() {
	spawn.SetBaseDirectory("./")
	spawn.Run()
	thrustWindow := window.NewWindow("http://breach.cc/", nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// make our top menus
	//applicationMenu, is essentially the menu bar
	applicationMenu := menu.NewMenu()
	//applicationMenuRoot is the first menu, on darwin this is always named the name of your application.
	applicationMenuRoot := menu.NewMenu()
	//File menu is our second menu
	fileMenu := menu.NewMenu()

	// Lets build our root menu.
	// the first argument to AddItem is a CommandID
	// A CommandID is used by Thrust Core to communicate back results and events.
	applicationMenuRoot.AddItem(1, "About")
	applicationMenuRoot.RegisterEventHandlerByCommandID(1,
		func(reply commands.CommandResponse, item *menu.MenuItem) {
			fmt.Println("About Handled")
		})
	// Now for the File menu
	fileMenu.AddItem(2, "Open")
	fileMenu.RegisterEventHandlerByCommandID(2,
		func(reply commands.CommandResponse, item *menu.MenuItem) {
			fmt.Println("Open Handled")
		})
	fileMenu.AddItem(3, "Edit")
	fileMenu.AddSeparator()
	fileMenu.AddItem(4, "Close")

	// Now we just need to plumb our menus together any way we want.

	applicationMenu.AddSubmenu(5, "Application", applicationMenuRoot)
	applicationMenu.AddSubmenu(6, "File", fileMenu)

	// Remember how in basic_browser, Window automatically self registered with the dispatcher.
	// unfortunately we have no such luck here.
	// I suppose this method could be added as an effect of SetApplicationMenu, but the effects of that need to be
	// Ironed out.
	// However, as least we only need to register the top level menu for events, all sub menus will delegate for the top menu.

	// Now we set it as our application Menu
	applicationMenu.SetApplicationMenu()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	dispatcher.RunLoop()
}
