package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const deliveryTable = "delivery"

type DeliveryRow struct {
	Uid           pgtype.UUID
	OrderUid      pgtype.UUID
	FromLongitude float64
	FromLatitude  float64
	ToLongitude   float64
	ToLatitude    float64
	Address       string
	Receiver      string
	Time          pgtype.Interval
	Price         int64
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

func NewDeliveryRow() *DeliveryRow {
	return &DeliveryRow{}
}

func (dr *DeliveryRow) FromEntity(delivery entity.Delivery) *DeliveryRow {
	dr.Uid = pgUidFromUUID(delivery.Uid)
	dr.OrderUid = pgUidFromUUID(delivery.OrderUid)
	dr.FromLongitude = delivery.FromLongitude
	dr.FromLatitude = delivery.FromLatitude
	dr.ToLongitude = delivery.ToLongitude
	dr.ToLatitude = delivery.ToLatitude
	dr.Address = delivery.Address
	dr.Receiver = delivery.Receiver
	dr.Time = pgtype.Interval{
		Microseconds: delivery.Time,
		Status:       pgtype.Present,
	}
	dr.CreatedAt = pgtype.Timestamp{
		Time:   delivery.CreatedAt,
		Status: pgStatusFromTime(delivery.CreatedAt),
	}
	dr.UpdatedAt = pgtype.Timestamp{
		Time:   delivery.UpdatedAt,
		Status: pgStatusFromTime(delivery.UpdatedAt),
	}
	return dr
}

func (dr *DeliveryRow) ToEntity() entity.Delivery {
	return entity.Delivery{
		Uid:           dr.Uid.Bytes,
		OrderUid:      dr.OrderUid.Bytes,
		FromLongitude: dr.FromLongitude,
		FromLatitude:  dr.FromLatitude,
		ToLongitude:   dr.ToLongitude,
		ToLatitude:    dr.ToLatitude,
		Address:       dr.Address,
		Receiver:      dr.Receiver,
		Time:          dr.Time.Microseconds,
		CreatedAt:     dr.CreatedAt.Time,
		UpdatedAt:     dr.UpdatedAt.Time,
	}
}

var deliveryTableColumns = []string{
	"uid", "order_uid", "from_longitude", "from_latitude",
	"to_longitude", "to_latitude", "address", "receiver",
	"delivery_time", "price", "created_at", "updated_at",
}

func (dr *DeliveryRow) Values() []interface{} {
	return []interface{}{
		dr.Uid, dr.OrderUid, dr.FromLongitude, dr.FromLatitude,
		dr.ToLongitude, dr.ToLatitude, dr.Address, dr.Receiver,
		dr.Time, dr.Price, dr.CreatedAt, dr.UpdatedAt,
	}
}

func (dr *DeliveryRow) Columns() []string {
	return deliveryTableColumns
}

func (dr *DeliveryRow) Table() string {
	return deliveryTable
}

func (dr *DeliveryRow) Scan(row pgx.Row) error {
	return row.Scan(
		&dr.Uid, &dr.OrderUid, &dr.FromLongitude, &dr.FromLatitude,
		&dr.ToLongitude, &dr.ToLatitude, &dr.Address, &dr.Receiver,
		&dr.Time, &dr.Price, &dr.CreatedAt, &dr.UpdatedAt,
	)
}

func (dr *DeliveryRow) ColumnsForUpdate() []string {
	return []string{
		"from_longitude", "from_latitude",
		"to_longitude", "to_latitude", "address", "receiver",
		"delivery_time", "price", "updated_at",
	}
}

func (dr *DeliveryRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		dr.FromLongitude, dr.FromLatitude,
		dr.ToLongitude, dr.ToLatitude, dr.Address, dr.Receiver,
		dr.Time, dr.Price, dr.UpdatedAt,
	}
}

func (dr *DeliveryRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": dr.Uid,
	}
}

func (dr *DeliveryRow) ConditionOrderUidEqual() sq.Eq {
	return sq.Eq{
		"order_uid": dr.OrderUid,
	}
}
