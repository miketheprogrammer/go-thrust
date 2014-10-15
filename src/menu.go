package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Menu struct {
	TargetID         int
	CommandHistory   []*Command
	ResponseHistory  []*CommandResponse
	WaitingResponses []*Command
	CommandQueue     []*Command
	Conn             net.Conn
	Ready            bool
	Displayed        bool
	Parent           *Menu
}

func (menu *Menu) Create(conn net.Conn) {
	menuCreate := Command{
		Action:     "create",
		ObjectType: "menu",
	}
	menu.Send(&menuCreate, conn)
}

func (menu *Menu) IsTarget(targetId int) bool {
	return targetId == menu.TargetID
}
func (menu *Menu) HandleError(reply CommandResponse) {

}

func (menu *Menu) HandleEvent(reply CommandResponse) {

}

func (menu *Menu) HandleReply(reply CommandResponse, conn net.Conn) {
	fmt.Println("MENU::Handling Response", reply)
	for k, v := range menu.WaitingResponses {
		if v.ID != reply.ID {
			continue
		}
		if len(menu.WaitingResponses) > 1 {
			menu.WaitingResponses = menu.WaitingResponses[:k+copy(menu.WaitingResponses[k:], menu.WaitingResponses[k+1:])]
		} else {
			menu.WaitingResponses = []*Command{}
		}

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
			menu.Displayed = true
		}

	}
}

func (menu *Menu) Send(command *Command, conn net.Conn) {
	ActionId += 1

	command.ID = ActionId

	fmt.Println(command)
	cmd, _ := json.Marshal(&command)
	fmt.Println("Writing", string(cmd), "\n", SOCKET_BOUNDARY)

	menu.WaitingResponses = append(menu.WaitingResponses, command)

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

func (menu *Menu) InsertItemAt(index, commandID int, label string, conn net.Conn) {
	command := Command{
		Method: "insert_item_at",
		Args: CommandArguments{
			CommandID: commandID,
			Index:     index,
			Label:     label,
		},
	}

	menu.Call(&command, conn)
}

func (menu *Menu) InsertSubmenuAt(index, commandID int, label string, child *Menu, conn net.Conn) {
	command := Command{
		Method: "insert_submenu_at",
		Args: CommandArguments{
			CommandID: commandID,
			Index:     index,
			Label:     label,
			MenuID:    child.TargetID,
		},
	}
	fmt.Println(command)
	child.Parent = menu

	menu.Call(&command, conn)
}

func (menu *Menu) SetApplicationMenu(conn net.Conn) {
	command := Command{
		Method: "set_application_menu",
	}

	menu.Call(&command, conn)
}
