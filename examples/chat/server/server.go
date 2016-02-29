package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/cloudspace/go-thrust/thrust"
	"github.com/gorilla/websocket"
	"github.com/tv42/birpc"
	"github.com/tv42/birpc/wetsock"
	"github.com/tv42/topic"
)

var (
	host = flag.String("host", "0.0.0.0", "IP address to bind to")
	port = flag.Int("port", 8000, "TCP port to listen on")
)

var html *template.Template = template.New("main")

func init() {
	template.Must(html.New("chat.html").Parse(string(chat_html)))
	template.Must(html.New("chat.css").Parse(string(chat_css)))
	template.Must(html.New("chat.js").Parse(string(chat_js)))
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

type Incoming struct {
	From    string
	Message string
}

type Outgoing struct {
	Time    time.Time `json:"time"`
	From    string    `json:"from"`
	Message string    `json:"message"`
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err := html.ExecuteTemplate(w, "chat.html", nil)
	if err != nil {
		log.Printf("Template error: %v", err)
	}
}

type Chat struct {
	broadcast *topic.Topic
	registry  *birpc.Registry
}

// Closing the socket from another goroutine causes a concurrent
// blocking read to return a net.errClosing error. It's pointless to
// spam logs with it. The error is not exported, so checking for it is
// ugly.
//
// Background: https://code.google.com/p/go/issues/detail?id=4373
func isErrClosing(err error) bool {
	switch err2 := err.(type) {
	case *net.OpError:
		err = err2.Err
	}
	return err.Error() == "use of closed network connection"
}

type nothing struct{}

func (c *Chat) Message(msg *Incoming, _ *nothing, ws *websocket.Conn) error {
	log.Printf("recv from %v:%#v\n", ws.RemoteAddr(), msg)

	c.broadcast.Broadcast <- Outgoing{
		Time:    time.Now(),
		From:    msg.From,
		Message: msg.Message,
	}
	return nil
}

func main() {
	prog := path.Base(os.Args[0])
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() > 0 {
		Usage()
		os.Exit(1)
	}

	log.Printf("Serving at http://%s:%d/", *host, *port)

	chat := Chat{}
	chat.broadcast = topic.New()
	chat.registry = birpc.NewRegistry()
	chat.registry.RegisterService(&chat)
	defer close(chat.broadcast.Broadcast)
	upgrader := websocket.Upgrader{}

	serve := func(w http.ResponseWriter, req *http.Request) {
		ws, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		endpoint := wetsock.NewEndpoint(chat.registry, ws)
		messages := make(chan interface{}, 10)
		chat.broadcast.Register(messages)
		go func() {
			defer chat.broadcast.Unregister(messages)
			for i := range messages {
				msg := i.(Outgoing)
				// Fire-and-forget.
				// TODO use .Notify when it exists
				_ = endpoint.Go("Chat.Message", msg, nil, nil)
			}
			// broadcast topic kicked us out for being too slow;
			// probably a hung TCP connection. let client
			// re-establish.
			log.Printf("Kicking slow client: %v", ws.RemoteAddr())
			ws.Close()
		}()

		if err := endpoint.Serve(); err != nil {
			log.Printf("websocket error from %v: %v", ws.RemoteAddr(), err)
		}
	}

	http.HandleFunc("/sock", serve)
	http.Handle("/", http.HandlerFunc(index))
	addr := fmt.Sprintf("%s:%d", *host, *port)

	thrust.InitLogger()
	thrust.Start()
	thrustWindow := thrust.NewWindow(thrust.WindowOptions{
		RootUrl: fmt.Sprintf("http://127.0.0.1:%d", *port),
	})
	thrustWindow.Show()
	thrustWindow.Focus()

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
