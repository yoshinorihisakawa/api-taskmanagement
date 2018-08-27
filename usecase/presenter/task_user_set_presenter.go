package presenter

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type TaskUserSetPresenter interface {
	ResponseTaskUserSet(t *model.Task, su *model.User, ru []*model.User) (*model.TaskUserSetResponse, error)
}
