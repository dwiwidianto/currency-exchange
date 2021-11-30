package controllers

import (
	"currency-exchange/business/users"
	"currency-exchange/controllers"
	"currency-exchange/controllers/users/request"
	"currency-exchange/controllers/users/response"
	"currency-exchange/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.UserUseCaseInterface
}

func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (controller *UserController) GetAllUsersController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := controller.userUsecase.GetAll(ctxNative)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromUserListDomain(data))
}

func (controller *UserController) CreateUsersController(c echo.Context) error {
	request := request.UserRegister{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctxNative := c.Request().Context()
	var data users.Domain
	data, err = controller.userUsecase.Create(ctxNative, request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainToCreateResponse(data))
}

func (controller *UserController) LoginController(c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	request := request.UserLogin{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, token, err = controller.userUsecase.Login(ctx, request.Password, request.Email)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromUserToDomainLogin(login, token))
}

func (controller *UserController) DeleteUserController(c echo.Context) error {
	id := c.Param("userId")
	convId, err := helpers.StringToInt(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.userUsecase.Delete(ctx, convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, convId)
}
