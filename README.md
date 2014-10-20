go-thrust
=========

Official GoLang Thrust Client for Thrust (https://github.com/breach/thrust) to be used in Breach.cc(http://breach.cc) via Official Go Implementation @ (http://github.com/miketheprogrammer/breach-go).

Running DEMO binary on OSX 10.9 or Linux (This is just for the demo, Thrust as it stands now is really a Library to be used in other applications.)
==================
Go Thrust will automatically install Thrust Core dependencies at runtime.


Start Go Thrust using binaries
```bash
Download the release for your distribution from the releases page.
unzip the contents.
run ./Thrust or the absolute path to the executable
```
Start Go Thrust using source and GoLang runtime
```bash
./install-go-deps.sh #installs current dependencies
make build.release
cd release/go-thrust/
./Thrust
```

Command Line Switches
========================
```bash
#LOGGING
#Inherits Switches from http://godoc.org/github.com/alexcesaro/log/stdlog
#Added below for ease, but for up to date info visit the page.
-log=loglevel
    #Log events at or above this level are logged.
-stderr=bool
    #Logs are written to standard error (stderr) instead of standard
    #output.
-flushlog=none
    #Until this level is reached nothing is output and logs are stored
    #in the memory. Once a log event is at or above this level, it
    #outputs all logs in memory as well as the future log events. This
    #feature should not be used with long-running processes.
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

- [ ] Refactor CallWhen* methods, Due to the nature of using GoRoutines, there is the chance that calls will execute out of the original order they were intended.

- [ ] Create a script to autodownload binaries

- [X] Refactor Logging

- [ ] SubMenus need order preservation

- [X] vendor folders need versioning

- [X] Need to fix Pathing for autoinstall and autorun. Relative paths will not work for most use cases.

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

