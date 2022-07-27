package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v4/stdlib" // pgx
)

const migrationDir string = "migrations"

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

	storage := &Storage{Connect: db}

	err = storage.migrate()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *Storage) migrate() error {
	driver, err := postgres.WithInstance(s.Connect, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationDir, migrationDir, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
