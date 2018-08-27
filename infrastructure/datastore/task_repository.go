package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
)

type taskRepository struct {
	db *gorm.DB
}

type TaskRepository interface {
	FindByID(id uint) (*model.Task, error)
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (taskRepository *taskRepository) FindByID(id uint) (*model.Task, error) {

	task := &model.Task{}
	err := taskRepository.db.Find(&task).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return task, nil
}
