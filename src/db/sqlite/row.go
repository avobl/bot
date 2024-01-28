package sqlite

import "database/sql"

type Row interface {
	Scan(dest ...any) error
	Err() error
}

func NewRowWrap(r *sql.Row) Row {
	return row{r: r}
}

type row struct {
	r *sql.Row
}

func (r row) Scan(dest ...any) error {
	return r.r.Scan(dest)
}

func (r row) Err() error {
	return r.r.Err()
}
