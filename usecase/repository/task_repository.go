package repository

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type TaskRepository interface {
	FindByID(id uint) (*model.Task, error)
}
