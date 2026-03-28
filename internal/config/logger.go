package config

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var skipper = func(c *echo.Context) bool {
	// Skip health check endpoint
	return c.Request().URL.Path == "/health"
}

var LoggerConfig = middleware.RequestLoggerConfig{
	LogStatus: true,
	LogURI:    true,
	Skipper:   skipper,
	LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
		fmt.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
		return nil
	},
}
