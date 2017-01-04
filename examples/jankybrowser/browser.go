package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/cloudspace/go-thrust/thrust"
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
	thrust.InitLogger()
	thrust.Start()

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: fmt.Sprintf("http://127.0.0.1:%d", *port),
	})
	thrustWindow.Show()
	thrustWindow.Focus()

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	http.Handle("/", http.HandlerFunc(index))
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic(err)
	}
}
