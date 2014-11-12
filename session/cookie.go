package session

import "time"

/*
cookie
source the source url
name the cookie name
value the cookie value
domain the cookie domain
path the cookie path
creation the creation date
expiry the expiration date
last_access the last time the cookie was accessed
secure is the cookie secure
http_only is the cookie only valid for HTTP
priority internal priority information
*/

type Cookie struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	// Need to check what type value is,
	// Im on train so have no wifi
	Domain     string    `json:"domain"`
	Path       string    `json:"path"`
	Creation   time.Date `json:"creation"`
	Expiry     time.Date `json:"expiry"`
	LastAccess time.Date `json:"last_access"`
	Secure     bool      `json:"secure"`
	HttpOnly   bool      `json:"http_only"`
	Priority   uint8     `json:"priority"`
	// need to check what type priority is
}
