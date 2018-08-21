package router

import (
	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/sample-api-hoop/infrastructure/api/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {

	// user関連
	e.POST("/users", handler.CreateUser)
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)

	// task関連
}
