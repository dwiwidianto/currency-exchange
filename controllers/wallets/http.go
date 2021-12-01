package controllers

import (
	"currency-exchange/business/wallets"
	"currency-exchange/controllers"
	"currency-exchange/controllers/wallets/response"
	"currency-exchange/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletController struct {
	walletUsecase wallets.WalletsUseCaseInterfaces
}

func NewWalletController(c *echo.Echo, uc wallets.WalletsUseCaseInterfaces) *WalletController {
	return &WalletController{
		walletUsecase: uc,
	}
}

func (controller *WalletController) GetAllWalletController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := controller.walletUsecase.GetAllWallets(ctxNative)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromWalletListDomain(data))
}

func (controller *WalletController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToInt(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.walletUsecase.DeleteWallets(ctx, convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, convId)
}
