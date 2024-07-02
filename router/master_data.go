package router

import (
	"example/helper"
	"example/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handleMasterData struct {
	MasterDataService service.MasterDataService
}

func NewMasterDataRouter(e *echo.Echo) {
	h := &handleMasterData{
		service.NewMasterDataService(),
	}

	router := e.Group("api/v1/master_data")
	router.GET("/skpd", h.ListSkpd)
	router.GET("/urusan", h.ListUrusan)
}

func (h handleMasterData) ListSkpd(c echo.Context) error {
	list, err := h.MasterDataService.ListSkpd()
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, list)
}

func (h handleMasterData) ListUrusan(c echo.Context) error {
	list, err := h.MasterDataService.ListUrusan()
	if err != nil {
		return helper.GenerateErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}

	return helper.GenerateSuccessResponse(c, list)
}
