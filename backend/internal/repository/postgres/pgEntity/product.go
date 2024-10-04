package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const ProductsTableName = "products"

var productsTableFields = []string{
	"uid", "category_uid", "name", "description", "ccal", "price", "created_at", "updated_at",
}

type ProductRow struct {
	Uid         pgtype.UUID
	CategoryUid pgtype.UUID
	Name        string
	Description string
	Ccal        int32
	Price       int64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func NewProductRow() *ProductRow {
	return &ProductRow{}
}

func (p *ProductRow) FromEntity(product entity.Product) *ProductRow {
	p.Uid = pgtype.UUID{
		Bytes:  product.Uid,
		Status: pgtype.Present,
	}
	p.CategoryUid = pgtype.UUID{
		Bytes:  product.CategoryUid,
		Status: pgtype.Present,
	}
	p.Name = product.Name
	p.Description = product.Description
	p.Ccal = product.Ccal
	p.Price = product.Price

	if product.CreatedAt.Unix() == 0 {
		p.CreatedAt = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	} else {
		p.CreatedAt = pgtype.Timestamp{
			Time:   product.CreatedAt,
			Status: pgtype.Present,
		}
	}

	if product.UpdatedAt.Unix() == 0 {
		p.UpdatedAt = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	} else {
		p.UpdatedAt = pgtype.Timestamp{
			Time:   product.UpdatedAt,
			Status: pgtype.Present,
		}
	}
	return p
}

func (p *ProductRow) ToEntity() entity.Product {
	return entity.Product{
		Uid:         p.Uid.Bytes,
		CategoryUid: p.CategoryUid.Bytes,
		Name:        p.Name,
		Description: p.Description,
		Ccal:        p.Ccal,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt.Time,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}

func (p *ProductRow) Values() []interface{} {
	return []interface{}{
		p.Uid, p.CategoryUid, p.Name, p.Description,
		p.Ccal, p.Price, p.CreatedAt, p.UpdatedAt,
	}
}

func (p *ProductRow) Columns() []string {
	return productsTableFields
}

func (p *ProductRow) Table() string {
	return ProductsTableName
}

func (p *ProductRow) ValuesForScan() []interface{} {
	return []interface{}{
		&p.Uid, &p.CategoryUid, &p.Name, &p.Description,
		&p.Ccal, &p.Price, &p.CreatedAt, &p.UpdatedAt,
	}
}

func (p *ProductRow) Scan(row pgx.Row) error {
	return row.Scan(p.ValuesForScan()...)
}

func (p *ProductRow) ColumnsForUpdate() []string {
	return []string{
		"category_uid", "name", "description", "ccal", "price", "updated_at",
	}
}

func (p *ProductRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		p.CategoryUid, p.Name, p.Description, p.Ccal,
		p.Price, p.UpdatedAt,
	}
}

type ProductRows struct {
	rows []*ProductRow
}

func NewProductRows() *ProductRows {
	return &ProductRows{}
}

func (pr *ProductRows) ScanAll(rows pgx.Rows) error {
	for rows.Next() {
		newRow := &ProductRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		pr.rows = append(pr.rows, newRow)
	}

	return nil
}

func (pr *ProductRows) ToEntity() []entity.Product {
	if len(pr.rows) == 0 {
		return nil
	}

	res := make([]entity.Product, len(pr.rows))
	for i := 0; i < len(pr.rows); i++ {
		res[i] = pr.rows[i].ToEntity()
	}
	return res
}

func (pr *ProductRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": pr.Uid}
}

func (pr *ProductRow) ConditionCategoryUidEqual() sq.Eq {
	return sq.Eq{"category_uid": pr.CategoryUid}
}
