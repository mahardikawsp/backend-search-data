package router

import "github.com/labstack/echo/v4"

func RegisterRouter(e *echo.Echo) {
	NewSearchRouter(e)
	NewMasterDataRouter(e)
	NewEconnectRouter(e)
	//TODO register another handler here if exists
	//...
}
