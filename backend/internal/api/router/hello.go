package router

import (
	"backend/internal/api/controller"

	"github.com/labstack/echo/v4"
)

func HelloRoute(e *echo.Echo) {
	e.GET("/", controller.SayHello)
}
