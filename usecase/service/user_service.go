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
	Create(u *model.User) error
	Get(u []*model.User) ([]*model.User, error)
	GetById(id uint) (*model.User, error)
	Update(u *model.User) (*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

func (userService *userService) Create(user *model.User) error {
	return userService.UserRepository.Store(user)
}

func (userService *userService) Get(users []*model.User) ([]*model.User, error) {
	u, err := userService.UserRepository.FindAll(users)
	if err != nil {
		return nil, err
	}
	return userService.UserPresenter.ResponseUsers(u), nil
}

func (userService *userService) GetById(id uint) (*model.User, error) {
	u, err := userService.UserRepository.FetchById(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userService *userService) Update(user *model.User) (*model.User, error) {
	u, err := userService.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
