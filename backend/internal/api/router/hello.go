package router

import (
	"backend/internal/api/controller"

	"github.com/labstack/echo/v4"
)

func initHelloRoute(e *echo.Group) {
	e.GET("/", controller.SayHello)
}
