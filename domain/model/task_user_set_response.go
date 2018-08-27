package model

type TaskUserSetResponse struct {
	TaskID      uint
	TaskName    string
	SendUser    User    `json:"send_user"`
	ReceiveUser []*User `json:"receive_user"`
}
