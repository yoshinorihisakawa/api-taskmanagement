package presenters

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type taskUserSetPresenter struct {
}

func NewTaskUserSetPresenter() TaskUserSetPresenter {
	return &taskUserSetPresenter{}
}

type TaskUserSetPresenter interface {
	ResponseTaskUserSet(t *model.Task, su *model.User, ru []*model.User) (*model.TaskUserSetResponse, error)
}

func (taskUserSetPresenter *taskUserSetPresenter) ResponseTaskUserSet(t *model.Task, su *model.User, ru []*model.User) (*model.TaskUserSetResponse, error) {

	tusr := &model.TaskUserSetResponse{
		TaskID:      t.ID,
		TaskName:    t.TaskName,
		SendUser:    *su,
		ReceiveUser: ru,
	}

	return tusr, nil
}
