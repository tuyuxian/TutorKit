package router

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {
	InitHelloRoute(e)
}
