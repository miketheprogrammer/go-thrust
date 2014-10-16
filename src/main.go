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

		fmt.Print("SOCKET::Line", line)
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
	otherSub := Menu{}
	// Calls to other methods after create are Queued until Create returns
	window.Create(conn)
	window.Show(conn)

	menu.Create(conn)
	menu.AddItem(2, "Root", conn)

	fileMenu.Create(conn)
	fileMenu.AddItem(3, "Open", conn)
	fileMenu.AddItem(4, "Close", conn)
	fileMenu.AddSeparator(conn)

	otherSub.Create(conn)
	otherSub.AddCheckItem(5, "Do 1", conn)
	otherSub.SetChecked(5, true, false, conn)
	otherSub.AddSeparator(conn)
	otherSub.AddItem(6, "Do 2", conn)

	fileMenu.AddSubmenu(7, "otherSub", &otherSub, conn)
	menu.AddSubmenu(1, "File", &fileMenu, conn)

	menu.SetApplicationMenu(conn)
	for {
		response := <-ch
		window.DispatchResponse(response, conn)
		menu.DispatchResponse(response, conn)
		if len(fileMenu.WaitingResponses) > 0 {
			for _, v := range fileMenu.WaitingResponses {
				fmt.Println("Waiting for", v.ID, v.Action, v.Method)
			}
		}

	}

}
