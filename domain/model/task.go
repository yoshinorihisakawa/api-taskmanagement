package model

import (
	"time"

	"github.com/yoshinorihisakawa/sample-di/model"
)

type Task struct {
	ID        uint          `gorm:"primary_key" json:"id"`
	TaskName  string        `json:"task_name" validate:"required"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt *time.Time    `json:"-"`
	DeletedAt *time.Time    `json:"-"`
	Users     []*model.User `json:"users" validate:"required"`
}
