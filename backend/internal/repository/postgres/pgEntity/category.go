package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const categoriesTableName = "categories"

var categoriesTableColumns = []string{
	"uid",
	"name",
	"description",
	"created_at",
	"updated_at",
}

type CategoryRow struct {
	Uid         pgtype.UUID
	Name        string
	Description string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func NewCategoryRow() *CategoryRow {
	return &CategoryRow{}
}

func (c *CategoryRow) FromEntity(category entity.Category) *CategoryRow {
	c.Uid = pgtype.UUID{
		Bytes:  category.Uid,
		Status: pgtype.Present,
	}
	c.Name = category.Name
	c.Description = category.Description

	if category.CreatedAt.Unix() != 0 {
		c.CreatedAt = pgtype.Timestamp{
			Time:   category.CreatedAt,
			Status: pgtype.Present,
		}
	} else {
		c.CreatedAt = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	}

	if category.UpdatedAt.Unix() != 0 {
		c.UpdatedAt = pgtype.Timestamp{
			Time:   category.UpdatedAt,
			Status: pgtype.Present,
		}
	} else {
		c.UpdatedAt = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	}
	return c
}

func (c *CategoryRow) ToEntity() entity.Category {
	return entity.Category{
		Uid:         c.Uid.Bytes,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt.Time,
		UpdatedAt:   c.UpdatedAt.Time,
	}
}

func (c *CategoryRow) Values() []interface{} {
	return []interface{}{c.Uid, c.Name, c.Description, c.CreatedAt, c.UpdatedAt}
}

func (c *CategoryRow) Columns() []string {
	return categoriesTableColumns
}

func (c *CategoryRow) Table() string {
	return categoriesTableName
}

func (c *CategoryRow) Scan(row pgx.Row) error {
	return row.Scan(&c.Uid, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
}

func (c *CategoryRow) ColumnsForUpdate() []string {
	return []string{
		"name",
		"description",
		"updated_at",
	}
}

func (c *CategoryRow) ValuesForUpdate() []interface{} {
	return []interface{}{c.Name, c.Description, c.UpdatedAt}
}

func (c *CategoryRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": c.Uid,
	}
}

type CategoriesRows struct {
	rows []*CategoryRow
}

func NewCategoriesRows() *CategoriesRows {
	return &CategoriesRows{}
}

func (cr *CategoriesRows) ScanAll(rows pgx.Rows) error {
	for rows.Next() {
		newRow := &CategoryRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		cr.rows = append(cr.rows, newRow)
	}

	return nil
}

func (cr *CategoriesRows) ToEntity() []entity.Category {
	if len(cr.rows) == 0 {
		return nil
	}

	res := make([]entity.Category, len(cr.rows))
	for i := 0; i < len(cr.rows); i++ {
		res[i] = cr.rows[i].ToEntity()
	}
	return res
}