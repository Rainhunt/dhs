package users

import (
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

type createUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

func (r *createUserDTO) validate() error {
	return runValidators(
		func() error { return validateEmail(r.Email) },
		func() error { return validateUsername(r.Username) },
		func() error { return validatePassword(r.Pass) },
	)
}

type editUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (r *editUserDTO) validate() error {
	return runValidators(
		func() error {
			if r.Email == "" {
				return nil
			}
			return validateEmail(r.Email)
		},
		func() error {
			if r.Username == "" {
				return nil
			}
			return validateUsername(r.Username)
		},
	)
}

type editUserPassDTO struct {
	Pass string `json:"pass"`
}

func (r *editUserPassDTO) validate() error {
	return validatePassword(r.Pass)
}

type loginDTO struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
