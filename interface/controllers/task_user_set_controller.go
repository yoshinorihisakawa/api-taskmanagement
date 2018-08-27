package controllers

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/service"
)

type taskUserSetController struct {
	taskUserSetService service.TaskUserSetService
}

type TaskUserSetController interface {
	GetTaskUserSet(id uint) (*model.TaskUserSetResponse, error)
}

func NewTaskUserSetController(ts service.TaskUserSetService) TaskUserSetController {
	return &taskUserSetController{ts}
}

func (taskUserSetController *taskUserSetController) GetTaskUserSet(id uint) (*model.TaskUserSetResponse, error) {

	u, err := taskUserSetController.taskUserSetService.GetTaskUserSet(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
