package common

import (
	"os"

	"github.com/alexcesaro/log"
	"github.com/alexcesaro/log/golog"
)

//Global ID tracking for Commands
//Could probably move this to a factory function
var ActionId uint = 0

const (
	SOCKET_BOUNDARY = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
)

var Log *golog.Logger

func InitLogger(sLevel string) {
	var level log.Level
	switch sLevel {
	case "debug":
		level = log.Debug
	case "info":
		level = log.Info
	case "none":
		level = log.None
	default:
		level = log.Error
	}
	Log = golog.New(os.Stdout, level)
	Log.Info("Thrust Client:: Initializing")
}
