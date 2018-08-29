package model

import "time"

type TaskUserSet struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	TaskID        uint       `json:"task_id"`
	Task          Task       `json:"-"`
	SendUserID    uint       `json:"send_user_id"`
	ReceiveUserID uint       `json:"receive_user_id"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}

type TaskUserSetRequest struct {
	TaskName    string    `json:"task_name" validate:"required"`
	SendUser    UserID    `json:"send_user" validate:"required"`
	ReceiveUser []*UserID `json:"receive_user" validate:"required"`
}
