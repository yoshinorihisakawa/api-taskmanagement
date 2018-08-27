package model

import (
	"time"
)

type Task struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	TaskName  string     `json:"task_name" validate:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
