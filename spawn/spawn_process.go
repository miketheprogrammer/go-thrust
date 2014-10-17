package spawn

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	. "github.com/miketheprogrammer/go-thrust/common"
)

func SpawnThrustCore(addr string, autoloaderDisabled bool) {

	var thrustExecPath string
	var thrustBoostrapPath string
	if strings.Contains(runtime.GOOS, "darwin") {
		thrustExecPath = "./vendor/darwin/x64/ThrustShell.app/Contents/MacOS/ThrustShell"
		thrustBoostrapPath = "./tools/bootstrap_darwin.sh"
	}
	if strings.Contains(runtime.GOOS, "linux") {
		thrustExecPath = "./vendor/linux/x64/thrust_shell"
		thrustBoostrapPath = "./tools/bootstrap_linux.sh"
	}

	if len(thrustExecPath) > 0 && autoloaderDisabled == false {
		if _, err := os.Stat(thrustExecPath); os.IsNotExist(err) {
			Log.Info("Could not find executable:", thrustExecPath)
			Log.Info("Attempting to Download and Install the Thrust Core Executable")

			installCmd := exec.Command("sh", thrustBoostrapPath)
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr

			installCmd.Start()
			Log.Info("Waiting for install to finish....")
			installErr := installCmd.Wait()

			if installErr != nil {
				Log.Errorf("Could not bootstrap, ErrorCode:", err)
			} else {
				Log.Info("... Done Bootstrapping")
			}
		}

		Log.Info("Attempting to start Thrust Core")
		Log.Debug("CMD:", thrustExecPath, "-socket-path="+addr)
		cmd := exec.Command(thrustExecPath, "-socket-path="+addr)
		//cmdIn, _ := cmd.StdinPipe()

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout

		cmd.Start()

		time.Sleep(time.Millisecond * 2000)
		Log.Info("Returning to Main Process.")
	} else {
		fmt.Println("===============WARNING================")
		fmt.Println("Auto Loading of thrust currently not supported for", runtime.GOOS)
		fmt.Println("Please run thrust executable manually")
		fmt.Println("===============END====================")
	}
}
