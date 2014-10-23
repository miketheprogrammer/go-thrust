package spawn

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/miketheprogrammer/go-thrust/common"
)

var base string

/*
SetBaseDirectory sets the base directory used in the other helper methods
*/
func SetBaseDirectory(b string) {
	base = b
}

/*
GetThrustDirectory returns the Directory where the unzipped thrust contents are.
Differs between builds based on OS
*/
func GetThrustDirectory() string {
	return base + "/vendor/darwin/x64/v" + common.THRUST_VERSION
}

/*
GetExecutablePath returns the path to the Thrust Executable
Differs between builds based on OS
*/
func GetExecutablePath() string {
	return GetThrustDirectory() + "/" + common.ApplicationName + ".app/Contents/MacOS/" + common.ApplicationName
}

/*
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS
*/
func GetDownloadUrl() string {
	return "https://github.com/breach/thrust/releases/download/v$V/thrust-v$V-darwin-x64.zip"
}

/*
SetThrustApplicationTitle sets the title in the Info.plist. This method only exists on Darwin.
*/
func Bootstrap() {
	if executableNotExist() == true {
		prepareExecutable()
		prepareInfoPropertiesListTemplate()
		writeInfoPropertiesList()
	}

}

func executableNotExist() bool {
	_, err := os.Stat(GetExecutablePath())
	return os.IsNotExist(err)
}

func prepareExecutable() {
	downloadFromUrl(GetDownloadUrl(), common.THRUST_VERSION)
	unzip(strings.Replace("/tmp/$V", "$V", common.THRUST_VERSION, 1), GetThrustDirectory())
	os.Rename(GetThrustDirectory()+"/ThrustShell.app/Contents/MacOS/ThrustShell", GetThrustDirectory()+"/ThrustShell.app/Contents/MacOS/"+common.ApplicationName)
	os.Rename(GetThrustDirectory()+"/ThrustShell.app", GetThrustDirectory()+"/"+common.ApplicationName+".app")
}

func prepareInfoPropertiesListTemplate() bool {
	plistPath := getInfoPropertiesListDirectory() + "/Info.plist"

	if _, err := os.Stat(plistPath + ".tmpl"); os.IsNotExist(err) {
		plist, err := ioutil.ReadFile(plistPath)

		if err != nil {
			fmt.Println(err)
			return false
		}

		plistTmpl := strings.Replace(string(plist), "ThrustShell", "$$", -1)
		plistTmpl = strings.Replace(plistTmpl, "org.breach.$$", "org.breach.ThrustShell", 1)
		plistTmpl = strings.Replace(plistTmpl, "$$Application", "ThrustShellApplication", 1)
		//func WriteFile(filename string, data []byte, perm os.FileMode) error

		err = ioutil.WriteFile(plistPath+".tmpl", []byte(plistTmpl), 0775)

		return true
	}

	return true
}

func writeInfoPropertiesList() {
	plistPath := getInfoPropertiesListDirectory() + "/Info.plist"
	if prepareInfoPropertiesListTemplate() == true {
		plistTmpl, err := ioutil.ReadFile(plistPath + ".tmpl")

		if err != nil {
			fmt.Println(err)
			return
		}

		plist := strings.Replace(string(plistTmpl), "$$", common.ApplicationName, -1)

		err = ioutil.WriteFile(plistPath, []byte(plist), 0775)
	} else {
		fmt.Println("Could not change title")
	}
}

func getInfoPropertiesListDirectory() string {
	return GetThrustDirectory() + "/" + common.ApplicationName + ".app/Contents"
}
