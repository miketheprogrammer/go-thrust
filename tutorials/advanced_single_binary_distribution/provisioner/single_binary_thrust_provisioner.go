package provisioner

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudspace/go-thrust/lib/spawn"
)

/*
Single Binary Distribution
*/
type SingleBinaryThrustProvisioner struct{}

/*
NewSingleBinaryThrustProvisioner instantiates an SingleBinaryThrustProvisioner

go:generate go-bindata -pkg $GOPACKAGE -o vendor.go vendor/
*/
func NewSingleBinaryThrustProvisioner() SingleBinaryThrustProvisioner {
	return SingleBinaryThrustProvisioner{}
}

/*
Provisions a thrust environment based on settings.
*/
func (sbtp SingleBinaryThrustProvisioner) Provision() error {
	err := sbtp.extractToPath(spawn.GetDownloadPath())
	if err != nil {
		return err
	}
	return spawn.Bootstrap()
}

func (sbtp SingleBinaryThrustProvisioner) extractToPath(filepath string) error {
	data, err := Asset("vendor/thrust-v0.7.6-darwin-x64.zip")
	if err != nil {
		fmt.Println("Error accessing thrust bindata")
		return err
	}
	fmt.Println("No error accessing thrust bindata")
	fmt.Println("Writing to", filepath)
	err = ioutil.WriteFile(filepath, data, os.ModePerm)
	if nil != err {
		return err
	}
	return nil
	// return spawn.UnzipExecutable(filepath)
}
