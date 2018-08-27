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

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

type handle struct {
	handler.UserHandler
	handler.TaskHandler
	handler.TaskUserSetHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return &handle{
		UserHandler:        i.NewUserHandler(),
		TaskHandler:        i.NewTaskHandler(),
		TaskUserSetHandler: i.NewTaskUserSetHandler(),
	}
}

// user関連
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

// task関連
func (i *interactor) NewTaskHandler() handler.TaskHandler {
	return handler.NewTaskHandler(i.NewTaskController())
}

func (i *interactor) NewTaskController() controllers.TaskController {
	return controllers.NewTaskController(i.NewTaskService())
}

func (i *interactor) NewTaskService() service.TaskService {
	return service.NewTaskService(
		i.NewTaskRepository(),
		i.NewTaskPresenter(),
	)
}

func (i *interactor) NewTaskRepository() repository.TaskRepository {
	return datastore.NewTaskRepository(i.conn)
}

func (i *interactor) NewTaskPresenter() presenter.TaskPresenter {
	return presenters.NewTaskPresenter()
}

// taskUserSet関連
func (i *interactor) NewTaskUserSetHandler() handler.TaskUserSetHandler {
	return handler.NewTaskUserSetHandler(i.NewTaskUserSetController())
}

func (i *interactor) NewTaskUserSetController() controllers.TaskUserSetController {
	return controllers.NewTaskUserSetController(i.NewTaskUserSetService())
}

func (i *interactor) NewTaskUserSetService() service.TaskUserSetService {
	return service.NewTaskUsetSetService(
		i.NewTaskRepository(),
		i.NewTaskPresenter(),
		i.NewUserRepository(),
		i.NewTaskUserSetRepository(),
		i.NewTaskUserSetPresenter(),
	)
}

func (i *interactor) NewTaskUserSetRepository() repository.TaskUserSetRepository {
	return datastore.NewTaskUserSetRepository(i.conn)
}

func (i *interactor) NewTaskUserSetPresenter() presenter.TaskUserSetPresenter {
	return presenters.NewTaskUserSetPresenter()
}
