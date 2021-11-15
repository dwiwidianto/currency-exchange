package controller

import (
	"net/http"
	config "widi443/currency-exchange/configs"
	model "widi443/currency-exchange/models"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {

	user := model.User{}
	c.Bind(&user)

	err := config.DBConnection.Save(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create user",
		"user":    user,
	})
}

func GetUserController(c echo.Context) error {
	var users []model.User

	err := config.DBConnection.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success",
		"data":    users,
	})
}
