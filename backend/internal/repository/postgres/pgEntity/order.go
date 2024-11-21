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
	"uid", "user_uid", "num", "sum", "status", "created_at", "updated_at",
}

type OrderRow struct {
	NewMaker[OrderRow]
	Uid       pgtype.UUID
	UserUid   pgtype.UUID
	Num       int64
	Sum       int64
	Status    string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func NewOrderRow() *OrderRow {
	return &OrderRow{}
}

func (p *OrderRow) FromEntity(order entity.Order) *OrderRow {
	p.Uid = pgtype.UUID{
		Bytes:  order.Uid,
		Status: pgtype.Present,
	}

	p.UserUid = pgtype.UUID{
		Bytes:  order.UserUid,
		Status: pgtype.Present,
	}

	p.Status = string(order.Status)
	p.Sum = order.Sum

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

	return p
}

func (p *OrderRow) ToEntity() entity.Order {
	return entity.Order{
		Uid:       p.Uid.Bytes,
		UserUid:   p.UserUid.Bytes,
		Num:       p.Num,
		Sum:       p.Sum,
		Status:    entity.OrderStatus(p.Status),
		CreatedAt: p.CreatedAt.Time,
		UpdatedAt: p.UpdatedAt.Time,
	}
}

func (p *OrderRow) Values() []interface{} {
	return []interface{}{
		p.Uid, p.UserUid, p.Num, p.Sum, p.Status, p.CreatedAt.Time, p.UpdatedAt.Time,
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
		&p.Uid, &p.UserUid, &p.Num, &p.Sum, &p.Status, &p.CreatedAt, &p.UpdatedAt,
	}
}

func (p *OrderRow) Scan(row pgx.Row) error {
	return row.Scan(p.ValuesForScan()...)
}

func (p *OrderRow) ColumnsForUpdate() []string {
	return []string{
		"num", "sum", "status", "updated_at",
	}
}

func (p *OrderRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		p.Num, p.Sum, p.Status, p.UpdatedAt,
	}
}

func (pr *OrderRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": pr.Uid}
}

func NewOrderRows() *Rows[*OrderRow, entity.Order] {
	return &Rows[*OrderRow, entity.Order]{}
}
