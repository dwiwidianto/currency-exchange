package routes

import (
	userController "currency-exchange/controllers/users"

	echo "github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController
}

func (controller RouteControllerList) RouteRegiester(c *echo.Echo) {

	users := c.Group("/user")
	users.GET("", controller.UserController.GetAllUsersController)
	users.POST("/create", controller.UserController.CreateUsersController)
	users.POST("/login", controller.UserController.LoginController)
}
