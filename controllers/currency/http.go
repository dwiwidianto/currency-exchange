package currency

import (
	"currency-exchange/business/currency"
	"currency-exchange/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CurrencyApiController struct {
	currencyApiUC currency.Usecase
}

func NewCurrencyApiController(e *echo.Echo, cu currency.Usecase) *CurrencyApiController {
	return &CurrencyApiController{
		currencyApiUC: cu,
	}
}

func (ctrl *CurrencyApiController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := ctrl.currencyApiUC.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, resp)
}
