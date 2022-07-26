package service

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // pgx
)

type Storage struct {
	Connect *sql.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Storage{Connect: db}, nil
}
