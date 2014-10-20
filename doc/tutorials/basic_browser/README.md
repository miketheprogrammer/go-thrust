# basic_browser
--

For the basic browser package we first need to add our standard imports for starting Thrust as well as delcare the package name. Since its an executable, we declare pacage main.

- Note: the Common package is imported using . notation, this imports all exported symbols into the local space. Symbols such as Log, InitLogger(), and SOCKET_BOUNDARY


```go
package main

import (
  "github.com/miketheprogrammer/go-thrust/commands"
  . "github.com/miketheprogrammer/go-thrust/common"
  "github.com/miketheprogrammer/go-thrust/connection"
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

Please add any flags you want parsed at the beginning as the next command will parse all flags.
The next command is InitLogger() this will initialize the core logger

```go
func main () {
  InitLogger()
}
```

Next we need a daemon to connect to, lets connect up our autodownload/autospawner for thrust core.
We do this by assigning the two outputs of spawn.SpawnThrustCore to connection.Stdout and connection.Stdin, dont confuse the order of these.

```go
func main () {
  InitLogger()

  connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()
}
```

Now, we need to initialize our connection threads.


```go
func main () {
  InitLogger()

  connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()

  err := connection.InitializeThreads()
  if err != nil {
    fmt.Println(err)
    os.Exit(2) // Whatever error code you want to throw here.
  }
}
```
Next lets grab our connection channels, and store them in two variables.
Please refer to the connection package documenation for more info on
what is returned from GetCommunicationChannels()

```go
func main () {
  // Initialize the Logger
  InitLogger()

  // Spawn Thrust core and connect it to the connection package
  connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()

  // Initialize the Connection packages threads
  err := connection.InitializeThreads()

  if err != nil {
    fmt.Println(err)
    os.Exit(2) // Whatever error code you want to throw here.
  }

  // Store the communcation channels
  out, in := connection.GetCommunicationChannels()
}
```

Now it is time to finially create our Window.
We will go through several methods in the next code snippet please read the comments

```go
func main () {
  // Initialize the Logger
  InitLogger()

  // Spawn Thrust core and connect it to the connection package
  connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()

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

}
```

There is one majore part left, your dispatching thread. We need to create a thread that will dispatch the commands to handlers. As soon as the api is standardized this may be removed, but for now, you handle it.

```go
func main () {
  // Initialize the Logger
  InitLogger()

  // Spawn Thrust core and connect it to the connection package
  connection.StdOut, connection.StdIn = spawn.SpawnThrustCore()

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

  // Register a handler for thrustWindow
  dispatcher.RegisterHandler(func(c commands.CommandResponse) {
    thrustWindow.DispatchResponse(c)
  })
  // Create your dispatching thread.
  for {
    select {
    case response := <-out.CommandResponses:
      dispatcher.Dispatch(response)
    default:
      break
    }
    time.Sleep(time.Microsecond * 10)
  }
}
```

You now have a successful window.