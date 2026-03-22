package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type UserDomain struct {
	router *router
}

func NewUserDomain(pool *pgxpool.Pool, jwtSecret string) *UserDomain {
	repo := newRepository(pool)
	service := newService(repo, jwtSecret)
	handler := newHandler(service)
	router := newRouter(handler)
	return &UserDomain{router: router}
}

func (u *UserDomain) Register(parent *echo.Group) {
	g := parent.Group("/users")
	u.router.registerRoutes(g)
}
