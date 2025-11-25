package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := echo.New()
	r.HideBanner = true

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	r.Logger.Fatal(r.Start(":" + port))
}
