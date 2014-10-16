package menu

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	. "github.com/miketheprogrammer/thrust-go/src/commands"
	. "github.com/miketheprogrammer/thrust-go/src/common"
)

type Menu struct {
	TargetID         int         `json:"target_id,omitempty"`
	WaitingResponses []*Command  `json:"awaiting_responses,omitempty"`
	CommandQueue     []*Command  `json:"command_queue,omitempty"`
	Conn             net.Conn    `json:"-"`
	Ready            bool        `json:"ready"`
	Displayed        bool        `json:"displayed"`
	Parent           *Menu       `json:"-"`
	Children         []*Menu     `json:"-"`
	Items            []*MenuItem `json:"items,omitempty"`
	EventRegistry    []int       `json:"events,omitempty"`
}

func (menu *Menu) Create(conn net.Conn) {
	menuCreate := Command{
		Action:     "create",
		ObjectType: "menu",
	}
	menu.WaitingResponses = append(menu.WaitingResponses, &menuCreate)
	menu.Send(&menuCreate, conn)
}

func (menu *Menu) IsTarget(targetId int) bool {
	return targetId == menu.TargetID
}
func (menu *Menu) HandleError(reply CommandResponse, conn net.Conn) {

}

func (menu *Menu) HandleEvent(reply CommandResponse, conn net.Conn) {
	for _, item := range menu.Items {
		fmt.Println("Looking for item to handle event")
		if reply.Event.CommandID == item.CommandID {
			fmt.Println("Event", item.CommandID, "Handled With Flags", reply.Event.EventFlags, "With Type", item.Type)
			item.HandleEvent(conn)
			return
		}
	}
}

func (menu *Menu) HandleReply(reply CommandResponse, conn net.Conn) {

	for k, v := range menu.WaitingResponses {
		if v.ID != reply.ID {
			continue
		}
		fmt.Println("MENU(", menu.TargetID, ")::Handling Response", reply)
		removeAt := func(k int) {
			if len(menu.WaitingResponses) > 1 {
				menu.WaitingResponses = menu.WaitingResponses[:k+copy(menu.WaitingResponses[k:], menu.WaitingResponses[k+1:])]
			} else {
				menu.WaitingResponses = []*Command{}
			}
		}
		defer removeAt(k)

		if menu.TargetID == 0 && v.Action == "create" {
			//Assume we have a reply to action:create
			if reply.Result.TargetID != 0 {
				menu.TargetID = reply.Result.TargetID
				fmt.Println("Received TargetID", "\nSetting Ready State")
				menu.Ready = true
			}
			for i, _ := range menu.CommandQueue {
				menu.CommandQueue[i].TargetID = menu.TargetID
				menu.Send(menu.CommandQueue[i], conn)
			}
			// Reinitialize empty command queue, and allow gc.
			menu.CommandQueue = []*Command{}
			return
		}

		if v.Action == "call" && v.Method == "set_application_menu" {
			fmt.Println("Received reply to set_application_menu", "Setting Menu Displayed to True")
			menu.setDisplayed(true)
		}

	}
}

func (menu *Menu) setDisplayed(displayed bool) {
	menu.Displayed = displayed

	for _, child := range menu.Items {
		if child.IsSubMenu() {
			child.SubMenu.setDisplayed(displayed)
		}
	}
}

func (menu *Menu) DispatchResponse(reply CommandResponse, conn net.Conn) {
	fmt.Println("Menu(", menu.TargetID, ")::Attempting to dispatch response")
	switch reply.Action {
	case "event":
		menu.HandleEvent(reply, conn)
	case "reply":
		menu.HandleReply(reply, conn)
	}

	for _, child := range menu.Items {
		if child.IsSubMenu() {
			child.SubMenu.DispatchResponse(reply, conn)
		}
	}
}

func (menu *Menu) Send(command *Command, conn net.Conn) {
	ActionId += 1

	command.ID = ActionId

	fmt.Println(command)
	cmd, _ := json.Marshal(command)
	fmt.Println("Writing", string(cmd), "\n", SOCKET_BOUNDARY)

	conn.Write(cmd)
	conn.Write([]byte("\n"))
	conn.Write([]byte(SOCKET_BOUNDARY))
}

func (menu *Menu) Call(command *Command, conn net.Conn) {
	command.Action = "call"
	command.TargetID = menu.TargetID

	if menu.Ready == false {
		menu.CommandQueue = append(menu.CommandQueue, command)
		return
	}
	menu.Send(command, conn)
}

func (menu *Menu) CallWhenReady(command *Command, conn net.Conn) {
	menu.WaitingResponses = append(menu.WaitingResponses, command)
	go func() {
		for {
			if menu.Ready {
				menu.Call(command, conn)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
}

func (menu *Menu) CallWhenChildStable(command *Command, child *Menu, conn net.Conn) {
	menu.WaitingResponses = append(menu.WaitingResponses, command)
	go func() {
		for {
			if child.IsStable() {
				command.Args.MenuID = child.TargetID
				menu.Call(command, conn)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
}

func (menu *Menu) CallWhenTreeStable(command *Command, conn net.Conn) {
	go func() {
		for {
			if menu.IsTreeStable() {
				command.Args.MenuID = menu.TargetID
				menu.WaitingResponses = append(menu.WaitingResponses, command)
				menu.Call(command, conn)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
}

func (menu *Menu) CallWhenDisplayed(command *Command, conn net.Conn) {
	go func() {
		for {
			if menu.Displayed {
				menu.WaitingResponses = append(menu.WaitingResponses, command)
				menu.Call(command, conn)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
}

func (menu *Menu) AddItem(commandID int, label string, conn net.Conn) {
	command := Command{
		Method: "add_item",
		Args: CommandArguments{
			CommandID: commandID,
			Label:     label,
		},
	}
	menuItem := MenuItem{
		CommandID: commandID,
		Label:     label,
		Parent:    menu,
		Type:      "item",
	}
	menu.Items = append(menu.Items, &menuItem)

	menu.CallWhenReady(&command, conn)
}

func (menu *Menu) AddCheckItem(commandID int, label string, conn net.Conn) {
	command := Command{
		Method: "add_check_item",
		Args: CommandArguments{
			CommandID: commandID,
			Label:     label,
		},
	}
	menuItem := MenuItem{
		CommandID: commandID,
		Label:     label,
		Type:      "check",
		Parent:    menu,
	}
	menu.Items = append(menu.Items, &menuItem)
	menu.CallWhenReady(&command, conn)
}

func (menu *Menu) AddRadioItem(commandID int, label string, groupID int, conn net.Conn) {
	command := Command{
		Method: "add_radio_item",
		Args: CommandArguments{
			CommandID: commandID,
			Label:     label,
			GroupID:   groupID,
		},
	}
	menuItem := MenuItem{
		CommandID: commandID,
		Label:     label,
		GroupID:   groupID,
		Parent:    menu,
		Type:      "radio",
	}
	menu.Items = append(menu.Items, &menuItem)
	menu.CallWhenReady(&command, conn)
}

func (menu *Menu) AddSubmenu(commandID int, label string, child *Menu, conn net.Conn) {
	command := Command{
		Method: "add_submenu",
		Args: CommandArguments{
			CommandID: commandID,
			Label:     label,
		},
	}

	// Assign Bidirectional navigation elements i.e. DoublyLinkedLists
	child.Parent = menu
	menuItem := MenuItem{
		CommandID: commandID,
		Label:     label,
		SubMenu:   child,
		Parent:    menu,
	}
	menu.Items = append(menu.Items, &menuItem)

	menu.CallWhenChildStable(&command, child, conn)
}

func (menu *Menu) SetChecked(commandID int, checked bool, asEvent bool, conn net.Conn) {
	command := Command{
		Method: "set_checked",
		Args: CommandArguments{
			CommandID: commandID,
			Checked:   checked,
		},
	}

	for _, item := range menu.Items {
		if item.IsCommandId(commandID) {
			item.Checked = checked
		}
	}
	//if asEvent == false {
	menu.CallWhenDisplayed(&command, conn)
	//}
}

func (menu *Menu) AddSeparator(conn net.Conn) {
	command := Command{
		Method: "add_separator",
	}
	menuItem := MenuItem{
		Type:   "separator",
		Parent: menu,
	}
	menu.Items = append(menu.Items, &menuItem)
	menu.CallWhenReady(&command, conn)
}

func (menu *Menu) SetApplicationMenu(conn net.Conn) {
	command := Command{
		Method: "set_application_menu",
		Args: CommandArguments{
			MenuID: menu.TargetID,
		},
	}

	// Thread to wait for Stable Menu State
	menu.CallWhenTreeStable(&command, conn)
}

/*
A menu is stable if and only if, it is Ready (meaning it was created successfully)
and it has no Commands awaiting Responses.
*/
func (menu *Menu) IsStable() bool {
	return menu.Ready && len(menu.WaitingResponses) == 0
}

/*
A Menu Tree is considered stable if and only if its children nodes report that they are stable.
Function is recursive, so factor that in to performance
*/
func (menu *Menu) IsTreeStable() bool {
	if !menu.IsStable() {
		return false
	}

	for _, child := range menu.Items {
		//fmt.Println("Checking child")
		if child.IsSubMenu() {
			if !child.SubMenu.IsTreeStable() {
				return false
			}
		}
	}

	return true
}
