package presenter

import "github.com/yoshinorihisakawa/api-taskmanagement/domain/model"

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}
