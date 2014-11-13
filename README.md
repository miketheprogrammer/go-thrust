go-thrust
=========
Current Go-Thrust Version 0.3.0 
Current Thrust Version 0.7.4

![Logo Thrust](http://i.imgur.com/DwFKI0J.png)
Official GoLang Thrust Client for Thrust (https://github.com/breach/thrust)

Quick Start, see it in action
==================
On Linux/Darwin

Start Go Thrust using source and GoLang runtime
```bash
./install-go-deps.sh #installs current dependencies
make tutorials
# then execute one of the many tutorials
#./tutorial/bin/basic_window
#./tutorial/bin/basic_menu
#./tutorial/bin/basic_menu_events
#./tutorial/bin/basic_multiple_windows
#./tutorial/bin/advanced_session
#./tutorial/bin/basic_webserver_app
```

On Windows just use go run
```bash
go run tutorial/basic_window.go
# ... etcetera
```

Forward:
==================
Go-Thrust is a cross platform GUI Application development base. It aims to provide the
essentials to create a cross platform application using web technologies such as, (HTML,CSS,JS),
as well as new technologies like Websockets, Webview, etcetera. Go-Thrust is the supported Go Library for accessing the underlying technology Thrust. Thrust builds upon the efforts of predecessors like ExoBrowser, Node-Webkit, and Atom-Shell to create an easily buildable Webbrowser base. Thrust consists essentially of a C/C++ implementation of the Chromium Content Library (Blink/V8) and abstracts away most of the nitty gritty platform specific implementations.

While Go Thrust and Thrust are in their infancy, they already provide some pretty nice features.
Expect many more to come.

Thrust exposes all of this beautifully over an Stdin/Stdout pipe speaking a JSONRPC protocol.

DOCUMENTATION (Docs and Tutorials)
================
* [Index](https://github.com/miketheprogrammer/go-thrust/tree/master/doc)

Current Platform Specific Cases:
================
Currently Darwin handles Application Menus different from other systems.
The "Root" menu is always named the project name. However on linux, the name is whatever you provide. This can be seen in demo.go, and can be alleviated by trying to use the same name, or by using different logic for both cases. I make no attempt to unify it at the library level, because the user deserves the freedom to do this themselves.

This thrust client exposes enough Methods to be fairly forwards compatible even without adding new helper methods. The beauty of working with a stable JSON RPC Protocol is that most methods are just helpers around build that data structure.

Helper methods receive the UpperCamelCase version of their relative names in Thrust.

i.e. insert_item_at == InsertItemAt

Please note **** That the intended use case of Application Menus is to only support 
OSX and Unity/X11 global menu bars. This means that you should implement most menus in html and javascript, using IPC/RPC to communicate with the host application. The side effect is primarily that Windows, and certain unix/linux systems will not load ApplicationMenus 


Using Go-Thrust as a Library and bundling your final release
==========================
To use go-thrust as a library, simple use the code in the same way you would use any GoLang library.

- More info to come on Prepping releases.

Current Platform Targets
================
Linux, Darwin, Windows.
Windows - Stable as of 0.3.0

Unfortunately because Thrust is built around C and Go its not exactly portable to Android and IOS, this may be possible in the future, but not yet. However, since we are primarily building a web app, we may be able to do something at a later time. For instance, there is work being done on allowing java to run go libraries. This would allow us to hopefully covert our application menus, to maybe just an html menu at the top, and just serve the page, instead of using Thrust. Thats all for the future however, for now, enjoy CPD for Linux,Darwin,Windows.


The Future of Go-Thrust
================
Any user of Go-Thrust should feel free to contribute Ideas, Code, anything that can help move this project in the right direction. The ideal goal is to stay as much out of the way, but still provide a useful system.

Some questions and thoughts: 
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

