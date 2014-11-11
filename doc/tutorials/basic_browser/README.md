# basic_browser
--

For the basic browser package we first need to add our standard imports for starting Thrust as well as delcare the package name. Since its an executable, we declare pacage main.

- Note: the Common package is imported using . notation, this imports all exported symbols into the local space. Symbols such as Log, InitLogger(), and SOCKET_BOUNDARY


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
  thrustWindow := window.Window{
    Url: "http://breach.cc/",
  }

  // Send a Create call to the Thrust core requesting the window to be created
  // Dont worry about return values, we handle that asynchronously behind the 
  // scenes
  thrustWindow.Create(nil)

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

Quick Note, dispatchers and other objects support registering handlers. You generally register a top level or root object with the dispatcher and then have that object process the even internally, for instance in the case of menus, register handlers per ActionId.

For simplicity sake, single instance objects such as Window (only that one for now) register themselves with their default handler.

This may change in the coming weeks, but for demo purposes we wanted the simplest code possible to get started. 

```go
func main () {
  // Initialize the Logger
  common.InitLogger()

  spawn.SetBaseDirectory("./")
  spawn.Run()

  // Create the window struct with default values
  thrustWindow := window.Window{
    Url: "http://breach.cc/",
  }

  // Send a Create call to the Thrust core requesting the window to be created
  // Dont worry about return values, we handle that asynchronously behind the 
  // scenes
  thrustWindow.Create(nil)

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

You now have a successful window.