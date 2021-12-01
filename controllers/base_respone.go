package controllers

import (
	"currency-exchange/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = helpers.CustomResponseMassageSuccess
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = helpers.CustomResponseMassageFiled
	response.Meta.Messages = []string{err.Error()}
	return c.JSON(http.StatusOK, response)
}
