package main

import "net"

type SizeHW struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type CommandArguments struct {
	RootUrl string `json:"root_url,omitempty"`
	Title   string `json:"title,omitempty"`
	Size    SizeHW `json:"size,omitempty"`
	X       int    `json:"x,omitempty"`
	Y       int    `json:"y,omitempty"`
}
type Command struct {
	ID         int              `json:"_id"`
	Action     string           `json:"_action"`
	ObjectType string           `json:"_type,omitempty"`
	Method     string           `json:"_method,omitempty"`
	TargetID   int              `json:"_target,omitempty"`
	Args       CommandArguments `json:"_args"`
}

func (c Command) Send(conn net.Conn) {

}

//{"_action":"reply","_error":"","_id":1,"_result":{"_target":3}}

type ResponseResult struct {
	TargetID int `json:"_target,omitempty"`
}

type CommandResponse struct {
	Action string         `json:"_action,omitempty"`
	Error  string         `json:"_error,omitempty"`
	ID     int            `json:"_id,omitempty"`
	Result ResponseResult `json:"_result,omitempty"`
}
