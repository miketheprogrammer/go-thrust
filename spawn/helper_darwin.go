package spawn

import (
	. "github.com/miketheprogrammer/go-thrust/common"
)

func GetThrustDirectory(base string) string {
	return base + "/vendor/darwin/x64/v" + THRUST_VERSION
}

func GetExecutablePath(base string) string {
	return GetThrustDirectory(base) + "/ThrustShell.app/Contents/MacOS/ThrustShell"
}

func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-darwin-x64.zip"
}
