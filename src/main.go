package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	. "github.com/miketheprogrammer/thrust-go/src/commands"
	. "github.com/miketheprogrammer/thrust-go/src/common"
	. "github.com/miketheprogrammer/thrust-go/src/menu"
	. "github.com/miketheprogrammer/thrust-go/src/spawn"
	. "github.com/miketheprogrammer/thrust-go/src/window"
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
	autoloaderDisabled := flag.Bool("disable-auto-loader", false, "disable auto running of thrust")
	flag.Parse()

	if len(addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}

	SpawnThrustCore(*addr, *autoloaderDisabled)
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
	checkList := Menu{}
	// Calls to other methods after create are Queued until Create returns
	window.Create(conn)
	window.Show(conn)

	menu.Create(conn)
	menu.AddItem(2, "Root", conn)

	fileMenu.Create(conn)
	fileMenu.AddItem(3, "Open", conn)
	fileMenu.AddItem(4, "Close", conn)
	fileMenu.AddSeparator(conn)

	checkList.Create(conn)
	checkList.AddCheckItem(5, "Do 1", conn)
	checkList.SetChecked(5, true, false, conn)
	checkList.AddSeparator(conn)
	checkList.AddCheckItem(6, "Do 2", conn)
	checkList.SetChecked(6, true, false, conn)

	fileMenu.AddSubmenu(7, "CheckList", &checkList, conn)
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
