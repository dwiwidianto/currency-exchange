package routes

import (
	userController "currency-exchange/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController
}

func (controller RouteControllerList) RouteRegiester(c *echo.Echo) {
	users := c.Group("/user")
	users.GET("", controller.UserController.GetAllUsers)
	users.POST("/login", controller.UserController.Login)
}
