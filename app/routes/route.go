package routes

import (
	transactionController "currency-exchange/controllers/transactions"
	userController "currency-exchange/controllers/users"

	echo "github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController        userController.UserController
	TransactionController transactionController.TransactionController
}

func (controller RouteControllerList) RouteRegiester(c *echo.Echo) {

	users := c.Group("user")
	users.GET("", controller.UserController.GetAllUsersController)
	users.POST("/create", controller.UserController.CreateUsersController)
	users.POST("/login", controller.UserController.LoginController)
	users.DELETE("/:userId", controller.UserController.DeleteUserController)

	transaction := c.Group("transaction")
	transaction.GET("", controller.TransactionController.GetAllTransactionController)
	transaction.POST("/create", controller.TransactionController.CreateTransactionController)
	transaction.DELETE("/:transactionId", controller.TransactionController.CreateTransactionController)

}
