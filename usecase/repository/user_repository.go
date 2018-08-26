package repository

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type UserRepository interface {
	Store(user *model.User) error
	FindAll(users []*model.User) ([]*model.User, error)
	FetchById(id uint) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
}
