package helper

import (
	"context"
	"database/sql"
	"fmt"
)

type txKey struct{}

type Transaction struct{}

func (t Transaction) InTransaction(ctx context.Context, db *sql.DB, tFunc func(ctx context.Context) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback()

	err = tFunc(context.WithValue(ctx, txKey{}, tx))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t Transaction) QueryContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}

	return db.QueryContext(ctx, query, args...)
}

func (t Transaction) QueryRowContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) *sql.Row {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return db.QueryRowContext(ctx, query, args...)
}

func (t Transaction) ExecContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return db.ExecContext(ctx, query, args...)
}

func (t Transaction) extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
