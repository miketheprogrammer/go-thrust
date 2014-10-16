thrust-go
=========

Official GoLang Thrust Client for Thrust (https://github.com/breach/thrust) to be used in Breach.cc(http://breach.cc) via Official Go Implementation @ (http://github.com/miketheprogrammer/breach-go)

Notes for me while I work on this Project:

Running on OSX 10.9
==================
Included in vendor are the binaries
Start thrust using
```bash
./vendor/darwin/10.9/ThrustShell.app/Contents/MacOS/ThrustShell -socket-path=/tmp/
thrust.sock
# This is no longer needed, Thrust-go will detect darwin as your runtime, and try to spawn this binary.
# if that does not work run the aformentioned binary and then run thrust with
# go run src/main.go -socket=/tmp/thrust.sock -disable-auto-loader=true
```
Start Thrust Go using
```bash
go run src/main.go -socket=/tmp/thrust.sock
```


Building from source
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
 ./out/Debug/ThrustShell.app/Contents/MacOS/ThrustShell -socket-path={somepathhere}


From here down only works while I have  func main in the repo which will be removed soon
Now back in the thrust-go directory

go build -o thrust src/main.go

./thrust -socket={socketaddrusedinstartingthrusthere}


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



Extra Notes
================
Example Menu State @ during set_application_menu

```javascript
{
    "target_id": 2,
    "awaiting_responses": [
        {
            "_id": 0,
            "_action": "",
            "_method": "set_application_menu",
            "_args": {
                "size": {},
                "menu_id": 2,
                "value": false
            }
        }
    ],
    "ready": true,
    "Displayed": false,
    "items": [
        {
            "command_id": 2,
            "label": "Root",
            "type": "item"
        },
        {
            "command_id": 1,
            "label": "File",
            "submenu": {
                "target_id": 3,
                "ready": true,
                "Displayed": false,
                "items": [
                    {
                        "command_id": 3,
                        "label": "Open",
                        "type": "item"
                    },
                    {
                        "command_id": 4,
                        "label": "Close",
                        "type": "item"
                    },
                    {
                        "type": "separator"
                    },
                    {
                        "command_id": 7,
                        "label": "CheckList",
                        "submenu": {
                            "target_id": 4,
                            "ready": true,
                            "Displayed": false,
                            "items": [
                                {
                                    "command_id": 5,
                                    "label": "Do 1",
                                    "type": "check",
                                    "checked": true
                                },
                                {
                                    "type": "separator"
                                },
                                {
                                    "command_id": 6,
                                    "label": "Do 2",
                                    "type": "check",
                                    "checked": true
                                }
                            ]
                        }
                    }
                ]
            }
        }
    ]
}
```

