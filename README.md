thrust-go
=========

Official GoLang Thrust Client for Thrust (https://github.com/breach/thrust) to be used in Breach.cc(http://breach.cc) via Official Go Implementation @ (http://github.com/miketheprogrammer/breach-go)

Running on OSX 10.9
==================
Included in vendor are the binaries
Start thrust using
```bash
./vendor/darwin/10.9/ThrustShell.app/Contents/MacOS/ThrustShell -socket-path=/tmp/
thrust.sock
# This is no longer needed, Thrust-go will detect darwin as your runtime, and try to spawn this binary.
# if that does not work run the aformentioned binary and then run thrust with
# go run demo.go -socket=/tmp/thrust.sock -disable-auto-loader=true
```
Start Thrust Go using
```bash
go run demo.go -socket=/tmp/thrust.sock
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
```bash
./scripts/bootstrap.py 
```
Then generate build files with:

```bash
GYP_GENERATORS=ninja gyp --depth . thrust_shell.gyp
```
Finally run ninja:

```bash
ninja -C out/Debug thrust_shell
```

Now you can run thrust 
I run it on OSX with the command
```bash
 ./out/Debug/ThrustShell.app/Contents/MacOS/ThrustShell -socket-path={somepathhere}
```


From here down only works while I have  func main in the repo which will be removed soon
Now back in the thrust-go directory

```bash
go build -o thrust src/main.go
./thrust -socket={socketaddrusedinstartingthrusthere}
```


Roadmap to v1.0 :
================
Please note Complete Support will never be toggled until Thrust core is stable.

- [X] Queue Requests prior to Object being created to matain a synchronous looking API without the need for alot of state checks
- [ ] Remove overuse of pointers in structs where modification will not take place
- [ ] Add Window Support
  - [X] Basic Support
  - [X] Refactor Connection usage
  - [ ] Complete Support 

- [ ] Add Menu Support
  - [X] Basic Support
  - [X] Refactor Connection usage
  - [ ] Complete Support

- [ ] Add Session Support
  - [ ] Basic Support
  - [ ] Complete Support

- [X] Implement Package Connection

- [x] Seperate out in to packages other than main
  - [X] Package Window
  - [X] Package Menu
  - [X] Package Commands
  - [X] Package Spawn

- [ ] Remove func Main as this is a Library
  - [ ] Should use Tests instead

- [X] Refactor how Dispatching occurs
  - [X] We should not have to manually dispatch, there should be a registration method 

- [ ] Refactor menu.SetChecked to accept a nillable menu item pointer, so we dont have to waste resources finding the item in the Tree

This thrust client exposes enough Methods to be fairly forwards compatible even without adding new helper methods. The beauty of working with a stable JSON RPC Protocol is that most methods are just helpers around build that data structure.

Helper methods receive the UpperCamelCase version of their relative names in Thrust.

i.e. insert_item_at == InsertItemAt



Extra Notes (This section quickly gets outdated, ill try to keep it updated, but it really just for reference)
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

