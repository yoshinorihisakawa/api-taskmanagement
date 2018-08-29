package router

import (
	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/api-taskmanagement/infrastructure/api/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {

	// user関連
	e.POST("/users", handler.CreateUser)
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)

	// taskUserSet関連
	//e.POST("/task_user_set", handler.CreateTask)
	e.GET("/task_user_sets/:id", handler.GetTaskUserSet)
	e.POST("/task_user_sets", handler.CreateTaskUserSet)
}
