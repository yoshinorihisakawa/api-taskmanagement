package model

import "time"

type TaskUserSet struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	TaskID        uint       `json:"task_id"`
	SendUserID    uint       `json:"send_user_id"`
	ReceiveUserID uint       `json:"receive_user_id"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}
