package sqlite

import "database/sql"

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Err() error
	Close() error
}

type rows struct {
	r *sql.Rows
}

func NewRowsWrap(r *sql.Rows) Rows {
	return rows{r: r}
}

func (r rows) Next() bool {
	return r.r.Next()
}

func (r rows) Scan(dest ...any) error {
	return r.r.Scan(dest...)
}

func (r rows) Err() error {
	return r.r.Err()
}

func (r rows) Close() error {
	return r.r.Close()
}
