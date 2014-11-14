package spawn

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	. "github.com/miketheprogrammer/go-thrust/common"
)

func downloadFromUrl(url, filepath, version string) (string, error) {
	url = strings.Replace(url, "$V", version, 2)
	fileName := strings.Replace(filepath, "$V", version, 1)
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
			case <-time.After(time.Second):
				if Log.LogInfo() {
					fmt.Print(".")
				}
			}

		}
	}()

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return "", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return "", err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return "", err
	}
	quit <- 1

	fmt.Println(n, "bytes downloaded.")

	return fileName, nil
}
