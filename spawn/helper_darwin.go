package spawn

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Unknwon/cae/zip"
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

func GetAppDirectory() string {
	return base + "/vendor/darwin/x64/v" + common.THRUST_VERSION + "/" + common.ApplicationName + ".app"
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

/*
executableNotExist checks if the executable does not exist
*/
func executableNotExist() bool {
	_, err := os.Stat(GetExecutablePath())
	return os.IsNotExist(err)
}

/*
prepareExecutable dowloads, unzips and does alot of other magic to prepare our thrust core build.
*/
func prepareExecutable() {
	path := downloadFromUrl(GetDownloadUrl(), common.THRUST_VERSION)
	//unzip(strings.Replace("/tmp/$V", "$V", common.THRUST_VERSION, 1), GetThrustDirectory())
	zip.ExtractTo(path, GetThrustDirectory())

	os.Rename(GetThrustDirectory()+"/ThrustShell.app/Contents/MacOS/ThrustShell", GetThrustDirectory()+"/ThrustShell.app/Contents/MacOS/"+common.ApplicationName)
	os.Rename(GetThrustDirectory()+"/ThrustShell.app", GetThrustDirectory()+"/"+common.ApplicationName+".app")

	applySymlinks()
}

/*
ApplySymLinks exists because our unzip utility does not respect deferred symlinks. It applies all the neccessary symlinks to make the thrust core exe connect to the thrust core libs.
*/
func applySymlinks() {
	fmt.Println("Applying Symlinks")
	fmt.Println(
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current"),
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Frameworks"),
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Libraries"),
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Resources"),
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/ThrustShell Framework"),
		os.Remove(GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/Libraries"))

	fmt.Println(
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/A",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current"),
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/Frameworks",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Frameworks"),
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/Libraries",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Libraries"),
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/Resources",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Resources"),
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/ThrustShell Framework",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/ThrustShell Framework"),
		os.Symlink(
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/A/Libraries/Libraries",
			GetAppDirectory()+"/Contents/Frameworks/ThrustShell Framework.framework/Versions/Current/Libraries"))

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
