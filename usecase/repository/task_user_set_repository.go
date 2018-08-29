package repository

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type TaskUserSetRepository interface {
	FindByID(id uint) (*model.TaskUserSet, error)
	FindAllIDsByID(id uint) ([]*model.TaskUserSet, error)
	Store(tus *model.TaskUserSet) error
}
