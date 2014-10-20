# dispatcher
--
    import "github.com/miketheprogrammer/go-thrust/dispatcher"


## Usage

#### func  Dispatch

```go
func Dispatch(command commands.CommandResponse)
```

#### func  RegisterHandler

```go
func RegisterHandler(f DispatcherHandleFunc)
```

#### func  RunLoop

```go
func RunLoop(outChannels *connection.Out)
```

#### type DispatcherHandleFunc

```go
type DispatcherHandleFunc func(commands.CommandResponse)
```
