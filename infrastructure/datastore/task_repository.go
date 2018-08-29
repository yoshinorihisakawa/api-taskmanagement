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
	Store(task *model.Task) (*model.Task, error)
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

func (taskRepository *taskRepository) Store(task *model.Task) (*model.Task, error) {
	// TODO:トランザクション開始

	err := taskRepository.db.Omit("updated_at").Save(task).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	t := &model.Task{}
	if err := taskRepository.db.Last(&t).Error; err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return t, nil
}
