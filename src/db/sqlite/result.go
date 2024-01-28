package sqlite

import "database/sql"

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type result struct {
	r sql.Result
}

func (r result) LastInsertId() (int64, error) {
	return r.r.LastInsertId()
}

func (r result) RowsAffected() (int64, error) {
	return r.r.RowsAffected()
}
