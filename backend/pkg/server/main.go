package server

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Run web server
func Run() {
	r := gin.Default()
	r.GET("/hello", handlers.HelloWorld())
	r.Run()
}
