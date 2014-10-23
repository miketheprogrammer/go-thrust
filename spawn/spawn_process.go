package spawn

/*
Package spawn implements methods and interfaces used in downloading and spawning the underlying thrust core binary.
*/
import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	. "github.com/miketheprogrammer/go-thrust/common"
)

/*
The SpawnThrustCore method is a bootstrap and run method.
It will try to detect an installation of thrust, if it cannot find it
it will download the version of Thrust detailed in the "common" package.
Once downloaded, it will launch a process.
Go-Thrust and all *-Thrust packages communicate with Thrust Core via Stdin/Stdout.
using -log=debug as a command switch will give you the most information about what is going on. -log=info will give you notices that stuff is happening.
Any log level higher than that will output nothing.
*/

func SpawnThrustCore(dir string) (io.ReadCloser, io.WriteCloser) {
	if len(dir) == 0 {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(usr.HomeDir)
		// Parses Flags
		dir = usr.HomeDir
	}

	SetBaseDirectory(dir)

	var thrustExecPath string

	thrustExecPath = GetExecutablePath()
	if len(thrustExecPath) > 0 {

		Bootstrap()

		thrustExecPath = GetExecutablePath()

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

		Log.Info("Thrust Core started.")
		return cmdOut, cmdIn
	} else {
		fmt.Println("===============WARNING================")
		fmt.Println("Current operating system not supported", runtime.GOOS)
		fmt.Println("===============END====================")
	}
	return nil, nil
}

func downloadFromUrl(url, version string) {
	url = strings.Replace(url, "$V", version, 2)
	fileName := strings.Replace("/tmp/$V", "$V", version, 1)
	if Log.LogInfo() {
		fmt.Println("Downloading", url, "to", fileName)
	}

	quit := make(chan int, 1)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Print("\n")
				return
			default:
				if Log.LogInfo() {
					fmt.Print(".")
				}
			}
			time.Sleep(time.Second)
		}
	}()

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	quit <- 1

	fmt.Println(n, "bytes downloaded.")
}
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	if Log.LogInfo() {
		fmt.Println("Unzipping", src, "to", dest)
	}

	quit := make(chan int, 1)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Print("\n")
				return
			default:
				if Log.LogInfo() {
					fmt.Print(".")
				}
			}
			time.Sleep(time.Second)
		}
	}()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0775)
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, 0775)
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0775)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	quit <- 1
	return nil
}
