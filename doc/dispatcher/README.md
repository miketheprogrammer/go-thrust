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
func RegisterHandler(f HandleFunc)
```

#### func  RunLoop

```go
func RunLoop(outChannels *connection.Out)
```

#### type HandleFunc

```go
type HandleFunc func(commands.CommandResponse)
```
