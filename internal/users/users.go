package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
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

func (u *UserDomain) Register(parent *echo.Group, jwtCfg echojwt.Config) {
	g := parent.Group("/users")
	u.router.registerRoutes(g, jwtCfg)
}
