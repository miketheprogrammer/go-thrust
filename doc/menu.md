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


#### func (*Menu) AddCheckItem

```go
func (menu *Menu) AddCheckItem(commandID uint, label string)
```
Add a CheckItem to both the internal representation of menu and the external
representation of menu

#### func (*Menu) AddItem

```go
func (menu *Menu) AddItem(commandID uint, label string)
```
Add a MenuItem to both the internal representation of menu and the external
representation of menu

#### func (*Menu) AddRadioItem

```go
func (menu *Menu) AddRadioItem(commandID uint, label string, groupID uint)
```
Add a RadioItem to both the internal representation of menu and the external
representation of menu

#### func (*Menu) AddSeparator

```go
func (menu *Menu) AddSeparator()
```
Add a Seperator to both the internal representation of menu and the external
representation of menu

#### func (*Menu) AddSubmenu

```go
func (menu *Menu) AddSubmenu(commandID uint, label string, child *Menu)
```
Add a SubMenu to both the internal representation of menu and the external
representation of menu

#### func (*Menu) AttachToWindow

```go
func (menu *Menu) AttachToWindow(w *window.Window)
```
On Linux and Windows systems, Attach the menu to a window

#### func (*Menu) Call

```go
func (menu *Menu) Call(command *Command)
```

#### func (*Menu) CallWhenChildStable

```go
func (menu *Menu) CallWhenChildStable(command *Command, child *Menu)
```

#### func (*Menu) CallWhenDisplayed

```go
func (menu *Menu) CallWhenDisplayed(command *Command)
```

#### func (*Menu) CallWhenReady

```go
func (menu *Menu) CallWhenReady(command *Command)
```

#### func (*Menu) CallWhenTreeStable

```go
func (menu *Menu) CallWhenTreeStable(command *Command)
```

#### func (*Menu) Create

```go
func (menu *Menu) Create(sendChannel *connection.In)
```

#### func (*Menu) DispatchResponse

```go
func (menu *Menu) DispatchResponse(reply CommandResponse)
```

#### func (*Menu) HandleError

```go
func (menu *Menu) HandleError(reply CommandResponse)
```

#### func (*Menu) HandleEvent

```go
func (menu *Menu) HandleEvent(reply CommandResponse)
```

#### func (*Menu) HandleReply

```go
func (menu *Menu) HandleReply(reply CommandResponse)
```

#### func (*Menu) IsStable

```go
func (menu *Menu) IsStable() bool
```
A menu is stable if and only if, it is Ready (meaning it was created
successfully) and it has no Commands awaiting Responses.

#### func (*Menu) IsTarget

```go
func (menu *Menu) IsTarget(targetId uint) bool
```

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
Recursively searches the Menu Tree for an item with the commandID Returns the
first found match. A proper menu should not reuse commandID's

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

#### func (*Menu) SendThread

```go
func (menu *Menu) SendThread()
```

#### func (*Menu) SetApplicationMenu

```go
func (menu *Menu) SetApplicationMenu()
```
On Darwin systems, Set the application menu in the UI

#### func (*Menu) SetChecked

```go
func (menu *Menu) SetChecked(commandID uint, checked bool)
```
Checks or Unchecks a CheckItem in the UI

#### func (*Menu) SetEnabled

```go
func (menu *Menu) SetEnabled(commandID uint, enabled bool)
```
Enables or Disables an item in the UI

#### func (*Menu) SetSendChannel

```go
func (menu *Menu) SetSendChannel(sendChannel *connection.In)
```

#### func (*Menu) SetVisible

```go
func (menu *Menu) SetVisible(commandID uint, visible bool)
```
Enables or Disables an item in the UI

#### func (*Menu) ToggleRadio

```go
func (menu *Menu) ToggleRadio(commandID, groupID uint, checked bool)
```
Checks or Unchecks a CheckItem in the UI

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
