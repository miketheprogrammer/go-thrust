package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/miketheprogrammer/thrust-go/commands"
	"github.com/miketheprogrammer/thrust-go/connection"
	"github.com/miketheprogrammer/thrust-go/dispatcher"
	"github.com/miketheprogrammer/thrust-go/menu"
	"github.com/miketheprogrammer/thrust-go/spawn"
	"github.com/miketheprogrammer/thrust-go/window"
)

func main() {
	addr := flag.String("socket", "", "unix socket where thrust is running")
	autoloaderDisabled := flag.Bool("disable-auto-loader", false, "disable auto running of thrust")
	flag.Parse()

	if len(*addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}

	spawn.SpawnThrustCore(*addr, *autoloaderDisabled)

	err := connection.InitializeThreads("unix", *addr)
	if err != nil {
		os.Exit(2)
	}
	out, in := connection.GetCommunicationChannels()

	thrustWindow := window.Window{}
	rootMenu := menu.Menu{}
	fileMenu := menu.Menu{}
	checkList := menu.Menu{}
	radioList := menu.Menu{}
	// Calls to other methods after create are Queued until Create returns
	thrustWindow.Create(in)
	thrustWindow.Show()

	rootMenu.Create(in)
	rootMenu.AddItem(2, "Root")

	fileMenu.Create(in)
	fileMenu.AddItem(3, "Open")
	fileMenu.AddItem(4, "Close")
	fileMenu.AddSeparator()

	checkList.Create(in)
	checkList.AddCheckItem(5, "Check 1")
	checkList.SetChecked(5, true)
	checkList.AddSeparator()
	checkList.AddCheckItem(6, "Check 2")
	checkList.SetChecked(6, true)
	checkList.SetEnabled(6, false)

	radioList.Create(in)
	radioList.AddRadioItem(7, "Radio 1-1", 1)
	radioList.AddRadioItem(8, "Radio 1-2", 1)
	radioList.AddSeparator()
	radioList.AddRadioItem(9, "Radio 2-1", 2)
	radioList.AddRadioItem(10, "Radio 2-2", 2)
	radioList.SetVisible(6, false)

	fileMenu.AddSubmenu(11, "CheckList", &checkList)
	fileMenu.AddSubmenu(12, "RadioList", &radioList)
	rootMenu.AddSubmenu(1, "File", &fileMenu)

	rootMenu.SetApplicationMenu()

	thrustWindow.Maximize()

	dispatcher.RegisterHandler(func(c commands.CommandResponse) {
		thrustWindow.DispatchResponse(c)
	})

	dispatcher.RegisterHandler(func(c commands.CommandResponse) {
		rootMenu.DispatchResponse(c)
	})
	for {
		select {
		case response := <-out.CommandResponses:
			//thrustWindow.DispatchResponse(response)
			//rootMenu.DispatchResponse(response)
			dispatcher.Dispatch(response)
			if len(fileMenu.WaitingResponses) > 0 {
				fileMenu.PrintRecursiveWaitingResponses()
			}
		default:
			break
		}
		time.Sleep(time.Microsecond * 10)

	}

}
