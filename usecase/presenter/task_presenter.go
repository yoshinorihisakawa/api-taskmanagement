package presenter

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type TaskPresenter interface {
	ResponseTasks(us []*model.Task) []*model.Task
}
