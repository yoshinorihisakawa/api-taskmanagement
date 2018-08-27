package presenters

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type taskPresenter struct {
}

func NewTaskPresenter() TaskPresenter {
	return &taskPresenter{}
}

type TaskPresenter interface {
	ResponseTasks(t []*model.Task) []*model.Task
}

func (taskPresenter *taskPresenter) ResponseTasks(t []*model.Task) []*model.Task {

	return t
}
