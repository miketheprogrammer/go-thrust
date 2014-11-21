package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/miketheprogrammer/go-thrust/lib/bindings/window"
	"github.com/miketheprogrammer/go-thrust/lib/dispatcher"
	"github.com/miketheprogrammer/go-thrust/lib/spawn"
)

var (
	port = flag.Int("port", 8000, "TCP port to listen on")
)
var html *template.Template = template.New("main")

func init() {
	template.Must(html.New("browser.html").Parse(string(browser_html)))

}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err := html.ExecuteTemplate(w, "browser.html", nil)
	if err != nil {
		log.Printf("Template error: %v", err)
	}
}

func main() {
	flag.Parse()
	spawn.SetBaseDirectory("./")
	spawn.Run()
	thrustWindow := window.NewWindow(fmt.Sprintf("http://127.0.0.1:%d", *port), nil)
	thrustWindow.Show()
	thrustWindow.Focus()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	go dispatcher.RunLoop()

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	http.Handle("/", http.HandlerFunc(index))
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic(err)
	}
}
