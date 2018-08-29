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
	CreateTaskUserSet(tus *model.TaskUserSetRequest) error
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

func (taskUserSetController *taskUserSetController) CreateTaskUserSet(tus *model.TaskUserSetRequest) error {
	// 内側の処理のための技術的な関心事を記載
	return taskUserSetController.taskUserSetService.CreateTaskUserSet(tus)
}
