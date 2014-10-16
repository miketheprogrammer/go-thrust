package spawn

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func SpawnThrustCore(addr string, autoloaderDisabled bool) {

	if len(addr) == 0 {
		fmt.Println("System cannot proceed without a socket to connect to. please use -socket={socket_addr}")
		os.Exit(2)
	}
	var thrustExecPath string

	if strings.Contains(runtime.GOOS, "darwin") {
		thrustExecPath = "./vendor/darwin/10.9/ThrustShell.app/Contents/MacOS/ThrustShell"
	}

	if len(thrustExecPath) > 0 && autoloaderDisabled == false {

		go func() {
			cmd := exec.Command(thrustExecPath, "-socket-path="+addr)
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
}
