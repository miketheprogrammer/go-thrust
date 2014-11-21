package thrust

import (
	"runtime"

	"github.com/miketheprogrammer/go-thrust/lib/bindings/menu"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/session"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/common"
	"github.com/miketheprogrammer/go-thrust/lib/connection"
	"github.com/miketheprogrammer/go-thrust/lib/dispatcher"
	"github.com/miketheprogrammer/go-thrust/lib/spawn"
)

func NewWindow(url string, sess *session.Session) *window.Window {
	return window.NewWindow(url, sess)
}

func NewSession(incognito, overrideDefaultSession bool, path string) *session.Session {
	return session.NewSession(incognito, overrideDefaultSession, path)
}

func NewMenu() *menu.Menu {
	return menu.NewMenu()
}

func Start() {
	spawn.Run()
	go dispatcher.RunLoop()
}

func SetProvisioner(p spawn.Provisioner) {
	spawn.SetProvisioner(p)
}

func LockThread() {
	for {
		runtime.Gosched()
	}
}

func InitLogger() {
	common.InitLogger("")
}

func DisableLogger() {
	common.InitLogger("none")
}

func Exit() {
	connection.CleanExit()
}
