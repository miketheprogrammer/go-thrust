package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miketheprogrammer/thrust-go/commands"
	. "github.com/miketheprogrammer/thrust-go/common"
	"github.com/miketheprogrammer/thrust-go/menu"
	"github.com/miketheprogrammer/thrust-go/spawn"
	"github.com/miketheprogrammer/thrust-go/window"
)

/*
Reader
Read from the unix socket connection, split on NewLine
Try to json.Unmarshal any value that is not the SOCKET_BOUNDARY
*/
func reader(r *bufio.Reader, ch chan commands.CommandResponse) {
	for {
		line, err := r.ReadString(byte('\n'))
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if !strings.Contains(line, SOCKET_BOUNDARY) {
			response := commands.CommandResponse{}
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

	if len(*addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}

	spawn.SpawnThrustCore(*addr, *autoloaderDisabled)
	conn, err := net.Dial("unix", *addr)

	defer conn.Close()

	if err != nil {
		os.Exit(2)
	}
	r := bufio.NewReader(conn)
	ch := make(chan commands.CommandResponse)

	go reader(r, ch)

	thrustWindow := window.Window{
		Conn: conn,
	}
	rootMenu := menu.Menu{}
	fileMenu := menu.Menu{}
	checkList := menu.Menu{}
	//radioList := Menu{}
	// Calls to other methods after create are Queued until Create returns
	thrustWindow.Create(conn)
	thrustWindow.Show(conn)

	rootMenu.Create(conn)
	rootMenu.AddItem(2, "Root", conn)

	fileMenu.Create(conn)
	fileMenu.AddItem(3, "Open", conn)
	fileMenu.AddItem(4, "Close", conn)
	fileMenu.AddSeparator(conn)

	checkList.Create(conn)
	checkList.AddCheckItem(5, "Do 1", conn)
	checkList.SetChecked(5, true, conn)
	checkList.AddSeparator(conn)
	checkList.AddCheckItem(6, "Do 2", conn)
	checkList.SetChecked(6, true, conn)
	checkList.SetEnabled(6, false, conn)

	fileMenu.AddSubmenu(7, "CheckList", &checkList, conn)
	rootMenu.AddSubmenu(1, "File", &fileMenu, conn)

	rootMenu.SetApplicationMenu(conn)

	for {
		response := <-ch
		thrustWindow.DispatchResponse(response, conn)
		rootMenu.DispatchResponse(response, conn)
		if len(fileMenu.WaitingResponses) > 0 {
			for _, v := range fileMenu.WaitingResponses {
				fmt.Println("Waiting for", v.ID, v.Action, v.Method)
			}
		}

	}

}
