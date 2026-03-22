package users

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rainhunt/dhs/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo      *repository
	jwtSecret string
}

func newService(r *repository, jwtSecret string) *service {
	return &service{
		repo:      r,
		jwtSecret: jwtSecret,
	}
}

func (s *service) createUser(ctx context.Context, email, username, pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	if _, err := s.repo.createUser(ctx, email, username, string(hashed)); err != nil {
		return "", err
	}
	return s.generateJWT(ctx, email, pass)
}

func (s *service) generateJWT(ctx context.Context, email, pass string) (string, error) {
	authCred, err := s.repo.getAuthCredentialsByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authCred.Pass), []byte(pass)); err != nil {
		return "", err
	}

	id, err := authCred.Id.UUIDValue()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.CustomClaims{
		IsAdmin: authCred.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: id.String(),
		},
	})
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *service) updateUserPass(ctx context.Context, newPass string, id pgtype.UUID) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return User{}, nil
	}
	return s.repo.editUserPass(ctx, string(hashed), id)
}
