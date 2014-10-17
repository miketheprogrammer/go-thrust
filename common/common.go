package common

import (
	"github.com/alexcesaro/log"
	"github.com/alexcesaro/log/stdlog"
)

//Global ID tracking for Commands
//Could probably move this to a factory function
var ActionId uint = 0

const (
	SOCKET_BOUNDARY   = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
	THRUST_VERSION    = "0.7.0"
	THRUST_GO_VERSION = "0.1.0"
)

var Log log.Logger

func InitLogger() {
	Log = stdlog.GetFromFlags()
	Log.Info("Thrust Client:: Initializing")
}
