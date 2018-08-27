package service

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/presenter"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	CreateUser(u *model.User) error
	GetUser(u []*model.User) ([]*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(u *model.User) (*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

func (userService *userService) CreateUser(user *model.User) error {
	return userService.UserRepository.Store(user)
}

func (userService *userService) GetUser(users []*model.User) ([]*model.User, error) {
	u, err := userService.UserRepository.FindAll(users)
	if err != nil {
		return nil, err
	}
	return userService.UserPresenter.ResponseUsers(u), nil
}

func (userService *userService) GetUserByID(id uint) (*model.User, error) {
	u, err := userService.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userService *userService) UpdateUser(user *model.User) (*model.User, error) {
	u, err := userService.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
