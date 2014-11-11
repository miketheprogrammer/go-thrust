# window
--
    import "github.com/miketheprogrammer/go-thrust/window"


## Usage

#### type Window

```go
type Window struct {
	TargetID         uint
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	CommandQueue     []*Command
	Url              string
	Title            string
	Ready            bool
	Displayed        bool
	SendChannel      *connection.In `json:"-"`
}
```


#### func (*Window) Call

```go
func (w *Window) Call(command *Command)
```

#### func (*Window) CallWhenDisplayed

```go
func (w *Window) CallWhenDisplayed(command *Command)
```

#### func (*Window) CallWhenReady

```go
func (w *Window) CallWhenReady(command *Command)
```

#### func (*Window) Create

```go
func (w *Window) Create(sess *session.Session)
```

#### func (*Window) DispatchResponse

```go
func (w *Window) DispatchResponse(reply CommandResponse)
```

#### func (*Window) Focus

```go
func (w *Window) Focus()
```

#### func (*Window) HandleError

```go
func (w *Window) HandleError(reply CommandResponse)
```

#### func (*Window) HandleEvent

```go
func (w *Window) HandleEvent(reply CommandResponse)
```

#### func (*Window) HandleReply

```go
func (w *Window) HandleReply(reply CommandResponse)
```

#### func (*Window) IsTarget

```go
func (w *Window) IsTarget(targetId uint) bool
```

#### func (*Window) Maximize

```go
func (w *Window) Maximize()
```

#### func (*Window) Minimize

```go
func (w *Window) Minimize()
```

#### func (*Window) Restore

```go
func (w *Window) Restore()
```

#### func (*Window) Send

```go
func (w *Window) Send(command *Command)
```

#### func (*Window) SetSendChannel

```go
func (w *Window) SetSendChannel(sendChannel *connection.In)
```

#### func (*Window) Show

```go
func (w *Window) Show()
```

#### func (*Window) UnFocus

```go
func (w *Window) UnFocus()
```

#### func (*Window) UnMaximize

```go
func (w *Window) UnMaximize()
```
