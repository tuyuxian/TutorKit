package router

import (
	"backend/internal/api/controller"

	"github.com/labstack/echo/v4"
)

func initUserRoute(e *echo.Group) {
	e.GET("/user/create", controller.UserCreate)
}
