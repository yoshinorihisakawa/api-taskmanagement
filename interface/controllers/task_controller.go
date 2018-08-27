package controllers

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/service"
)

type taskController struct {
	taskService service.TaskService
}

type TaskController interface {
	CreateTask(task *model.Task) error
}

func NewTaskController(ts service.TaskService) TaskController {
	return &taskController{ts}
}

func (taskController *taskController) CreateTask(task *model.Task) error {
	// 内側の処理のための技術的な関心事を記載
	return taskController.taskService.Create(task)
}
