package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

const productsCountTableName = "products_count"

type ProductsCountRow struct {
	ProductUid    pgtype.UUID
	StockQuantity int64
}

func NewProductsCountRow(productUid uuid.UUID, count int64) *ProductsCountRow {
	return &ProductsCountRow{
		ProductUid: pgtype.UUID{
			Bytes:  productUid,
			Status: pgtype.Present,
		},
		StockQuantity: count,
	}
}

func (pc *ProductsCountRow) Count() int64 {
	return pc.StockQuantity
}

func (pc *ProductsCountRow) Values() []interface{} {
	return []interface{}{pc.ProductUid, pc.StockQuantity}
}

func (pc *ProductsCountRow) Columns() []string {
	return []string{"product_uid", "stock_quantity"}
}

func (pc *ProductsCountRow) Table() string {
	return productsCountTableName
}

func (pc *ProductsCountRow) Scan(row pgx.Row) error {
	return row.Scan(&pc.ProductUid, &pc.StockQuantity)
}

func (pc *ProductsCountRow) ColumnsForUpdate() []string {
	return []string{"stock_quantity"}
}

func (pc *ProductsCountRow) ValuesForUpdate() []interface{} {
	return []interface{}{pc.StockQuantity}
}

func (pc *ProductsCountRow) ConditionProductUidEqual() sq.Eq {
	return sq.Eq{"product_uid": pc.ProductUid}
}
