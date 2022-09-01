package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	}
}
