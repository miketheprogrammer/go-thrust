# basic_browser
--

For the basic browser package we first need to add our standard imports for starting Thrust as well as delcare the package name. Since its an executable, we declare pacage main.

- Quick note about synchronous vs. asynchronous. Unless otherwise stated, most API commands are IO Bound, and asynchronous in nature. However, they look synchronous. Most of the nitty gritty details are abstracted away from the user, with a sort of priority queue / state driven api. Some calls will only execute when certain states have been met. For instance, you cannot set a menu as the application menu until the menu and all its items and submenus have been created and have stabilized meaning they have no events waiting to return.

```go
package main

import (
  "github.com/miketheprogrammer/go-thrust/common"
  "github.com/miketheprogrammer/go-thrust/dispatcher"
  "github.com/miketheprogrammer/go-thrust/spawn"
  "github.com/miketheprogrammer/go-thrust/window"
)
```
Next lets create our func main 

```go
func main() {
  
}
```

Please add any flags you want parsed at the
beginning of your main function. 

```go
func main () {
  common.InitLogger()
}
```

Next we need a daemon to connect to, lets connect up our autodownload/autospawner for thrust core. But first, will set it to use our local directory for download storage. It will always install to a directory called vendor, wherever the path you provide is.
If no path is provided, it will use the current users home directory (at least for now.)

```go
func main () {
  common.InitLogger()

  spawn.SetBaseDirectory("./")
  spawn.Run()
}
```

Now it is time to create our Window.
We will go through several methods in the next code snippet please read the comments

```go
func main () {
  // Initialize the Logger
  common.InitLogger()

  spawn.SetBaseDirectory("./")
  spawn.Run()

  // Create the window struct with default values
  thrustWindow := window.NewWindow("http://breach.cc/", nil)

  // Show our new window.
  thrustWindow.Show()

  // Maximize our new window
  thrustWindow.Maximize()

  // Focus our window
  thrustWindow.Focus()
}
```

Now our final step, running the dispatcher loop. This or some other loop must be blocking to keep the process open.

The dispatcher is your key to accessing many of the internal. The method RunLoop is a helper that will serve as the standard loop, however you can feel free to implement your own, if you dont mind any compatibility issues with forward releases.

Quick Note, dispatchers and other objects support registering handlers. Generally, objects such as Window, Menu, Session, will self register with the dispatcher.

This may change in the coming weeks, but for demo purposes we wanted the simplest code possible to get started. 

```go
func main () {
  // Initialize the Logger
  common.InitLogger()

  spawn.SetBaseDirectory("./")
  spawn.Run()

  // Create the window struct with default values
  thrustWindow := window.NewWindow("http://breach.cc/", nil)



  // Show our new window.
  thrustWindow.Show()

  // Maximize our new window
  thrustWindow.Maximize()

  // Focus our window
  thrustWindow.Focus()

  // Finally run the dispatcher loop
  dispatcher.RunLoop()
  // Dont run as a goroutine unless you have another blocking action on this thread.
}
```

You now have a successful window showing you http://breach.cc