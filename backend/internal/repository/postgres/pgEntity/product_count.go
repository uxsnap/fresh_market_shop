package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const productsCountTableName = "products_count"

type ProductCountRow struct {
	ProductUid    pgtype.UUID
	StockQuantity int64
}

func NewProductCountRow(productUid uuid.UUID, count int64) *ProductCountRow {
	return &ProductCountRow{
		ProductUid: pgtype.UUID{
			Bytes:  productUid,
			Status: pgtype.Present,
		},
		StockQuantity: count,
	}
}

func (pc *ProductCountRow) New() *ProductCountRow {
	return &ProductCountRow{}
}

func (pc *ProductCountRow) Count() int64 {
	return pc.StockQuantity
}

func (pc *ProductCountRow) Values() []interface{} {
	return []interface{}{pc.ProductUid, pc.StockQuantity}
}

func (pc *ProductCountRow) Columns() []string {
	return []string{"product_uid", "stock_quantity"}
}

func (pc *ProductCountRow) Table() string {
	return productsCountTableName
}

func (pc *ProductCountRow) Scan(row pgx.Row) error {
	return row.Scan(pc.ValuesForScan()...)
}

func (pc *ProductCountRow) ColumnsForUpdate() []string {
	return []string{"stock_quantity"}
}

func (pc *ProductCountRow) ValuesForUpdate() []interface{} {
	return []interface{}{pc.StockQuantity}
}

func (pc *ProductCountRow) ValuesForScan() []interface{} {
	return []interface{}{
		&pc.ProductUid, &pc.StockQuantity,
	}
}

func (pc *ProductCountRow) ConditionProductUidEqual() sq.Eq {
	return sq.Eq{"product_uid": pc.ProductUid}
}

func (op *ProductCountRow) FromEntity(productCount entity.ProductCount) *ProductCountRow {
	return &ProductCountRow{
		ProductUid: pgtype.UUID{
			Bytes:  productCount.ProductUid,
			Status: pgtype.Present,
		},
		StockQuantity: productCount.Count,
	}
}

func (op *ProductCountRow) ToEntity() entity.ProductCount {
	return entity.ProductCount{
		ProductUid: op.ProductUid.Bytes,
		Count:      op.StockQuantity,
	}
}

func NewProductCountRows() *Rows[*ProductCountRow, entity.ProductCount] {
	return &Rows[*ProductCountRow, entity.ProductCount]{}
}
