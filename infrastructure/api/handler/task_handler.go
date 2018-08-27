package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/interface/controllers"
)

type taskHandler struct {
	taskController controllers.TaskController
}

type TaskHandler interface {
	CreateTask(c echo.Context) error
}

func NewTaskHandler(tc controllers.TaskController) TaskHandler {
	return &taskHandler{taskController: tc}
}

func (uh *taskHandler) CreateTask(c echo.Context) error {
	// リクエスト用のEntityを作成
	req := &model.Task{}

	// bind
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	// validate
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uh.taskController.CreateTask(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
