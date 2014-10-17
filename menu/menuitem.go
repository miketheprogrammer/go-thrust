package menu

import . "github.com/miketheprogrammer/thrust-go/common"

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
	Checked   bool   `json:"checked"`
	Enabled   bool   `json:"enabled"`
	Visible   bool   `json:"visible"`
	Parent    *Menu  `json:"-"`
}

func (mi MenuItem) IsSubMenu() bool {
	return mi.SubMenu != nil
}

func (mi MenuItem) IsCheckItem() bool {
	return mi.Type == "check"
}

func (mi MenuItem) IsRadioItem() bool {
	return mi.Type == "radio"
}

func (mi MenuItem) IsGroupID(groupID int) bool {
	return mi.GroupID == groupID
}

func (mi MenuItem) IsCommandID(commandID int) bool {
	return mi.CommandID == commandID
}

func (mi MenuItem) HandleEvent() {
	Log.Debug("EventType", mi.Type)
	switch mi.Type {
	case "check":
		Log.Debug("Toggling Checked(", mi.Checked, ")", "to", "checked(", !mi.Checked, ")")
		mi.Parent.SetChecked(mi.CommandID, !mi.Checked)
	case "radio":
		Log.Debug("Toggling RadioChecked(", mi.Checked, ")", "to", "checked(", !mi.Checked, ")")
		mi.Parent.ToggleRadio(mi.CommandID, mi.GroupID, !mi.Checked)
	}

}
