package pgEntity

import (
	"github.com/jackc/pgtype"
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

func NewOrderProductsRow() *OrderProductsRow {
	return &OrderProductsRow{}
}

type OrderProductsRows struct {
	rows []*OrderProductsRow
}

func NewOrderProductsRows() *OrderProductsRows {
	return &OrderProductsRows{}
}

func (p *OrderProductsRow) Table() string {
	return OrderProductsTableName
}
