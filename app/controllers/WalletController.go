package controller

import (
	"net/http"
	config "widi443/currency-exchange/app/configs"
	model "widi443/currency-exchange/app/models"

	"github.com/labstack/echo/v4"
)

func CreateWallet(c echo.Context) error {
	wallet := model.Wallet{}
	c.Bind(&wallet)

	err := config.DBConnection.Save(&wallet).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success create wallet",
		"data":    wallet,
	})
}

func GetWallet(c echo.Context) error {
	var wallet []model.Wallet

	err := config.DBConnection.Preload("User").Find(&wallet).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get data Wallet",
		"data":    wallet,
	})
}
