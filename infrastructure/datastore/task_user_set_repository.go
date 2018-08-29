package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
)

type taskUserSetRepository struct {
	db *gorm.DB
}

type TaskUserSetRepository interface {
	FindByID(id uint) (*model.TaskUserSet, error)
	FindAllIDsByID(id uint) ([]*model.TaskUserSet, error)
	Store(tus *model.TaskUserSet) error
}

func NewTaskUserSetRepository(db *gorm.DB) TaskUserSetRepository {
	return &taskUserSetRepository{db}
}

func (taskUserSetRepository *taskUserSetRepository) FindByID(id uint) (*model.TaskUserSet, error) {
	taskUserSet := &model.TaskUserSet{}
	err := taskUserSetRepository.db.Where("task_id = ?", id).Find(&taskUserSet).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return taskUserSet, nil
}

func (taskUserSetRepository *taskUserSetRepository) FindAllIDsByID(id uint) ([]*model.TaskUserSet, error) {
	tu := []*model.TaskUserSet{}
	err := taskUserSetRepository.db.Select("receive_user_id").Where("task_id = ?", id).Find(&tu).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return tu, nil
}

func (taskUserSetRepository *taskUserSetRepository) Store(tus *model.TaskUserSet) error {
	return taskUserSetRepository.db.Omit("updated_at").Save(tus).Error
}
