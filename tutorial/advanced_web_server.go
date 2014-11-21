package main

import (
	"fmt"
	"net/http"

	"github.com/miketheprogrammer/go-thrust"
	"github.com/miketheprogrammer/go-thrust/lib/session"
	"github.com/miketheprogrammer/go-thrust/tutorial/provisioner"
	"github.com/miketheprogrammer/go-thrust/x/web"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func main() {
	http.HandleFunc("/", handler)
	webHandler := web.NewWebHandler()
	http.Handle("/web", webHandler)

	// thrust.Start() must always come before any bindings are created.
	thrust.DisableLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()
	mysession := session.NewSession(false, false, "cache")

	thrustWindow := thrust.NewWindow("http://localhost:8080/", mysession)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

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
