package helper

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
)

type (
	Response struct {
		Data  interface{} `json:"data,omitempty"`
		Error interface{} `json:"error,omitempty"`
	}

	ErrorResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Errors  interface{} `json:"errors"`
	}

	PaginateResponse struct {
		Total       int         `json:"total"`
		PerPage     int         `json:"per_page"`
		CurrentPage int         `json:"current_page"`
		TotalPage   int         `json:"total_page"`
		Items       interface{} `json:"items"`
		Offset      int         `json:"offset"`
	}

	PaginateParam struct {
		Page  int
		Limit int
		Count int
	}
)

func GenerateErrorResponse(c echo.Context, code int, message string, err error) error {
	fmt.Println("error response : ", message, err)
	return c.JSON(code, Response{Error: ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  err,
	}})
}

func GenerateSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Data: data,
	})
}

func Paginate(p *PaginateParam, result interface{}) *PaginateResponse {
	if p.Page < 1 {
		fmt.Println("true kurang 1")
		p.Page = 1
	}

	var paginator PaginateResponse
	var offset int

	offset, p.Limit = GetLimitOffset(p.Page, p.Limit)
	if p.Count == 0 {
		result = []int{}
	}

	paginator.Total = p.Count
	paginator.Items = result
	paginator.PerPage = p.Limit
	paginator.CurrentPage = p.Page
	paginator.Offset = offset
	paginator.TotalPage = int(math.Ceil(float64(p.Count) / float64(p.Limit)))

	return &paginator
}

func GetLimitOffset(page, limitIn int) (offset, limitOut int) {
	limitOut = limitIn

	if page < 1 {
		page = 1
	}

	if limitOut == 0 {
		limitOut = 10
	} else if limitOut > 30 {
		limitOut = 25
	}

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limitOut
	}

	return offset, limitOut
}
