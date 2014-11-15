package spawn

import (
	"os"
	"strings"

	"github.com/miketheprogrammer/go-thrust/common"
)

/*
GetThrustDirectory returns the Directory where the unzipped thrust contents are.
Differs between builds based on OS
*/
func GetThrustDirectory() string {
	return base + "./vendor/linux/x64/v" + common.THRUST_VERSION
}

/*
GetExecutablePath returns the path to the Thrust Executable
Differs between builds based on OS
*/
func GetExecutablePath() string {
	return GetThrustDirectory() + "/thrust_shell"
}

/*
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS
*/
func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-linux-x64.zip"
}

/*
SetThrustApplicationTitle sets the title in the Info.plist. This method only exists on Darwin.
*/
func Bootstrap() {
	if executableNotExist() == true {
		prepareExecutable()
	}
}

func executableNotExist() bool {
	_, err := os.Stat(GetExecutablePath())
	return os.IsNotExist(err)
}

func prepareExecutable() {
	downloadFromUrl(GetDownloadUrl(), "/tmp/$V", common.THRUST_VERSION)
	unzip(strings.Replace("/tmp/$V", "$V", common.THRUST_VERSION, 1), GetThrustDirectory())
}
