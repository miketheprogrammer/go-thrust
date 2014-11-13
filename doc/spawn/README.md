# spawn
--
    import "github.com/miketheprogrammer/go-thrust/spawn"


## Usage

#### func  Bootstrap

```go
func Bootstrap()
```
SetThrustApplicationTitle sets the title in the Info.plist. This method only
exists on Darwin.

#### func  GetAppDirectory

```go
func GetAppDirectory() string
```

#### func  GetDownloadUrl

```go
func GetDownloadUrl() string
```
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS

#### func  GetExecutablePath

```go
func GetExecutablePath() string
```
GetExecutablePath returns the path to the Thrust Executable Differs between
builds based on OS

#### func  GetThrustDirectory

```go
func GetThrustDirectory() string
```
GetThrustDirectory returns the Directory where the unzipped thrust contents are.
Differs between builds based on OS

#### func  Run

```go
func Run(autoDownloadEnabled bool) (io.ReadCloser, io.WriteCloser)
```
The SpawnThrustCore method is a bootstrap and run method. It will try to detect
an installation of thrust, if it cannot find it it will download the version of
Thrust detailed in the "common" package. Once downloaded, it will launch a
process. Go-Thrust and all *-Thrust packages communicate with Thrust Core via
Stdin/Stdout. using -log=debug as a command switch will give you the most
information about what is going on. -log=info will give you notices that stuff
is happening. Any log level higher than that will output nothing.

#### func  SetBaseDirectory

```go
func SetBaseDirectory(dir string) error
```
SetBaseDirectory sets the base directory used in the other helper methods
