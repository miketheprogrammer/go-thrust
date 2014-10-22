# spawn
--
    import "github.com/miketheprogrammer/go-thrust/spawn"


## Usage

#### func  SpawnThrustCore

```go
func SpawnThrustCore(dir string) (io.ReadCloser, io.WriteCloser)
```
The SpawnThrustCore method is a bootstrap and run method. It will try to detect
an installation of thrust, if it cannot find it it will download the version of
Thrust detailed in the "common" package. Once downloaded, it will launch a
process. Go-Thrust and all *-Thrust packages communicate with Thrust Core via
Stdin/Stdout. using -log=debug as a command switch will give you the most
information about what is going on. -log=info will give you notices that stuff
is happening. Any log level higher than that will output nothing.
