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
