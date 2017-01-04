package main

import (
	"fmt"
	"net/http"

	"github.com/cloudspace/go-thrust/lib/bindings/window"
	"github.com/cloudspace/go-thrust/lib/commands"
	"github.com/cloudspace/go-thrust/thrust"
	"github.com/cloudspace/go-thrust/tutorials/provisioner"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlIndex)
}

func main() {
	http.HandleFunc("/", handler)
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()

	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: "http://localhost:8080/",
	})
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()
	thrustWindow.OpenDevtools()
	_, err := thrustWindow.HandleRemote(func(er commands.EventResult, this *window.Window) {
		fmt.Println("RemoteMessage Recieved:", er.Message.Payload)
		// Keep in mind once we have the message, lets say its json of some new type we made,
		// We can unmarshal it to that type.
		// Same goes for the other way around.
		this.SendRemoteMessage("boop")
	})
	if err != nil {
		fmt.Println(err)
		thrust.Exit()
	}
	// See, we dont use thrust.LockThread() because we now have something holding the process open
	http.ListenAndServe(":8080", nil)
}

var htmlIndex string = `
<html>
	<head>
    <script>
    	function write(text) {
    		var e = document.getElementById("log");
    		e.innerHTML = e.innerHTML + text;
    	}
      var start;
    	THRUST.remote.listen(function (event) {
    		var now = (new Date);
    		console.log("ROUNDTRIP IN " + (now - start) + "ms" + " with a length of " + event.payload.length)
    		write( "RESPONSE:" + event.payload + " received in " + (now - start) + "ms<br/>");

    	})
			
			setInterval(function () {
				start = (new Date)
				THRUST.remote.send("beep")
			},1000)
    </script>
	</head>
  <body>
   <p id="log"></p>
  </body>
</html>
`
