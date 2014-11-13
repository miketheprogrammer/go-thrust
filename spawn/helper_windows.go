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
	return base + "\\vendor\\windows\\ia32\\v" + common.THRUST_VERSION
}

/*
GetExecutablePath returns the path to the Thrust Executable
Differs between builds based on OS
*/
func GetExecutablePath() string {
	return GetThrustDirectory() + "\\thrust_shell.exe"
}

/*
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS
*/
func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-windows-ia32.zip"
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
	common.Log.Debug(os.Getenv("TEMP") + "\\$V")
	downloadFromUrl(GetDownloadUrl(), os.Getenv("TEMP")+"\\$V", common.THRUST_VERSION)
	unzip(strings.Replace(os.Getenv("TEMP")+"\\$V", "$V", common.THRUST_VERSION, 1), GetThrustDirectory())
}
