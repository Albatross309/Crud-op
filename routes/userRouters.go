package routes

import (
	userServices "crud-app/services"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.POST("/users", userServices.CreateUser)
	e.GET("/users", userServices.GetUsers)
	e.GET("/users/:id", userServices.GetUser)
	e.PUT("/users/:id", userServices.UpdateUser)
	e.DELETE("/users/:id", userServices.DeleteUser)
}
