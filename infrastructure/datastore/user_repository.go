package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Store(user *model.User) error
	FindAll(users []*model.User) ([]*model.User, error)
	FindByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	FindAllByID(id []*int64) ([]*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Store(user *model.User) error {
	return userRepository.db.Omit("updated_at").Save(user).Error

}

func (userRepository *userRepository) FindAll(users []*model.User) ([]*model.User, error) {

	err := userRepository.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return users, nil
}

func (userRepository *userRepository) FindByID(id uint) (*model.User, error) {
	user := &model.User{}
	err := userRepository.db.First(&user, id).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return user, nil
}

func (userRepository *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := userRepository.db.Save(&user).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return user, nil
}

func (userRepository *userRepository) FindAllByID(id []*int64) ([]*model.User, error) {
	users := []*model.User{}
	fmt.Println(&id[0])
	err := userRepository.db.Where("id in (?)", id).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return users, nil
}
