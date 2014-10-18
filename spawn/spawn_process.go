package spawn

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	. "github.com/miketheprogrammer/go-thrust/common"
)

func SpawnThrustCore() (io.ReadCloser, io.WriteCloser) {

	var thrustExecPath string
	var thrustBoostrapPath string
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if strings.Contains(runtime.GOOS, "darwin") {
		thrustExecPath = dir + "/vendor/darwin/x64/v" + THRUST_VERSION + "/ThrustShell.app/Contents/MacOS/ThrustShell"
		thrustBoostrapPath = dir + "/tools/bootstrap_darwin.sh"
	}
	if strings.Contains(runtime.GOOS, "linux") {
		thrustExecPath = dir + "/vendor/linux/x64/v" + THRUST_VERSION + "/thrust_shell"
		thrustBoostrapPath = dir + "/tools/bootstrap_linux.sh"
	}

	if len(thrustExecPath) > 0 {
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
		Log.Debug("CMD:", thrustExecPath)
		cmd := exec.Command(thrustExecPath)
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
		fmt.Println("Current operating system not supported", runtime.GOOS)
		fmt.Println("===============END====================")
	}
	return nil, nil
}
