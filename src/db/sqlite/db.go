package sqlite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// ErrNoRows decouples dependency on sql package.
var ErrNoRows = sql.ErrNoRows

// Provider is an interface that wraps the standard sql.DB interface.
type Provider interface {
	Init() error
	Close() error
	PrepareContext(ctx context.Context, query string) (Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type Config struct {
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}

type sqlite struct {
	connection *sql.DB
	config     *Config
}

// GetProvider returns not initialized Provider.
func GetProvider(conf *Config) Provider {
	return &sqlite{config: conf}
}

// Init initializes the database connection.
func (sq *sqlite) Init() error {
	db, err := sql.Open("sqlite3", sq.config.DBName)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(sq.config.MaxIdleConns)
	db.SetMaxOpenConns(sq.config.MaxOpenConns)

	sq.connection = db

	return nil
}

func (sq *sqlite) Close() error {
	return sq.connection.Close()
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
