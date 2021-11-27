package controllers

import (
	"currency-exchange/business/users"
	"currency-exchange/controllers"
	"currency-exchange/controllers/users/request"
	"currency-exchange/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.UserUseCaseInterface
}

func NewUserController(c *echo.Echo, uc users.UserUseCaseInterface) *UserController {
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

func (ctrl *UserController) LoginController(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)
	user, err := ctrl.userUsecase.Login(*userLogin.ToDomain(), ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromUsersDomain(user))
}
