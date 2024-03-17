package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/avobl/bot/src/config"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// Connection - Initialized database connection.
	Connection Provider

	// ErrNoRows decouples dependency on sql package.
	ErrNoRows = sql.ErrNoRows
)

// Provider is an interface that wraps the standard sql.DB interface.
type Provider interface {
	PrepareContext(ctx context.Context, query string) (Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type sqlite struct {
	connection *sql.DB
}

// Load initializes the database connection.
func Load(ctx context.Context) error {
	db, err := sql.Open("sqlite3", config.C.SQLite.Dbname)
	if err != nil {
		return fmt.Errorf("sql open: %w", err)
	}

	go func() {
		<-ctx.Done()
		if err = db.Close(); err != nil {
			slog.WarnContext(ctx, "sqlite close: %v", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(config.C.SQLite.MaxIdleConns)
	db.SetMaxOpenConns(config.C.SQLite.MaxOpenConns)

	Connection = &sqlite{connection: db}

	return nil
}

func (sq *sqlite) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	st, err := sq.connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmt{stmt: st}, nil
}

func (sq *sqlite) ExecContext(ctx context.Context, query string, args ...any) (Result, error) {
	res, err := sq.connection.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &result{r: res}, nil
}

func (sq *sqlite) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	rs, err := sq.connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &rows{r: rs}, nil
}

func (sq *sqlite) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	return sq.connection.QueryRowContext(ctx, query, args...)
}

func (sq *sqlite) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	txn, err := sq.connection.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &tx{tx: txn}, nil
}
