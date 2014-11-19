package spawn

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/sadasant/go-thrust/common"
)

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
				Log.Critical(err)
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
