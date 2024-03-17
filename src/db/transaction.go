package db

import (
	"context"
	"database/sql"
)

type Tx interface {
	StmtContext(ctx context.Context, query string) (Stmt, error)
	Commit() error
	Rollback() error
}

type tx struct {
	tx *sql.Tx
}

func (t tx) StmtContext(ctx context.Context, query string) (Stmt, error) {
	st, err := t.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmt{stmt: st}, nil
}

func (t tx) Commit() error {
	return t.tx.Commit()
}

func (t tx) Rollback() error {
	return t.tx.Rollback()
}
