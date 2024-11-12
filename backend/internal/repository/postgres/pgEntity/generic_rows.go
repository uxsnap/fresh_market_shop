package pgEntity

import (
	"github.com/jackc/pgx/v4"
)

type ToEntityConverter[K any] interface {
	ToEntity() K
}

type Row[K any] interface {
	Scan(row pgx.Row) error
	ToEntityConverter[K]
}

// T == *pgEntity....Row эта Row реализует методы => []T == []*pgEntity...Row
type Rows[T Row[K], K any] struct {
	rows []T
}

func NewRows[T Row[K], K any]() *Rows[T, K] {
	return &Rows[T, K]{}
}

func (rs *Rows[T, K]) ScanAll(rows pgx.Rows) error {
	rs.rows = []T{}

	for rows.Next() {
		// new (T) => new (*pgEntity.Row) ==> * (*pgEntity.Row) ==> * -> nil
		newRow := new(T)

		if err := T.Scan(*newRow, rows); err != nil {
			return err
		}
		rs.rows = append(rs.rows, *newRow)
	}
	return nil
}

func (rs *Rows[T, K]) ToEntity() []K {
	if len(rs.rows) == 0 {
		return nil
	}

	res := make([]K, len(rs.rows))
	for i := 0; i < len(rs.rows); i++ {
		res[i] = rs.rows[i].ToEntity() // ?why?
	}
	return res
}
