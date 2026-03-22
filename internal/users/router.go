package users

import (
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type router struct {
	handler *handler
}

func newRouter(handler *handler) *router {
	return &router{handler: handler}
}

func (r *router) registerRoutes(g *echo.Group, jwtCfg echojwt.Config) {
	auth := g.Group("", echojwt.WithConfig(jwtCfg))
	g.POST("/signup", r.handler.createUser)
	g.POST("/login", r.handler.login)
	auth.GET("/", r.handler.createUser)
}
