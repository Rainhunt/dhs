package users

import "github.com/labstack/echo/v4"

type router struct {
	handler *handler
}

func newRouter(handler *handler) *router {
	return &router{handler: handler}
}

func (r *router) registerRoutes(g *echo.Group) {
	g.POST("/signup", r.handler.createUser)
	g.POST("/login", r.handler.login)
}
