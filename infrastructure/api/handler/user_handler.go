package handler

import (
	"context"

	"net/http"

	"strconv"

	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/sample-api-hoop/domain/model"
	"github.com/yoshinorihisakawa/sample-api-hoop/interface/controllers"
)

type userHandler struct {
	userController controllers.UserController
}

type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
}

func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) CreateUser(c echo.Context) error {

	// リクエスト用のEntityを作成
	req := &model.User{}

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

	err := uh.userController.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
func (uh *userHandler) GetUsers(c echo.Context) error {

	req := &model.User{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.userController.GetUsers()
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *userHandler) GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: "invalid id"})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.userController.GetUser(uint(id))
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, u)
	}

	return c.JSON(http.StatusOK, u)
}

func (uh *userHandler) UpdateUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	if id < 1 {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: "invalid id"})
	}

	// リクエスト用のEntityを作成
	req := &model.User{}

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

	u, err := uh.userController.UpdateUser(req, uint(id))
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, u)
	}

	return c.JSON(http.StatusOK, u)
}
