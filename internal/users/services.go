package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo *repository
}

func newService(r *repository) *service {
	return &service{repo: r}
}

func (s *service) createUser(ctx context.Context, email, username, pass string) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	return s.repo.createUser(ctx, email, username, string(hashed))
}

func (s *service) updateUserPass(ctx context.Context, newPass string, id pgtype.UUID) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return User{}, nil
	}
	return s.repo.editUserPass(ctx, string(hashed), id)
}
