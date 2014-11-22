package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/miketheprogrammer/go-thrust"
	"github.com/miketheprogrammer/go-thrust/lib/bindings/menu"
	"github.com/miketheprogrammer/go-thrust/lib/commands"
	"github.com/miketheprogrammer/go-thrust/tutorial/provisioner"
)

func main() {
	//db, err := leveldb.OpenFile("./tmp/leveldb/my.db", nil)

	thrust.InitLogger()
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	thrust.Start()

	mainWindow := thrust.NewWindow("http://localhost:8081/asset/index.html", nil)
	mainWindow.Show()
	mainWindow.Maximize()
	mainWindow.Focus()

	//mainWindow.OpenDevtools()

	m := martini.Classic()

	m.Get("/asset/:file", AssetHandler)
	m.Get("/contextmenu", (func() func(http.ResponseWriter, *http.Request) {
		contextmenu := thrust.NewMenu()
		contextmenu.AddItem(1, "Open Devtools")
		contextmenu.RegisterEventHandlerByCommandID(1, func(reply commands.CommandResponse, item *menu.MenuItem) {
			mainWindow.OpenDevtools()
		})
		contextmenu.AddItem(2, "Close Devtools")
		contextmenu.RegisterEventHandlerByCommandID(3, func(reply commands.CommandResponse, item *menu.MenuItem) {
			mainWindow.CloseDevtools()
		})
		return func(w http.ResponseWriter, req *http.Request) {
			contextmenu.Popup(mainWindow)
		}

	})())

	log.Fatal(http.ListenAndServe("127.0.0.1:8081", m))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello,  ` world!\n")
}

func AssetHandler(w http.ResponseWriter, req *http.Request, params martini.Params) {
	bytes, err := Asset("data/" + params["file"])
	if err != nil {
		io.WriteString(w, params["file"])
		return
	}

	io.WriteString(w, string(bytes))
	return
}
