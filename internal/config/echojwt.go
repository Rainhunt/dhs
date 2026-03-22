package config

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type CustomClaims struct {
	IsAdmin bool `json:"isAdmin"`
	jwt.RegisteredClaims
}

func NewEchoJWTConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		SigningKey: []byte(secret),
	}
}
