# common
--
    import "github.com/miketheprogrammer/go-thrust/common"


## Usage

```go
const (
	THRUST_GO_VERSION = "0.3.1"
)
```

```go
const (
	SOCKET_BOUNDARY = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
)
```

```go
var ActionId uint = 0
```

Global ID tracking for Commands Could probably move this to a factory function

```go
var Log *golog.Logger
```

#### func  InitLogger

```go
func InitLogger(sLevel string)
```
