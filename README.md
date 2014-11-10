go-thrust
=========
Current Go-Thrust Version 0.2.4
Current Thrust Version 0.7.4

Official GoLang Thrust Client for Thrust (https://github.com/breach/thrust)

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
mkdir -f release/
mkdir -f release/go-thrust
go build -o release/go-thrust/Thrust demo.go
./release/go-thrust/Thrust -basedir=./release/go-thrust
```

Command Line Switches
========================
```bash
-basedir=string (DEMO ONLY)
    #where you want the core files to be stored.
    #on initial run, all files 

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

DOCUMENTATION
================
* [Index](https://github.com/miketheprogrammer/go-thrust/tree/master/doc)

Current Platform Specific Cases:
================
Currently Darwin handles Application Menus different from other systems.
The "Root" menu is always named the project name. However on linux, the name is whatever you provide. This can be seen in demo.go, and can be alleviated by trying to use the same name, or by using different logic for both cases. I make no attempt to unify it at the library level, because the user deserves the freedom to do this themselves.

Roadmap to v1.0 :
================
Please note Complete Support *may* not be toggled until Thrust core is stable.

- [ ] Add kiosk support (core.0.7.3)
- [X] Queue Requests prior to Object being created to matain a synchronous looking API without the need for alot of state checks
- [ ] Remove overuse of pointers in structs where modification will not take place
- [ ] Add Window Support
  - [X] Basic Support
  - [X] Refactor Connection usage
  - [ ] Complete Support 
    - Accessors (core.0.7.3)
    - Events (core.0.7.3)

- [ ] Add Menu Support
  - [X] Basic Support
  - [X] Refactor Connection usage
  - [ ] Complete Support

- [ ] Add Session Support
  - [X] Basic Support
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

- [X] Refactor CallWhen* methods, Due to the nature of using GoRoutines, there is the chance that calls will execute out of the original order they were intended.

- [X] Create a script to autodownload binaries

- [X] Refactor Logging

- [ ] SubMenus need order preservation

- [X] vendor folders need versioning

- [X] Need to fix Pathing for autoinstall and autorun. Relative paths will not work for most use cases.

This thrust client exposes enough Methods to be fairly forwards compatible even without adding new helper methods. The beauty of working with a stable JSON RPC Protocol is that most methods are just helpers around build that data structure.

Helper methods receive the UpperCamelCase version of their relative names in Thrust.

i.e. insert_item_at == InsertItemAt


Using Go-Thrust as a Library and bundling your final release
==========================
To use go-thrust as a library, simple use the code in the same way you would use any GoLang library.

- More info to come on Prepping releases.



Extra Notes (This section quickly gets outdated, ill try to keep it updated, but it really just for reference)
================

- Future Architecture
There needs to be consideration into how an application will be created.
While we should leave the core API's open enough to allow different architectures, there should be a primary guiding one.

The UI Layer
    - Should we do single page app style? using a framework.
        - Probably more responsive
    - Or should it be driven by templates in GO
        - Of course has the cost of network, and rebuilding socket connections on every load.
        - Benefit: When webview tag is added, various components can be registered, rather than being a single page app.





```

