package controller

import (

	"backend/internal/api/util"
	"github.com/labstack/echo/v4"
)

func UserCreate(ctx echo.Context) error {
	return util.Success(ctx, "hi")
}