package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
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
			fmt.Println(err)
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
	fileMenu := Menu{}

	window.Create(conn)
	window.Show(conn)

	menu.Create(conn)
	fileMenu.Create(conn)

	fileMenu.InsertItemAt(0, 1, "Open", conn)
	fileMenu.InsertItemAt(1, 2, "Close", conn)

	for {
		response := <-ch
		window.HandleReply(response, conn)
		menu.HandleReply(response, conn)
		fileMenu.HandleReply(response, conn)
		if menu.Ready == true &&
			fileMenu.Ready &&
			fileMenu.Parent != &menu &&
			len(fileMenu.WaitingResponses) == 0 {
			menu.InsertSubmenuAt(1, 3, "File", &fileMenu, conn)
		}
		// if window.Ready && window.Displayed == false {
		// 	fmt.Println("Window Ready")
		// 	window.Show(conn)
		// }

	}

}
