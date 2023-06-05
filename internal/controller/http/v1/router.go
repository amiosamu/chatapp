package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo) {

	handler.Use(middleware.Recover())
	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })

}
