package service

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/presenter"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/repository"
)

type taskService struct {
	TaskRepository repository.TaskRepository
	TaskPresenter  presenter.TaskPresenter
}

type TaskService interface {
	Create(u *model.Task) error
}

func NewTaskService(tr repository.TaskRepository, pre presenter.TaskPresenter) TaskService {
	return &taskService{tr, pre}
}

func (taskService *taskService) Create(task *model.Task) error {
	return nil
}
