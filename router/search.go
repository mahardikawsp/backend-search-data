package router

import (
	"context"
	"example/helper"
	"example/model"
	"example/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type handleSearch struct {
	SearchService service.SearchService
}

func NewSearchRouter(e *echo.Echo) {
	h := &handleSearch{
		service.NewSearchService(),
	}

	router := e.Group("api/v1/search")
	router.GET("", h.List)
}

func (h handleSearch) List(c echo.Context) error {
	ctx := context.Background()
	param := new(model.SearchParam)

	if err := c.Bind(param); err != nil {
		log.Error(err)
		return helper.GenerateErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	list, err := h.SearchService.List(ctx, param)
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, list)
}
