package main

import (
	"context"
	"log"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rainhunt/dhs/internal/config"
	"github.com/rainhunt/dhs/internal/db"
	"github.com/rainhunt/dhs/internal/users"
)

func main() {
	e := echo.New()
	r := e.Group("")

	r.Use(middleware.Recover())
	r.Use(middleware.RequestLogger())

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pool, err := db.NewPgxPool(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	jwtCfg := config.NewEchoJWTConfig(cfg.Jwt.Secret)

	userDomain := users.NewUserDomain(pool, cfg.Jwt.Secret)
	userDomain.Register(r, jwtCfg)

	sc := echo.StartConfig{Address: ":" + strconv.Itoa(cfg.Server.Port)}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
