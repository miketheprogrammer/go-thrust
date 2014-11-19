package platform

import (
	"fmt"

	"github.com/miketheprogrammer/go-thrust/common"
)

type Platform struct {
	Modules []Module
}

type Module struct {
	Path string
	Data []byte
}

func NewPlatform() *Platform {
	platform := &Platform{}
	platform.LoadModuleJS("platform.js")
	return platform
}

func (p *Platform) LoadModuleJS(path string) {
	// Asset provided by go-bindata
	bytes, err := Asset(path)
	if err != nil {
		common.Log.Alert(err)
	}
	fmt.Println(bytes, err)
	p.Modules = append(p.Modules, Module{
		Path: path,
		Data: bytes,
	})
}
