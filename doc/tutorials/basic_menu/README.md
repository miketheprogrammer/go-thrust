# basic_menu
--

Ok starting off with out basic_browser template

```go

package main

import (
  "github.com/miketheprogrammer/go-thrust/common"
  "github.com/miketheprogrammer/go-thrust/dispatcher"
  "github.com/miketheprogrammer/go-thrust/spawn"
  "github.com/miketheprogrammer/go-thrust/window"
)

func main() {
  common.InitLogger()
  spawn.SetBaseDirectory("./")
  spawn.Run()
  thrustWindow := window.Window{Url: "http://breach.cc/"}
  thrustWindow.Create(nil)
  thrustWindow.Show()
  thrustWindow.Maximize()
  thrustWindow.Focus()
  // BLOCKING - Dont run before youve excuted all commands you want first.
  dispatcher.RunLoop()
}

```
We need to add the following import.

"github.com/miketheprogrammer/go-thrust/menu"


Next lets make our top level menus
Note all the following code comes before 
the call to dispatcher.RunLoop

```go

  // make our top menus
  //applicationMenu, is essentially the menu bar
  applicationMenu := menu.Menu{}
  //applicationMenuRoot is the first menu, on darwin this is always named the name of your application.
  applicationMenuRoot := menu.Menu{}
  //File menu is our second menu
  fileMenu := menu.Menu{}

```

As you may have noticed by now every item has a Create all. We need to execute those calls in the next snippet, however you may be asking why it is not automatically done in the constructor? Create actually makes a JSONRPC call to Thrust Core to create an object of that type on the heap. This is often what you want, but not always, sometimes you may want to prebuild the model here, and then create it.

Lets now create our menu bar, and the base application root menu.
```go

  // Create our menu bar
  applicationMenu.Create()

  // Lets build our root menu.
  applicationMenuRoot.Create()

  applicationMenuRoot.AddItem(1, "About")

```

Now the same goes for our File menu

```go

  // Now for the File menu
  fileMenu.Create()
  fileMenu.AddItem(2, "Open")
  fileMenu.AddItem(3, "Edit")
  fileMenu.AddSeparator()
  fileMenu.AddItem(4, "Close")

```

Now we just have to plumb together our menus in a way that makes sense. This is not very complex in this example since they are only top level menus.

```go

  applicationMenu.AddSubmenu(5, "Application", &applicationMenuRoot)
  applicationMenu.AddSubmenu(6, "File", &fileMenu)

```

Keep in mind, on darwin systems the first menu attached to the top level menu(menu bar) is the Application menu, it always inherits the name of the application.

In our case the default name is Go Thrust

Remember how in basic_browser, Window automatically self registered with the dispatcher.
unfortunately we have no such luck here.
I suppose this method could be added as an effect of SetApplicationMenu, but the effects of that need to be
Ironed out.
However, as least we only need to register the top level menu for events, all sub menus will delegate for the top menu.

```go

  dispatcher.RegisterHandler(applicationMenu.DispatchResponse)

```

Now we just set the application menu and run the loop

```go

applicationMenu.SetApplicationMenu()

dispatcher.RunLoop()

```
