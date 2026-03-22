package users

import (
	"errors"
	"net/mail"
	"unicode"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id       pgtype.UUID `json:"id"`
	Email    string      `json:"email"`
	Username string      `json:"user"`
}

type authCredentials struct {
	Id      pgtype.UUID `json:"id"`
	IsAdmin bool        `json:"isAdmin"`
	Pass    string      `json:"pass"`
}

type customClaims struct {
	IsAdmin bool `json:"isAdmin"`
	jwt.RegisteredClaims
}

type createUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

func (r *createUserDTO) validate() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("please enter a valid email")
	}
	if utf8.RuneCountInString(r.Username) < 3 {
		return errors.New("username must be at least 3 characters in length")
	}
	if utf8.RuneCountInString(r.Username) > 16 {
		return errors.New("username must be equal to or less than 16 characters in length")
	}
	var passwordHasUpper, passwordHasLower, passwordHasSymbol, passwordHasDigit bool
	for _, p := range r.Pass {
		switch {
		case unicode.IsUpper(p):
			passwordHasUpper = true
		case unicode.IsLower(p):
			passwordHasLower = true
		case unicode.IsSymbol(p):
			passwordHasSymbol = true
		case unicode.IsDigit(p):
			passwordHasDigit = true
		}
	}
	switch {
	case utf8.RuneCountInString(r.Pass) < 8:
		return errors.New("password must be at least 8 characters in length")
	case !passwordHasUpper:
		return errors.New("password must contain an uppercase letter")
	case !passwordHasLower:
		return errors.New("password must contin a lowercase letter")
	case !passwordHasSymbol:
		return errors.New("password must contain a symbol")
	case !passwordHasDigit:
		return errors.New("password must contain a number")
	}
	return nil
}

type loginDTO struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
