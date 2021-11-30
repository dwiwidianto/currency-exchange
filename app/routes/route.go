package routes

import (
	"currency-exchange/app/middlewares"
	userController "currency-exchange/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController userController.UserController
}

func (controller *RouteControllerList) RouteRegiester(e *echo.Echo) {

	api := e.Group("v1/")
	api.POST("login", controller.UserController.LoginController)

	api.GET("users", controller.UserController.GetAllUsersController, middleware.JWTWithConfig(controller.JWTMiddleware), middlewares.IsAdmin)
	api.POST("users", controller.UserController.CreateUsersController)
	api.DELETE("users/:userId", controller.UserController.DeleteUserController, middleware.JWTWithConfig(controller.JWTMiddleware), middlewares.IsAdmin)
}
