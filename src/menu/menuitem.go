package menu

import (
	"fmt"
	"net"
)

type Item interface {
	IsSubMenu() bool
	IsCommandId() bool
	Handle()
}

type MenuItem struct {
	CommandID int    `json:"command_id,omitempty"`
	Label     string `json:"label,omitempty"`
	GroupID   int    `json:"group_id,omitempty"`
	SubMenu   *Menu  `json:"submenu,omitempty"`
	Type      string `json:"type,omitempty"`
	Checked   bool   `json:"checked,omitempty"`
	Parent    *Menu  `json:"-"`
}

func (mi MenuItem) IsSubMenu() bool {
	return mi.SubMenu != nil
}

func (mi MenuItem) IsCommandId(commandID int) bool {
	return mi.CommandID == commandID
}

func (mi MenuItem) HandleEvent(conn net.Conn) {
	fmt.Println("EventType", mi.Type)
	switch mi.Type {
	case "check":
		fmt.Println("Toggling Checked(", mi.Checked, ")", "to", "checked(", !mi.Checked, ")")
		mi.Parent.SetChecked(mi.CommandID, !mi.Checked, false, conn)
	}
}
