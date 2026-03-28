package users

import (
	"errors"
	"net/mail"
	"unicode"
	"unicode/utf8"
)

func runValidators(validators ...func() error) error {
	for _, v := range validators {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}

func validateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("please enter a valid email")
	}
	return nil
}

func validateUsername(username string) error {
	if utf8.RuneCountInString(username) < 3 {
		return errors.New("username must be at least 3 characters in length")
	}
	if utf8.RuneCountInString(username) > 16 {
		return errors.New("username must be equal to or less than 16 characters in length")
	}
	return nil
}

func validatePassword(pass string) error {
	var passwordHasUpper, passwordHasLower, passwordHasSymbol, passwordHasDigit bool
	for _, p := range pass {
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
	case utf8.RuneCountInString(pass) < 8:
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
