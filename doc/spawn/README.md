# spawn
--
    import "github.com/miketheprogrammer/go-thrust/spawn"


## Usage

#### func  GetDownloadUrl

```go
func GetDownloadUrl() string
```
GetDownloadUrl returns the interpolatable version of the Thrust download url
Differs between builds based on OS

#### func  GetExecutablePath

```go
func GetExecutablePath(base string) string
```
GetExecutablePath returns the path to the Thrust Executable Differs between
builds based on OS

#### func  GetThrustDirectory

```go
func GetThrustDirectory(base string) string
```
GetThrustDirectory returns the Directory where the unzipped thrust contents are.
Differs between builds based on OS

#### func  SpawnThrustCore

```go
func SpawnThrustCore(dir string) (io.ReadCloser, io.WriteCloser)
```
