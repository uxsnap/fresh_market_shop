package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const deliveryAddressesTable = "delivery_addresses"

type DeliveryAddressRow struct {
	NewMaker[DeliveryAddressRow]

	Uid        pgtype.UUID
	UserUid    pgtype.UUID
	AddressUid pgtype.UUID
	Entrance   int64
	Code       int64
	Floor      int64
	Apartment  int64
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

func NewDeliveryAddressRow() *DeliveryAddressRow {
	return &DeliveryAddressRow{}
}

func (dr *DeliveryAddressRow) FromEntity(address entity.DeliveryAddress) *DeliveryAddressRow {
	dr.Uid = pgUidFromUUID(address.Uid)
	dr.UserUid = pgUidFromUUID(address.UserUid)
	dr.AddressUid = pgUidFromUUID(address.AddressUid)
	dr.Entrance = address.Entrance
	dr.Code = address.Code
	dr.Floor = address.Floor
	dr.Apartment = address.Apartment
	dr.CreatedAt = pgtype.Timestamp{
		Time:   address.CreatedAt,
		Status: pgStatusFromTime(address.CreatedAt),
	}
	dr.UpdatedAt = pgtype.Timestamp{
		Time:   address.UpdatedAt,
		Status: pgStatusFromTime(address.UpdatedAt),
	}
	return dr
}

func (dr *DeliveryAddressRow) ToEntity() entity.DeliveryAddress {
	return entity.DeliveryAddress{
		Uid:        dr.Uid.Bytes,
		UserUid:    dr.UserUid.Bytes,
		AddressUid: dr.AddressUid.Bytes,
		Entrance:   dr.Entrance,
		Code:       dr.Code,
		Floor:      dr.Floor,
		Apartment:  dr.Apartment,
		CreatedAt:  dr.CreatedAt.Time,
		UpdatedAt:  dr.UpdatedAt.Time,
	}
}

var deliveryAddressesTableColumns = []string{
	"uid", "user_uid", "address_uid", "entrance", "code", "floor", "apartment", "created_at", "updated_at",
}

func (dr *DeliveryAddressRow) Values() []interface{} {
	return []interface{}{
		dr.Uid, dr.UserUid, dr.AddressUid, dr.Entrance, dr.Code, dr.Floor,
		dr.Apartment, dr.CreatedAt, dr.UpdatedAt,
	}
}

func (dr *DeliveryAddressRow) Columns() []string {
	return deliveryAddressesTableColumns
}

func (dr *DeliveryAddressRow) Table() string {
	return deliveryAddressesTable
}

func (dr *DeliveryAddressRow) Scan(row pgx.Row) error {
	return row.Scan(
		&dr.Uid, &dr.UserUid, &dr.AddressUid, &dr.Entrance, &dr.Code, &dr.Floor,
		&dr.Apartment, &dr.CreatedAt, &dr.UpdatedAt,
	)
}

func (dr *DeliveryAddressRow) ValuesForScan() []interface{} {
	return []interface{}{
		&dr.Uid, &dr.UserUid, &dr.AddressUid, &dr.Entrance, &dr.Code, &dr.Floor,
		&dr.Apartment, &dr.CreatedAt, &dr.UpdatedAt,
	}
}

func (dr *DeliveryAddressRow) ColumnsForUpdate() []string {
	return []string{
		"address_uid", "entrance", "code", "floor", "apartment", "updated_at",
	}
}

func (dr *DeliveryAddressRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		dr.AddressUid, dr.Entrance, dr.Code, dr.Floor,
		dr.Apartment, dr.UpdatedAt,
	}
}

func (dr *DeliveryAddressRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": dr.Uid,
	}
}

func (dr *DeliveryAddressRow) ConditionUserUidEqual() sq.Eq {
	return sq.Eq{
		"user_uid": dr.UserUid,
	}
}

func NewDeliveryAddressRows() *Rows[*DeliveryAddressRow, entity.DeliveryAddress] {
	return &Rows[*DeliveryAddressRow, entity.DeliveryAddress]{}
}
