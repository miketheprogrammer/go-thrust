package spawn

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	. "github.com/miketheprogrammer/go-thrust/common"
)

func SpawnThrustCore(addr string, autoloaderDisabled bool) (io.ReadCloser, io.WriteCloser) {

	var thrustExecPath string
	var thrustBoostrapPath string
	if strings.Contains(runtime.GOOS, "darwin") {
		thrustExecPath = "./vendor/darwin/x64/v" + THRUST_VERSION + "/ThrustShell.app/Contents/MacOS/ThrustShell"
		thrustBoostrapPath = "./tools/bootstrap_darwin.sh"
	}
	if strings.Contains(runtime.GOOS, "linux") {
		thrustExecPath = "./vendor/linux/x64/v" + THRUST_VERSION + "/thrust_shell"
		thrustBoostrapPath = "./tools/bootstrap_linux.sh"
	}

	if len(thrustExecPath) > 0 && autoloaderDisabled == false {
		if _, err := os.Stat(thrustExecPath); os.IsNotExist(err) {
			Log.Info("Could not find executable:", thrustExecPath)
			Log.Info("Attempting to Download and Install the Thrust Core Executable")

			installCmd := exec.Command("sh", thrustBoostrapPath, THRUST_VERSION)
			if Log.LogDebug() {
				installCmd.Stdout = os.Stdout
				installCmd.Stderr = os.Stderr
			}

			installDoneChan := make(chan bool, 1)
			installCmd.Start()
			fmt.Print("Installing ")
			go func() {
				for {
					select {
					case <-installDoneChan:
						return
					default:
						fmt.Print(".")
						time.Sleep(time.Second)
					}
				}
			}()

			installErr := installCmd.Wait()
			installDoneChan <- true
			if installErr != nil {
				Log.Errorf("Could not bootstrap, ErrorCode:", err)
			} else {
				Log.Info("... Done Bootstrapping")
			}
		}

		Log.Info("Attempting to start Thrust Core")
		Log.Debug("CMD:", thrustExecPath, "-socket-path="+addr)
		cmd := exec.Command(thrustExecPath, "-socket-path="+addr)
		cmdIn, e1 := cmd.StdinPipe()
		cmdOut, e2 := cmd.StdoutPipe()

		if e1 != nil {
			fmt.Println(e1)
			os.Exit(2)
		}

		if e2 != nil {
			fmt.Println(e2)
			os.Exit(2)
		}

		if Log.LogDebug() {
			cmd.Stderr = os.Stdout
		}

		cmd.Start()

		//time.Sleep(time.Millisecond * 1000)
		Log.Info("Returning to Main Process.")
		return cmdOut, cmdIn
	} else {
		fmt.Println("===============WARNING================")
		fmt.Println("Auto Loading of thrust currently not supported for", runtime.GOOS)
		fmt.Println("Please run thrust executable manually")
		fmt.Println("===============END====================")
	}
	return nil, nil
}
