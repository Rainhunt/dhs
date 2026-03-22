package main

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rainhunt/dhs/internal/config"
	"github.com/rainhunt/dhs/internal/db"
	"github.com/rainhunt/dhs/internal/users"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	r := e.Group("")

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pool, err := db.NewPgxPool(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	userDomain := users.NewUserDomain(pool, cfg.Jwt.Secret)
	userDomain.Register(r)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.Server.Port)))
}
