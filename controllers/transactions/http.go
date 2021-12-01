package controllers

import (
	"currency-exchange/business/transactions"
	"currency-exchange/controllers"
	"currency-exchange/controllers/transactions/request"
	"currency-exchange/controllers/transactions/response"
	"currency-exchange/helpers"

	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	usecase transactions.TransactionUseCaseInterfaces
}

func NewTransactionController(uc transactions.TransactionUseCaseInterfaces) *TransactionController {
	return &TransactionController{
		usecase: uc,
	}
}

func (controller *TransactionController) CreateTransactionController(c echo.Context) error {
	var userInput request.CreateTransaction
	err := c.Bind(&userInput)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	inputs, _ := controller.usecase.CreateTransaction(*userInput.ToDomain())
	return controllers.NewSuccessResponse(c, response.FromDomain(inputs))
}

// func (controller *TransactionController) CreateTransactionController(c echo.Context) error {
// 	request := request.CreateTransaction{}
// 	var err error
// 	err = c.Bind(&request)
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	// ctxNative := c.Request().Context()
// 	var data transactions.Domain
// 	data, err = controller.usecase.CreateTransaction(*request.ToDomain())
// 	if err != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controllers.NewSuccessResponse(c, response.FromDomain(data))
// }

func (controller *TransactionController) GetAllTransactionController(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.usecase.GetAllTransaction(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromListDomain(data))
}

func (controller *TransactionController) DeleteTransactionController(c echo.Context) error {
	getId := c.Param("transactionId")
	id, err := helpers.StringToInt(getId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, id)
}
