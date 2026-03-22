package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rainhunt/dhs/internal/db/queries"
)

type repository struct {
	q *queries.Queries
}

func newRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		q: queries.New(pool),
	}
}

func (r *repository) getUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return User{}, err
	}

	return User{
		Id:       row.ID,
		Email:    row.Email,
		Username: row.Username,
	}, nil
}

func (r *repository) getAuthCredentialsByEmail(ctx context.Context, email string) (authCredentials, error) {
	row, err := r.q.GetAuthCredentialsByEmail(ctx, email)
	if err != nil {
		return authCredentials{}, err
	}

	return authCredentials{
		Id:      row.ID,
		IsAdmin: row.IsAdmin,
		Pass:    row.Pass,
	}, nil
}

func (r *repository) listUsers(ctx context.Context, limit, offset int32) ([]User, error) {
	rows, err := r.q.ListUsers(ctx, queries.ListUsersParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
	}

	users := make([]User, len(rows))
	for i, row := range rows {
		users[i] = User{
			Id:       row.ID,
			Email:    row.Email,
			Username: row.Username,
		}
	}

	return users, nil
}

func (r *repository) createUser(ctx context.Context, email, username, pass string) (User, error) {
	row, err := r.q.CreateUser(ctx, queries.CreateUserParams{Email: email, Username: username, Pass: pass})
	if err != nil {
		return User{}, err
	}

	return User{
		Id:       row.ID,
		Email:    row.Email,
		Username: row.Username,
	}, nil
}

func (r *repository) editUser(ctx context.Context, email, username string) (User, error) {
	row, err := r.q.EditUser(ctx, queries.EditUserParams{Email: email, Username: username})
	if err != nil {
		return User{}, err
	}

	return User{
		Id:       row.ID,
		Email:    row.Email,
		Username: row.Username,
	}, nil
}

func (r *repository) editUserPass(ctx context.Context, pass string, id pgtype.UUID) (User, error) {
	row, err := r.q.EditUserPass(ctx, queries.EditUserPassParams{ID: id, Pass: pass})
	if err != nil {
		return User{}, err
	}

	return User{
		Id:       row.ID,
		Email:    row.Email,
		Username: row.Username,
	}, nil
}

func (r *repository) deleteUser(ctx context.Context, id pgtype.UUID) error {
	err := r.q.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
