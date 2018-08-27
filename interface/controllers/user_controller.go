package controllers

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	CreateUser(user *model.User) error
	GetUsers() ([]*model.User, error)
	GetUser(id uint) (*model.User, error)
	UpdateUser(user *model.User, id uint) (*model.User, error)
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (userController *userController) CreateUser(user *model.User) error {
	// 内側の処理のための技術的な関心事を記載
	return userController.userService.CreateUser(user)
}

func (userController *userController) GetUsers() ([]*model.User, error) {
	user := []*model.User{}
	u, err := userController.userService.GetUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userController *userController) GetUser(id uint) (*model.User, error) {

	u, err := userController.userService.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userController *userController) UpdateUser(user *model.User, id uint) (*model.User, error) {
	user.ID = id
	u, err := userController.userService.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
