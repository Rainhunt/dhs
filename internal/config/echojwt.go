package config

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func GetCustomClaims(c *echo.Context) (*CustomClaims, error) {
	token, err := echo.ContextGet[*jwt.Token](c, "user")
	if err != nil {
		return nil, echo.ErrUnauthorized.Wrap(err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, echo.ErrForbidden
	}
	return claims, nil
}

func StringToPgUUID(plainText string) (pgtype.UUID, error) {
	var pgUuid pgtype.UUID
	if err := pgUuid.Scan(plainText); err != nil {
		return pgtype.UUID{}, err
	}
	if !pgUuid.Valid {
		return pgtype.UUID{}, errors.New("missing or malformed UUID")
	}
	return pgUuid, nil
}
