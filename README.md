thrust-go
=========

GoLang Thrust Client for Breach.cc 

Notes for me while I work on this Project:

Current command to run is:
```bash
go run src/* -socket=/var/folders/nx/lqsp2r_93nlbdt_7mxkwrjr00000gn/T/_thrust_shell.sock
```

Building bleeding edge from Source
======
Get Chromium DepotTools from here
http://www.chromium.org/developers/how-tos/install-depot-tools
Make sure to follow instructions for setting path.

Clone Thrust Core from here
https://github.com/breach/thrust/

From the thrust core directory run 
First run the boostrap script:

./scripts/bootstrap.py 
Then generate build files with:

GYP_GENERATORS=ninja gyp --depth . thrust_shell.gyp
Finally run ninja:

ninja -C out/Debug thrust_shell

Now you can run thrust 
I run it on OSX with the command
 ./out/Debug/ThrustShell.app/Contents/MacOS/ThrustShell 

Look for a line similar to this
[19883:1015/143211:INFO:api_server.cc(388)] Listening on: /var/folders/nx/lqsp2r_93nlbdt_7mxkwrjr00000gn/T/_thrust_shell.sock

And copy that socket address

From here down only works while I have  func main in the repo which will be removed soon
Now back in the thrust-go directory

go build -o thrust src/* 

./thrust -socket={socketaddrhere}


Todo:
================

- [ ] Queue Requests prior to Object being created to matain a synchronous looking API without the need for alot of state checks
- [ ] Remove overuse of pointers in structs where modification will not take place
- [ ] Add Window Support
  - [X] Basic Support
  - [ ] Complete Support 

- [ ] Add Menu Support
  - [X] Basic Support
  - [ ] Complete Support

- [ ] Add Session Support
  - [ ] Basic Support
  - [ ] Complete Support

- [ ] Seperate out in to packages other than main
  - [ ] Package Window
  - [ ] Package Menu
  - [ ] Package Session
  - [ ] Package Commands
  - [ ] Package Core - Import for all Window, Menu, Session, Commands 

- [ ] - Remove func Main as this is a Library
  - [ ] Should use Tests instead


This thrust client exposes enough Methods to be fairly forwards compatible even without adding new helper methods. The beauty of working with a stable JSON RPC Protocol is that most methods are just helpers around build that data structure.

Helper methods receive the UpperCamelCase version of their relative names in Thrust.

i.e. insert_item_at == InsertItemAt



