package pgEntity

import (
	"encoding/json"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const OrderProductsTableName = "order_products"

var ordersProductsTableFields = []string{
	"order_uid",
	"product_uid",
	"count",
}

type OrderProductsRow struct {
	OrderUid   pgtype.UUID
	ProductUid pgtype.UUID
	Count      int64
}

func (o *OrderProductsRow) FromEntity(op entity.OrderProducts) *OrderProductsRow {
	o.OrderUid = pgtype.UUID{
		Bytes:  op.OrderUid,
		Status: pgtype.Present,
	}
	o.ProductUid = pgtype.UUID{
		Bytes:  op.ProductUid,
		Status: pgtype.Present,
	}

	o.Count = op.Count

	return o
}

func (o *OrderProductsRow) ToEntity() entity.OrderProducts {
	return entity.OrderProducts{
		OrderUid:   o.OrderUid.Bytes,
		ProductUid: o.ProductUid.Bytes,
		Count:      o.Count,
	}
}

func NewOrderProductsRow() *OrderProductsRow {
	return &OrderProductsRow{}
}

func (op *OrderProductsRow) New() *OrderProductsRow {
	return &OrderProductsRow{}
}

func (o *OrderProductsRow) Scan(row pgx.Row) error {
	return row.Scan(&o.OrderUid, &o.ProductUid, &o.Count)
}

type OrderProductsRows struct {
	*Rows[*OrderProductsRow, entity.OrderProducts]
}

func NewOrderProductsRows() *OrderProductsRows {
	return &OrderProductsRows{
		&Rows[*OrderProductsRow, entity.OrderProducts]{},
	}
}

func (p *OrderProductsRow) Columns() []string {
	return ordersProductsTableFields
}

func (p *OrderProductsRow) Table() string {
	return OrderProductsTableName
}

func (p *OrderProductsRow) ValuesForScan() []interface{} {
	return []interface{}{
		&p.OrderUid,
		&p.ProductUid,
		&p.Count,
	}
}

func (op *OrderProductsRows) FromJson(bts []byte) error {
	op.rows = nil

	if err := json.Unmarshal(bts, &op.rows); err != nil {
		return err
	}
	return nil
}
