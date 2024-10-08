package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const deliveryAddressesTable = "delivery_addresses"

type DeliveryAddressRow struct {
	Uid         pgtype.UUID
	UserUid     pgtype.UUID
	Latitude    float64
	Longitude   float64
	CityName    string
	StreetName  string
	HouseNumber int64
	Building    int64
	Floor       int64
	Apartment   int64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func NewDeliveryAddressRow() *DeliveryAddressRow {
	return &DeliveryAddressRow{}
}

func (dr *DeliveryAddressRow) FromEntity(address entity.DeliveryAddress) *DeliveryAddressRow {
	dr.Uid = pgUidFromUUID(address.Uid)
	dr.UserUid = pgUidFromUUID(address.UserUid)
	dr.Latitude = address.Latitude
	dr.Longitude = address.Longitude
	dr.CityName = address.CityName
	dr.StreetName = address.StreetName
	dr.HouseNumber = address.HouseNumber
	dr.Building = address.Building
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
		Uid:         dr.Uid.Bytes,
		UserUid:     dr.UserUid.Bytes,
		Latitude:    dr.Latitude,
		Longitude:   dr.Longitude,
		CityName:    dr.CityName,
		StreetName:  dr.StreetName,
		HouseNumber: dr.HouseNumber,
		Building:    dr.Building,
		Floor:       dr.Floor,
		Apartment:   dr.Apartment,
		CreatedAt:   dr.CreatedAt.Time,
		UpdatedAt:   dr.UpdatedAt.Time,
	}
}

var deliveryAddressesTableColumns = []string{
	"uid", "user_uid", "latitude", "longitude", "city_name", "street_name", "house_number",
	"building", "floor", "apartment", "created_at", "updated_at",
}

func (dr *DeliveryAddressRow) Values() []interface{} {
	return []interface{}{
		dr.Uid, dr.UserUid, dr.Latitude, dr.Longitude, dr.CityName,
		dr.StreetName, dr.HouseNumber, dr.Building, dr.Floor,
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
		&dr.Uid, &dr.UserUid, &dr.Latitude, &dr.Longitude, &dr.CityName,
		&dr.StreetName, &dr.HouseNumber, &dr.Building, &dr.Floor,
		&dr.Apartment, &dr.CreatedAt, &dr.UpdatedAt,
	)
}

func (dr *DeliveryAddressRow) ColumnsForUpdate() []string {
	return []string{
		"latitude", "longitude", "city_name", "street_name", "house_number",
		"building", "floor", "apartment", "updated_at",
	}
}

func (dr *DeliveryAddressRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		dr.Latitude, dr.Longitude, dr.CityName,
		dr.StreetName, dr.HouseNumber, dr.Building, dr.Floor,
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
		"user_uid": dr.Uid,
	}
}

type DeliveryAddressRows struct {
	rows []*DeliveryAddressRow
}

func NewDeliveryAddressRows() *DeliveryAddressRows {
	return &DeliveryAddressRows{}
}

func (rs *DeliveryAddressRows) ScanAll(rows pgx.Rows) error {
	rs.rows = []*DeliveryAddressRow{}
	for rows.Next() {
		newRow := &DeliveryAddressRow{}
		if err := newRow.Scan(rows); err != nil {
			return err
		}

		rs.rows = append(rs.rows, newRow)
	}
	return nil
}

func (rs *DeliveryAddressRows) ToEntity() []entity.DeliveryAddress {
	if len(rs.rows) == 0 {
		return nil
	}

	res := make([]entity.DeliveryAddress, len(rs.rows))
	for i := 0; i < len(rs.rows); i++ {
		res[i] = rs.rows[i].ToEntity()
	}
	return res
}
