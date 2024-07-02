package router

import (
	"context"
	"example/helper"
	"example/model"
	"example/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handleEconnectService struct {
	EconnectService service.EconnectService
}

func NewEconnectRouter(e *echo.Echo) {
	h := &handleEconnectService{
		service.NewEconnectService(),
	}

	router := e.Group("api/v1/econnect")

	router.POST("/create", h.Create)
	router.GET("", h.List)
	router.GET("/detail/:id", h.Detail)
	router.PUT("/update/:id", h.Update)
}

func (h handleEconnectService) Create(c echo.Context) error {
	ctx := context.Background()
	param := new(model.EconnectParam)

	err := c.Bind(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	err, httpStatus := h.EconnectService.Create(param, ctx)
	if err != nil {
		return helper.GenerateErrorResponse(c, httpStatus, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, "success")
}

func (h handleEconnectService) List(c echo.Context) error {
	param := new(model.FilterOptions)

	err := c.Bind(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	data, err := h.EconnectService.List(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, data)
}

func (h handleEconnectService) Detail(c echo.Context) error {
	param := new(model.EconnectParam)

	err := c.Bind(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	data, err := h.EconnectService.Detail(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, data)
}

func (h handleEconnectService) Update(c echo.Context) error {
	ctx := context.Background()
	param := new(model.EconnectParam)

	err := c.Bind(param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	err, httpStatus := h.EconnectService.Update(param, ctx)
	if err != nil {
		return helper.GenerateErrorResponse(c, httpStatus, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, "success")
}
