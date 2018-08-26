package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yoshinorihisakawa/api-taskmanagement/infrastructure/api/handler"
	"github.com/yoshinorihisakawa/api-taskmanagement/infrastructure/datastore"
	"github.com/yoshinorihisakawa/api-taskmanagement/interface/controllers"
	"github.com/yoshinorihisakawa/api-taskmanagement/interface/presenters"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/presenter"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/repository"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/service"
)

type interactor struct {
	conn *gorm.DB
}

type Iteractor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *gorm.DB) Iteractor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewUserHandler()
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controllers.UserController {
	return controllers.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.conn)
}

func (i *interactor) NewUserPresenter() presenter.UserPresenter {
	return presenters.NewUserPresenter()
}
