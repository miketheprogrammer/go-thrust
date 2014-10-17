package commands

import (
	"net"
)

type SizeHW struct {
	Width  uint `json:"width,omitempty"`
	Height uint `json:"height,omitempty"`
}

/*
Covers all possible argument combinations.
Makes use of omit empty to adapt to different use cases
*/
type CommandArguments struct {
	RootUrl   string `json:"root_url,omitempty"`
	Title     string `json:"title,omitempty"`
	Size      SizeHW `json:"size,omitempty"`
	X         int    `json:"x,omitempty"`
	Y         int    `json:"y,omitempty"`
	CommandID uint   `json:"command_id,omitempty"`
	Label     string `json:"label,omitempty"`
	MenuID    uint   `json:"menu_id,omitempty"` // this should never be 0 anyway
	WindowID  uint   `json:"window_id,omitempty"`
	GroupID   uint   `json:"group_id,omitempty"`
	Value     bool   `json:"value"`
}
type Command struct {
	ID         uint             `json:"_id"`
	Action     string           `json:"_action"`
	ObjectType string           `json:"_type,omitempty"`
	Method     string           `json:"_method"`
	TargetID   uint             `json:"_target,omitempty"`
	Args       CommandArguments `json:"_args"`
}

func (c Command) Send(conn net.Conn) {

}

//{"_action":"reply","_error":"","_id":1,"_result":{"_target":3}}

type ReplyResult struct {
	TargetID uint `json:"_target,omitempty"`
}

type EventResult struct {
	CommandID  uint `json:"command_id,omitempty"`
	EventFlags int  `json:"event_flags,omitempty"`
}

type CommandResponse struct {
	Action string      `json:"_action,omitempty"`
	Error  string      `json:"_error,omitempty"`
	ID     uint        `json:"_id,omitempty"`
	Result ReplyResult `json:"_result,omitempty"`
	Event  EventResult `json:"_event,omitempty"`
}
