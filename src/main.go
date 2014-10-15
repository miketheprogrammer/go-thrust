package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var ActionId = 0

const (
	SOCKET_BOUNDARY = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
)

type ThrustApiObject interface {
	Create(conn net.Conn)
	Call(action string, command *Command, conn net.Conn)
	IsTarget(targetId int) bool
	Handle(reply CommandResponse)
}

func reader(r *bufio.Reader, ch chan CommandResponse) {
	// buf := make([]byte, 2048)
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			panic(err)
		}
		if !strings.Contains(line, SOCKET_BOUNDARY) {
			response := CommandResponse{}
			json.Unmarshal([]byte(line), &response)
			fmt.Println(response)
			ch <- response
		}

		fmt.Println("Got line", line)
	}
	// for {
	// 	n, err := r.Read(buf[:])
	// 	if err != nil {
	// 		return
	// 	}
	// 	println("Client got:", string(buf[0:n]))
	// }
}

func main() {
	addr := flag.String("socket", "", "unix socket where thrust is running")
	flag.Parse()
	fmt.Println(*addr)

	if len(*addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}
	conn, err := net.Dial("unix", *addr)

	defer conn.Close()

	if err != nil {
		os.Exit(2)
	}
	r := bufio.NewReader(conn)
	ch := make(chan CommandResponse)

	go reader(r, ch)

	window := Window{
		Conn: conn,
	}
	window.Create(conn)

	for {
		response := <-ch
		window.HandleReply(response)
		if window.Ready && window.Displayed == false {
			window.Show(conn)
		}
		time.Sleep(time.Millisecond * 10000)

	}

}
