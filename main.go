package main

import (
	"example/config"
	"example/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	cfg := config.New()
	cfg.SetupConfig()
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | [${method}] ${uri} | ${status} | ${latency_human} | error=${error}\n",
	}))

	router.RegisterRouter(e)

	e.Logger.Fatal(e.Start(":8003"))
}
