package server

import (
	"backend/internal/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Run web server
func Run() {
	srv := echo.New()
	logger, _ := zap.NewProduction()
	srv.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	api := srv.Group("/api/v1")
	router.RegisterRoutes(api)
	srv.Logger.Fatal(srv.Start(":8080"))
}
