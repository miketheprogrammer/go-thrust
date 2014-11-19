package web

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/miketheprogrammer/go-thrust/web/platform"
)

type WebHandler struct {
	Platform *platform.Platform
}

func NewWebHandler() WebHandler {
	wh := WebHandler{}
	wh.Platform = platform.NewPlatform()
	return wh
}

func (wh WebHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	writer := bufio.NewWriter(w)
	for _, module := range wh.Platform.Modules {
		fmt.Println(module)
		for _, byte_ := range module.Data {
			fmt.Println(byte_)
			writer.WriteByte(byte_)
		}
	}
	return
}
