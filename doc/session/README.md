# session
--
    import "github.com/miketheprogrammer/go-thrust/session"


## Usage

#### type Cookie

```go
type Cookie struct {
	Source     string `json:"source"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Domain     string `json:"domain"`
	Path       string `json:"path"`
	Creation   int64  `json:"creation"`
	Expiry     int64  `json:"expiry"`
	LastAccess int64  `json:"last_access"`
	Secure     uint   `json:"secure"`
	HttpOnly   bool   `json:"http_only"`
	Priority   uint   `json:"priority"`
}
```

Cookie source the source url name the cookie name value the cookie value domain
the cookie domain path the cookie path creation the creation date expiry the
expiration date last_access the last time the cookie was accessed secure is the
cookie secure http_only is the cookie only valid for HTTP priority internal
priority information

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
func (ds DummySession) InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session)
```

#### func (DummySession) InvokeCookiesAdd

```go
func (ds DummySession) InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session) bool
```

#### func (DummySession) InvokeCookiesDelete

```go
func (ds DummySession) InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session) bool
```

#### func (DummySession) InvokeCookiesFlush

```go
func (ds DummySession) InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session) bool
```

#### func (DummySession) InvokeCookiesLoad

```go
func (ds DummySession) InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie)
```
For Simplicity type declarations

#### func (DummySession) InvokeCookiesLoadForKey

```go
func (ds DummySession) InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie)
```

#### func (DummySession) InvokeCookiesUpdateAccessTime

```go
func (ds DummySession) InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session) bool
```

#### type Session

```go
type Session struct {
	TargetID                 uint
	CookieStore              bool
	OffTheRecord             bool
	Path                     string
	Ready                    bool
	CommandHistory           []*Command
	ResponseHistory          []*CommandResponse
	WaitingResponses         []*Command
	CommandQueue             []*Command
	SendChannel              *connection.In
	SessionOverrideInterface SessionInvokable
}
```

Session is the core API Binding object used to communicate with Thrust.

#### func  NewSession

```go
func NewSession(incognito, overrideDefaultSession bool, path string) *Session
```
NewSession is a constructor that takes 3 arguments, incognito which is a
boolean, meaning dont persist session state after close. overrideDefaultSession
which is a boolean that till tell thrust core to try to invoke session methods
from us. path string is the path to store session data.

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
	InvokeCookiesLoad(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie)
	InvokeCookiesLoadForKey(args *commands.CommandResponseArguments, session *Session) (cookies []Cookie)
	InvokeCookiesFlush(args *commands.CommandResponseArguments, session *Session) bool
	InvokeCookiesAdd(args *commands.CommandResponseArguments, session *Session) bool
	InvokeCookiesUpdateAccessTime(args *commands.CommandResponseArguments, session *Session) bool
	InvokeCookiesDelete(args *commands.CommandResponseArguments, session *Session) bool
	InvokeCookieForceKeepSessionState(args *commands.CommandResponseArguments, session *Session)
}
```
