package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	. "github.com/miketheprogrammer/thrust-go/src/commands"
	. "github.com/miketheprogrammer/thrust-go/src/common"
	. "github.com/miketheprogrammer/thrust-go/src/menu"
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
	fmt.Println(*addr)

	if len(*addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}
	var thrustExecPath string

	if strings.Contains(runtime.GOOS, "darwin") {
		thrustExecPath = "./vendor/darwin/10.9/ThrustShell.app/Contents/MacOS/ThrustShell"
	}

	if len(thrustExecPath) > 0 && *autoloaderDisabled == false {

		go func() {
			cmd := exec.Command(thrustExecPath, "-socket-path=/tmp/thrust.sock")
			cmdIn, _ := cmd.StdinPipe()
			cmdOut, _ := cmd.StdoutPipe()
			cmdErr, _ := cmd.StderrPipe()

			cmd.Start()
			defer cmdIn.Close()

			for {
				outBytes, _ := ioutil.ReadAll(cmdOut)
				errBytes, _ := ioutil.ReadAll(cmdErr)

				fmt.Print(string(outBytes))
				fmt.Print(string(errBytes))

				time.Sleep(time.Millisecond * 10)
			}
		}()
		time.Sleep(time.Millisecond * 1000)
	} else {
		fmt.Println("===============WARNING================")
		fmt.Println("Auto Loading of thrust currently not supported for", runtime.GOOS)
		fmt.Println("Please run thrust executable manually")
		fmt.Println("===============END====================")
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
