package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/interface/controllers"
)

type taskUserSetHandler struct {
	taskUserSetController controllers.TaskUserSetController
}

type TaskUserSetHandler interface {
	GetTaskUserSet(c echo.Context) error
	CreateTaskUserSet(c echo.Context) error
}

func NewTaskUserSetHandler(tc controllers.TaskUserSetController) TaskUserSetHandler {
	return &taskUserSetHandler{taskUserSetController: tc}
}

func (uh *taskUserSetHandler) GetTaskUserSet(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: "invalid id"})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.taskUserSetController.GetTaskUserSet(uint(id))
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, u)
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *taskUserSetHandler) CreateTaskUserSet(c echo.Context) error {
	// リクエスト用のEntityを作成
	req := &model.TaskUserSetRequest{}

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

	err := uh.taskUserSetController.CreateTaskUserSet(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
