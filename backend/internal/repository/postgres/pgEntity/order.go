package pgEntity

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const OrdersTableName = "orders"

var ordersTableFields = []string{
	"uid", "user_uid", "num", "status", "created_at", "updated_at",
}

type OrderRow struct {
	Uid       pgtype.UUID
	UserUid   pgtype.UUID
	Num       int64
	Status    string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func NewOrderRow() *OrderRow {
	return &OrderRow{}
}

func (p *OrderRow) FromEntity(order entity.Order) (*OrderRow, error) {
	p.Uid = pgtype.UUID{
		Bytes:  order.Uid,
		Status: pgtype.Present,
	}

	p.UserUid = pgtype.UUID{
		Bytes:  order.UserUid,
		Status: pgtype.Present,
	}

	p.Status = order.Status

	if order.CreatedAt.Unix() <= 0 {
		p.CreatedAt = pgtype.Timestamp{
			Time:   time.Now().UTC(),
			Status: pgtype.Null,
		}
	} else {
		p.CreatedAt = pgtype.Timestamp{
			Time:   order.CreatedAt,
			Status: pgtype.Present,
		}
	}

	if order.UpdatedAt.Unix() <= 0 {
		p.UpdatedAt = pgtype.Timestamp{
			Time:   time.Now().UTC(),
			Status: pgtype.Null,
		}
	} else {
		p.UpdatedAt = pgtype.Timestamp{
			Time:   order.UpdatedAt,
			Status: pgtype.Present,
		}
	}

	return p, nil
}

func (p *OrderRow) ToEntity() entity.Order {
	return entity.Order{
		Uid:       p.Uid.Bytes,
		UserUid:   p.UserUid.Bytes,
		Num:       p.Num,
		Status:    p.Status,
		CreatedAt: p.CreatedAt.Time,
		UpdatedAt: p.UpdatedAt.Time,
	}
}

func (p *OrderRow) Values() []interface{} {
	return []interface{}{
		p.Uid, p.UserUid, p.Num, p.Status, p.CreatedAt.Time, p.UpdatedAt.Time,
	}
}

func (p *OrderRow) Columns() []string {
	return ordersTableFields
}

func (p *OrderRow) Table() string {
	return OrdersTableName
}

func (p *OrderRow) ValuesForScan() []interface{} {
	return []interface{}{
		&p.Uid, &p.UserUid, &p.Num, &p.Status, &p.CreatedAt, &p.UpdatedAt,
	}
}

func (p *OrderRow) Scan(row pgx.Row) error {
	return row.Scan(p.ValuesForScan()...)
}

func (p *OrderRow) ColumnsForUpdate() []string {
	return []string{
		"num", "sum", "updated_at",
	}
}

func (p *OrderRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		p.Num, p.Status, p.UpdatedAt,
	}
}

type OrderRows struct {
	rows []*OrderRow
}

func NewOrderRows() *OrderRows {
	return &OrderRows{}
}

func (pr *OrderRows) ScanAll(rows pgx.Rows) error {
	pr.rows = []*OrderRow{}
	for rows.Next() {
		newRow := &OrderRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		pr.rows = append(pr.rows, newRow)
	}

	return nil
}

func (pr *OrderRows) ToEntity() []entity.Order {
	if len(pr.rows) == 0 {
		return nil
	}

	res := make([]entity.Order, len(pr.rows))
	for i := 0; i < len(pr.rows); i++ {
		res[i] = pr.rows[i].ToEntity()
	}
	return res
}

func (pr *OrderRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": pr.Uid}
}
