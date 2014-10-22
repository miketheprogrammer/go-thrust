package spawn

import (
	. "github.com/miketheprogrammer/go-thrust/common"
)

func GetThrustDirectory(based string) string {
	return base + "/vendor/linux/x64/v" + THRUST_VERSION
}

func GetExecutablePath(based string) string {
	return GetThrustDirectory(base) + "/thrust_shell"
}

func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-linux-x64.zip"
}
