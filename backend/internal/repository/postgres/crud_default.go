package repositoryPostgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type Row interface {
	Values() []interface{}
	Columns() []string
	Table() string
	Scan(row pgx.Row) error
	ColumnsForUpdate() []string
	ValuesForUpdate() []interface{}
}

type Rows interface {
	ScanAll(rows pgx.Rows) error
}

func (r *BasePgRepository) Create(ctx context.Context, row Row) error {
	stmt, args, err := sq.Insert(row.Table()).
		PlaceholderFormat(sq.Dollar).
		Columns(row.Columns()...).
		Values(row.Values()...).ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB().Exec(ctx, stmt, args...)
	return err
}

func (r *BasePgRepository) GetOne(ctx context.Context, row Row, condition sq.Sqlizer) error {
	stmt, args, err := sq.Select(row.Columns()...).
		From(row.Table()).
		PlaceholderFormat(sq.Dollar).
		Where(condition).ToSql()
	if err != nil {
		return err
	}

	return row.Scan(r.DB().QueryRow(ctx, stmt, args...))
}

func (r *BasePgRepository) GetWithLimit(ctx context.Context, row Row, dest Rows, condition sq.Sqlizer, limit uint64, offset uint64) error {
	stmt, args, err := sq.Select(row.Columns()...).
		From(row.Table()).
		PlaceholderFormat(sq.Dollar).
		Where(condition).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		return err
	}

	fmt.Println(stmt)

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return dest.ScanAll(rows)
}

func (r *BasePgRepository) Update(ctx context.Context, row Row, condition sq.Sqlizer) error {
	columnsForUpdate := row.ColumnsForUpdate()
	valuesForUpdate := row.ValuesForUpdate()

	sqlBuilder := sq.Update(row.Table()).PlaceholderFormat(sq.Dollar)

	for i := 0; i < len(columnsForUpdate); i++ {
		sqlBuilder = sqlBuilder.Set(columnsForUpdate[i], valuesForUpdate[i])
	}

	sqlBuilder = sqlBuilder.Where(condition)

	stmt, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB().Exec(ctx, stmt, args...)
	return err
}

func (r *BasePgRepository) Delete(ctx context.Context, row Row, condition sq.Sqlizer) error {
	stmt, args, err := sq.Delete(row.Table()).PlaceholderFormat(sq.Dollar).Where(condition).ToSql()
	if err != nil {
		return err
	}
	_, err = r.DB().Exec(ctx, stmt, args...)
	return err
}

func (r *BasePgRepository) GetSome(ctx context.Context, row Row, dest Rows, condition sq.Sqlizer) error {
	stmt, args, err := sq.Select(row.Columns()...).
		From(row.Table()).
		PlaceholderFormat(sq.Dollar).
		Where(condition).ToSql()
	if err != nil {
		return err
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return dest.ScanAll(rows)
}

func (r *BasePgRepository) WithPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
