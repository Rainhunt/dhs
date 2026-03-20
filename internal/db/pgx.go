package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rainhunt/dhs/internal/config"
)

func NewPgxPool(cfg *config.Config) (*pgxpool.Pool, error) {
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		dbCfg.User,
		os.Getenv(dbCfg.PassEnv),
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db at postgres://%s:XXX@%s:%d/%s: %w",
			dbCfg.User,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.Name,
			err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db ping failed at postgres://%s:XXX@%s:%d/%s: %w",
			dbCfg.User,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.Name,
			err)
	}
	fmt.Printf("Connected to PostgreSQL at postgres://%s:XXX@%s:%d/%s",
		dbCfg.User,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name)

	return pool, nil
}
