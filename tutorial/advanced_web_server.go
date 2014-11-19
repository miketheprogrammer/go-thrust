package main

import (
	"fmt"
	"net/http"

	"github.com/miketheprogrammer/go-thrust/dispatcher"
	"github.com/miketheprogrammer/go-thrust/session"
	"github.com/miketheprogrammer/go-thrust/spawn"
	"github.com/miketheprogrammer/go-thrust/web"
	"github.com/miketheprogrammer/go-thrust/window"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func main() {
	http.HandleFunc("/", handler)
	webHandler := web.NewWebHandler()
	http.Handle("/web", webHandler)
	spawn.SetBaseDirectory("./")
	spawn.Run()

	mysession := session.NewSession(false, false, "cache")

	thrustWindow := window.NewWindow("http://localhost:8080/", mysession)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	// NonBLOCKING - note in other examples this was blocking.
	go dispatcher.RunLoop()

	http.ListenAndServe(":8080", nil)
}

var htmlIndex string = `
<html>
  <body>
    <h2> Welcome to Go-Thrust <h3>
    <img height="50px" width="120px" src="http://i.imgur.com/DwFKI0J.png"/>
    <script>
      window.onload = function() {
        setTimeout(function() {
          var webview = document.createElement("webview");
          document.body.appendChild(webview);
          webview.src = "http://www.google.com";
          //webview.classList.add("active");
        }, 0);
      }
    </script>
  </body>
</html>
`
