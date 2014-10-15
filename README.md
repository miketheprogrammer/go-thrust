thrust-go
=========

GoLang Thrust Client for Breach.cc 

Notes for me while I work on this Project:

Current command to run is:
```bash
go run src/* -socket=/var/folders/nx/lqsp2r_93nlbdt_7mxkwrjr00000gn/T/_thrust_shell.sock
```

Todo:

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



