package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type OrderProduct struct {
	Uid   pgtype.UUID
	Count int64
}

type OrderProducts struct {
	Rows []*OrderProduct
}

func NewOrderProductsRows() *OrderProducts {
	return &OrderProducts{}
}

func (op *OrderProducts) FromEntity(orderProducts entity.OrderProducts) *OrderProducts {
	op.Rows = make([]*OrderProduct, len(orderProducts.Products))

	for i, v := range orderProducts.Products {
		op.Rows[i] = &OrderProduct{
			Uid: pgtype.UUID{
				Bytes:  v.Uid,
				Status: pgtype.Present,
			},
			Count: v.Count,
		}
	}

	return op
}

func (p *OrderProduct) ValuesForScan() []interface{} {
	return []interface{}{
		&p.Uid, &p.Count,
	}
}

func (p *OrderProduct) Scan(row pgx.Row) error {
	return row.Scan(p.ValuesForScan()...)
}

func (pr *OrderProducts) ScanAll(rows pgx.Rows) error {
	pr.Rows = []*OrderProduct{}

	for rows.Next() {
		newRow := &OrderProduct{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		pr.Rows = append(pr.Rows, newRow)
	}

	return nil
}
