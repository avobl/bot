package db

import (
	"context"
	"database/sql"
)

type Stmt interface {
	QueryContext(ctx context.Context, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, args ...any) Row
	ExecContext(ctx context.Context, args ...any) (Result, error)
	Close() error
}

type stmt struct {
	stmt *sql.Stmt
}

func (s stmt) QueryRowContext(ctx context.Context, args ...any) Row {
	return NewRowWrap(s.stmt.QueryRowContext(ctx, args))
}

func (s stmt) QueryContext(ctx context.Context, args ...any) (Rows, error) {
	rows, err := s.stmt.QueryContext(ctx, args)
	if err != nil {
		return nil, err
	}
	return NewRowsWrap(rows), nil
}

func (s stmt) ExecContext(ctx context.Context, args ...any) (Result, error) {
	return s.stmt.ExecContext(ctx, args)
}

func (s stmt) Close() error {
	return s.stmt.Close()
}
