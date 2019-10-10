package defs

type MemberInfo struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Realname  string `json:"realname"`
	Sex       string `json:"sex"`
	Signature string `json:"signature"`
}
