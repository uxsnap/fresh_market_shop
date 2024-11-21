package pgEntity

import (
	"strings"

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
	NewMaker[CategoryRow]

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

func (c *CategoryRow) ConditionNameLike() sq.Like {
	return sq.Like{
		"LOWER(name)": "%" + strings.ToLower(c.Name) + "%",
	}
}

func NewCategoriesRows() *Rows[*CategoryRow, entity.Category] {
	return &Rows[*CategoryRow, entity.Category]{}
}
