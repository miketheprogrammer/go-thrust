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

#### func  SetBaseDirectory

```go
func SetBaseDirectory(b string)
```
SetBaseDirectory sets the base directory used in the other helper methods

#### func  SpawnThrustCore

```go
func SpawnThrustCore(dir string) (io.ReadCloser, io.WriteCloser)
```
