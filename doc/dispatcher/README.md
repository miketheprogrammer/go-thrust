# dispatcher
--
    import "github.com/miketheprogrammer/go-thrust/dispatcher"


## Usage

#### func  Dispatch

```go
func Dispatch(command commands.CommandResponse)
```
Dispatch dispatches a CommandResponse to every handler in the registry

#### func  RegisterHandler

```go
func RegisterHandler(f HandleFunc)
```
RegisterHandler registers a HandleFunc f to receive a CommandResponse when one
is sent to the system.

#### func  Run

```go
func Run(outChannels *connection.Out)
```

#### func  RunLoop

```go
func RunLoop()
```
RunLoop starts a loop that receives CommandResponses and dispatches them. This
is a helper method, but you could just implement your own, if you only need this
loop to be the blocking loop. For Instance, in a HTTP Server setting, you might
want to run this as a goroutine and then let the servers blocking handler keep
the process open. As long as there are commands in the channel, this loop will
dispatch as fast as possible, when all commands are exhausted this loop will run
on iterations of 10 microseconds optimistically.

#### type HandleFunc

```go
type HandleFunc func(commands.CommandResponse)
```
