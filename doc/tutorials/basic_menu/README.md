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
  thrustWindow := window.NewWindow("http://breach.cc/", nil)
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
the call to dispatcher.RunLoop, It doesnt have to, you can run dispatcher runloop whenever you want, but it will block the process.

As long as you have something blocking the main thread from exiting the process will stay open.

for instance the following pseudo code.
http.server.start()
go createApplicationMenu()
go dispatcher.RunLoop()

Anyway lets proceed

```go

  // make our top menus
  applicationMenu := menu.NewMenu()
  //applicationMenuRoot is the first menu, on darwin this is always named the name of your application.
  applicationMenuRoot := menu.NewMenu()
  //File menu is our second menu
  fileMenu := menu.NewMenu()

```

Lets now create our menu bar, and the base application root menu.
```go
  applicationMenuRoot.AddItem(1, "About")
```

Now the same goes for our File menu

```go

  // Now for the File menu
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

Now we just set the application menu and run the loop

```go

applicationMenu.SetApplicationMenu()

dispatcher.RunLoop()

```
