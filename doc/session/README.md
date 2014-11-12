# session
--
    import "github.com/miketheprogrammer/go-thrust/session"


## Usage

#### type Cookie

```go
type Cookie struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	// Need to check what type value is,
	// Im on train so have no wifi
	Domain     string    `json:"domain"`
	Path       string    `json:"path"`
	Creation   time.Date `json:"creation"`
	Expiry     time.Date `json:"expiry"`
	LastAccess time.Date `json:"last_access"`
	Secure     bool      `json:"secure"`
	HttpOnly   bool      `json:"http_only"`
	Priority   uint8     `json:"priority"`
}
```


#### type DummySession

```go
type DummySession struct{}
```


#### func  NewDummySession

```go
func NewDummySession() (dummy *DummySession)
```

#### func (DummySession) InvokeCookieForceKeepSessionState

```go
func (d DummySession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesAdd

```go
func (d DummySession) InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesDelete

```go
func (d DummySession) InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesFlush

```go
func (d DummySession) InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesLoad

```go
func (d DummySession) InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesLoadForKey

```go
func (d DummySession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesUpdateAccessTime

```go
func (d DummySession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session)
```

#### type Session

```go
type Session struct {
	TargetID                 uint
	CookieStore              bool
	OffTheRecord             bool
	Ready                    bool
	CommandHistory           []*Command
	ResponseHistory          []*CommandResponse
	WaitingResponses         []*Command
	CommandQueue             []*Command
	SendChannel              *connection.In
	SessionOverrideInterface SessionInvokable
}
```


#### func  NewSession

```go
func NewSession(incognito, overrideDefaultSession bool, saveType string) *Session
```

#### func (*Session) DispatchResponse

```go
func (session *Session) DispatchResponse(response CommandResponse)
```

#### func (*Session) HandleInvoke

```go
func (session *Session) HandleInvoke(reply CommandResponse)
```

#### func (*Session) HandleReply

```go
func (session *Session) HandleReply(reply CommandResponse)
```

#### func (*Session) Send

```go
func (session *Session) Send(command *Command)
```

#### func (*Session) SetInvokable

```go
func (session *Session) SetInvokable(si SessionInvokable)
```

#### type SessionInvokable

```go
type SessionInvokable interface {
	InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session)
	InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session)
	InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session)
}
```

Methods prefixed with Invoke are methods that can be called by ThrustCore, this
differs to our standard call/reply, or event actions, since we are now the
responder.

SessionInvokable is an interface designed to allow you to create your own
Session Store. Simple build a structure that supports these methods, and call
session.SetInvokable(myInvokable)
