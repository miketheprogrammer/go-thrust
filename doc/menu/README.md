# menu
--
    import "github.com/miketheprogrammer/go-thrust/menu"


## Usage

#### type ChildCommand

```go
type ChildCommand struct {
	Command *commands.Command
	Child   *Menu
}
```


#### type Item

```go
type Item interface {
	IsSubMenu() bool
	IsCommandId() bool
	Handle()
}
```


#### type Menu

```go
type Menu struct {
	TargetID         uint                                                 `json:"target_id,omitempty"`
	WaitingResponses []*Command                                           `json:"awaiting_responses,omitempty"`
	CommandQueue     []*Command                                           `json:"command_queue,omitempty"`
	Ready            bool                                                 `json:"ready"`
	Displayed        bool                                                 `json:"displayed"`
	Parent           *Menu                                                `json:"-"`
	Children         []*Menu                                              `json:"-"`
	Items            []*MenuItem                                          `json:"items,omitempty"`
	EventRegistry    []uint                                               `json:"events,omitempty"`
	SendChannel      *connection.In                                       `json:"-"`
	Sync             MenuSync                                             `jons:"-"`
	ReplyHandlers    map[uint]func(reply CommandResponse, item *MenuItem) `json:"_"`
}
```

Menu is the basic object for creating and working with Menu's Provides all the
necessary attributes and methods to work with asynchronous calls to the menu
API. The TargetID is assigned by ThrustCore, so on init of this object, there is
no TargetID. A Goroutine is dispatched to get the targetID.

#### func (*Menu) AddCheckItem

```go
func (menu *Menu) AddCheckItem(commandID uint, label string)
```
AddCheckItem adds a CheckItem to both the internal representation of menu and
the external representation of menu

#### func (*Menu) AddItem

```go
func (menu *Menu) AddItem(commandID uint, label string)
```
AddItem adds a MenuItem to both the internal representation of menu and the
external representation of menu

#### func (*Menu) AddRadioItem

```go
func (menu *Menu) AddRadioItem(commandID uint, label string, groupID uint)
```
AddRadioItem adds a RadioItem to both the internal representation of menu and
the external representation of menu

#### func (*Menu) AddSeparator

```go
func (menu *Menu) AddSeparator()
```
AddSeperator adds a Seperator Item to both the internal representation of menu
and the external representation of menu.

#### func (*Menu) AddSubmenu

```go
func (menu *Menu) AddSubmenu(commandID uint, label string, child *Menu)
```
AddSubmenu adds a SubMenu to both the internal representation of menu and the
external representation of menu

#### func (*Menu) Call

```go
func (menu *Menu) Call(command *Command)
```
Call turns a Command into an action:call, there are two main types of Actions
for outgoing commands, create/call. There may be more added later.

#### func (*Menu) CallWhenChildStable

```go
func (menu *Menu) CallWhenChildStable(command *Command, child *Menu)
```
CallWhenChildStable queues up "Calls" to go out only when the state of the Child
is Stable. Stable means that the child is Ready and has no AwaitingResponses

#### func (*Menu) CallWhenDisplayed

```go
func (menu *Menu) CallWhenDisplayed(command *Command)
```
CallWhenDisplayed queues up "Calls" to go out only when the menu is Displayed

#### func (*Menu) CallWhenReady

```go
func (menu *Menu) CallWhenReady(command *Command)
```
CallWhenReady queues up "Calls" to go out only when the Menu State is "Ready"

#### func (*Menu) CallWhenTreeStable

```go
func (menu *Menu) CallWhenTreeStable(command *Command)
```
CallWhenTreeStable queues up "Calls" to go out only when the state of the menu
is Stable. Stable means that the menu is Ready and has no AwaitingResponses

#### func (*Menu) Create

```go
func (menu *Menu) Create()
```
Create a new menu object. Dispatches a call to ThrustCore to generate the object
and return the new TargetID in a reply.

#### func (*Menu) DispatchResponse

```go
func (menu *Menu) DispatchResponse(reply CommandResponse)
```
DispatchResponse dispatches CommandResponses to the proper delegates (Error,
Event, Reply)

#### func (*Menu) HandleError

```go
func (menu *Menu) HandleError(reply CommandResponse)
```
HandleError is a handler for Error responses from ThrustCore This should be
changed to private as soon as API stabilizes.

#### func (*Menu) HandleEvent

```go
func (menu *Menu) HandleEvent(reply CommandResponse)
```
HandleEvent is a handler for Event responses from ThrustCore This should be
changed to private as soon as API stabilizes.

#### func (*Menu) HandleReply

```go
func (menu *Menu) HandleReply(reply CommandResponse)
```
HandleReply is a handler for Reply responses from ThrustCore This should be
changed to private as soon as API stabilizes.

#### func (*Menu) IsStable

```go
func (menu *Menu) IsStable() bool
```
IsStable returns the a boolean value indicating that the menu is Ready and has
no WaitingResponses

#### func (*Menu) IsTarget

```go
func (menu *Menu) IsTarget(targetId uint) bool
```
IsTarget checks if the current menu is the menu we are looking for.

#### func (*Menu) IsTreeStable

```go
func (menu *Menu) IsTreeStable() bool
```
A Menu Tree is considered stable if and only if its children nodes report that
they are stable. Function is recursive, so factor that in to performance

#### func (*Menu) ItemAtCommandID

```go
func (menu *Menu) ItemAtCommandID(commandID uint) *MenuItem
```
ItemAtCommandID recursively searches the Menu Tree for an item with the
commandID. Returns the first found match. A proper menu should not reuse
commandID's

#### func (*Menu) Popup

```go
func (menu *Menu) Popup(w *window.Window)
```
Popup creates a popup menu on the given window

#### func (Menu) PrintRecursiveWaitingResponses

```go
func (menu Menu) PrintRecursiveWaitingResponses()
```
DEBUG Functions

#### func (*Menu) RadioGroupAtGroupID

```go
func (menu *Menu) RadioGroupAtGroupID(groupID uint) []*MenuItem
```
Find all menu items that belong to group identified by groupID Not recursive, as
a group should be identified at the same level. Since it is not recursive you
can theoretically reuse a groupID but problems could creep up elsewhere, so
please use unique groupID for radio items

#### func (*Menu) RegisterEventHandlerByCommandID

```go
func (menu *Menu) RegisterEventHandlerByCommandID(commandID uint, handler func(reply CommandResponse, item *MenuItem))
```

#### func (*Menu) Send

```go
func (menu *Menu) Send(command *Command)
```
Send emits a Command over the Command SendChannel to be delivered to ThrustCore

#### func (*Menu) SendThread

```go
func (menu *Menu) SendThread()
```
SendThread is a Thread for Sending Commands based on current state of the Menu.
Some commands require other events in the system to have already taken place.
This thread ensures that you can run almost any command at anytime, and have it
take place in the correct order. This further insures that the underlying
ThrustCore api does not crash, do to improper api knowledge.

#### func (*Menu) SetApplicationMenu

```go
func (menu *Menu) SetApplicationMenu()
```
SetApplicationMenu sets the Application Menu on system that support global
application level menus such as x11, unity, darwin

#### func (*Menu) SetChecked

```go
func (menu *Menu) SetChecked(commandID uint, checked bool)
```
SetChecked Checks or Unchecks a CheckItem in the UI

#### func (*Menu) SetEnabled

```go
func (menu *Menu) SetEnabled(commandID uint, enabled bool)
```
SetEnabled sets whether or not a given item can receive actions via the UI.

#### func (*Menu) SetSendChannel

```go
func (menu *Menu) SetSendChannel(sendChannel *connection.In)
```
SetSendChannel is a helper Setter for SendChannel, in case we make it private in
the future. Use this for full forwards compatibility.

#### func (*Menu) SetVisible

```go
func (menu *Menu) SetVisible(commandID uint, visible bool)
```
SetVisible sets a boolean visibility attribute in the UI for a menu item with
the given commandID.

#### func (*Menu) ToggleRadio

```go
func (menu *Menu) ToggleRadio(commandID, groupID uint, checked bool)
```
ToggleRadio Checks or Unchecks a RadioItem in the UI. It is used by the default
event handler to turn on the expected item, and turn of other items in the
group.

#### type MenuItem

```go
type MenuItem struct {
	CommandID uint   `json:"command_id,omitempty"`
	Label     string `json:"label,omitempty"`
	GroupID   uint   `json:"group_id,omitempty"`
	SubMenu   *Menu  `json:"submenu,omitempty"`
	Type      string `json:"type,omitempty"`
	Checked   bool   `json:"checked"`
	Enabled   bool   `json:"enabled"`
	Visible   bool   `json:"visible"`
	Parent    *Menu  `json:"-"`
}
```


#### func (MenuItem) HandleEvent

```go
func (mi MenuItem) HandleEvent()
```

#### func (MenuItem) IsCheckItem

```go
func (mi MenuItem) IsCheckItem() bool
```

#### func (MenuItem) IsCommandID

```go
func (mi MenuItem) IsCommandID(commandID uint) bool
```

#### func (MenuItem) IsGroupID

```go
func (mi MenuItem) IsGroupID(groupID uint) bool
```

#### func (MenuItem) IsRadioItem

```go
func (mi MenuItem) IsRadioItem() bool
```

#### func (MenuItem) IsSubMenu

```go
func (mi MenuItem) IsSubMenu() bool
```

#### type MenuSync

```go
type MenuSync struct {
	/* Channels for singaling queues */
	ReadyChan       chan bool
	DisplayedChan   chan bool
	ChildStableChan chan uint
	TreeStableChan  chan bool

	/*Queues for preserving command order and Priority*/
	ReadyQueue     []*commands.Command
	DisplayedQueue []*commands.Command
	// Not Exactly a Queue, more of a priority queue. Send out the first child that is stable.
	ChildStableQueue []*ChildCommand
	TreeStableQueue  []*commands.Command

	/* Channels for control */
	QuitChan chan bool
}
```
