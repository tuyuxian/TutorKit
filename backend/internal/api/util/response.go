package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	ErrHint string      `json:"hint,omitempty"`
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Data:    data,
		Error:   "",
		Success: true,
	})
}

func Error(c echo.Context, status int, data string, err error) error {
	ret := Response{
		Data:    nil,
		Message: data,
		Success: false,
	}
	return c.JSON(status, ret)
}