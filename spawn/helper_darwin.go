package spawn

import (
	. "github.com/miketheprogrammer/go-thrust/common"
)

/*
GetThrustDirectory returns the Directory where the unzipped thrust contents are.
Differs between builds based on OS
*/
func GetThrustDirectory(base string) string {
	return base + "/vendor/darwin/x64/v" + THRUST_VERSION
}

/*
GetExecutablePath returns the path to the Thrust Executable
Differs between builds based on OS
*/
func GetExecutablePath(base string) string {
	return GetThrustDirectory(base) + "/ThrustShell.app/Contents/MacOS/ThrustShell"
}

/*
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS
*/
func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-darwin-x64.zip"
}
