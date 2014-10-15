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

//Global ID tracking for Commands
//Could probably move this to a factory function
var ActionId int = 0

const (
	SOCKET_BOUNDARY = "--(Foo)++__THRUST_SHELL_BOUNDARY__++(Bar)--"
)

/*
Reader
Read from the unix socket connection, split on NewLine
Try to json.Unmarshal any value that is not the SOCKET_BOUNDARY
*/
func reader(r *bufio.Reader, ch chan CommandResponse) {
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
	menu := Menu{}
	window.Create(conn)
	menu.Create(conn)
	setMenu := func() {
		menu.InsertItemAt(1, 1, "MyItem", conn)
		menu.InsertItemAt(2, 2, "MyItem", conn)
		time.Sleep(time.Millisecond * 2000)
		//menu.SetApplicationMenu(conn)
	}
	for {
		response := <-ch
		window.HandleReply(response)
		menu.HandleReply(response)
		if window.Ready && window.Displayed == false {
			fmt.Println("Window Ready")
			window.Show(conn)
		}

		if menu.Ready && menu.Displayed == false {
			setMenu()
			setMenu = func() {}
		}

	}

}
