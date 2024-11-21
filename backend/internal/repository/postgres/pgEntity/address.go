package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const addressTable = "addresses"

type AddressRow struct {
	NewMaker[AddressRow]

	Uid         pgtype.UUID
	CityUid     pgtype.UUID
	Street      string
	HouseNumber string
	Latitude    float64
	Longitude   float64
}

func NewAddressRow() *AddressRow {
	return &AddressRow{}
}

func (a *AddressRow) FromEntity(address entity.Address) *AddressRow {
	a.Uid = pgUidFromUUID(address.Uid)
	a.CityUid = pgUidFromUUID(address.CityUid)
	a.Street = address.Street
	a.HouseNumber = address.HouseNumber
	a.Latitude = address.Latitude
	a.Longitude = address.Longitude

	return a
}

func (a *AddressRow) ToEntity() entity.Address {
	return entity.Address{
		Uid:         a.Uid.Bytes,
		CityUid:     a.CityUid.Bytes,
		Street:      a.Street,
		HouseNumber: a.HouseNumber,
		Latitude:    a.Latitude,
		Longitude:   a.Longitude,
	}
}

var addressTableColumns = []string{
	"uid", "city_uid", "street", "house_number", "latitude", "longitude",
}

func (a *AddressRow) Values() []interface{} {
	return []interface{}{
		a.Uid, a.CityUid, a.Street, a.HouseNumber, a.Latitude, a.Longitude,
	}
}

func (c *AddressRow) Columns() []string {
	return addressTableColumns
}

func (c *AddressRow) Table() string {
	return addressTable
}

func (a *AddressRow) Scan(row pgx.Row) error {
	return row.Scan(
		&a.Uid, &a.CityUid, &a.Street, &a.HouseNumber, &a.Latitude, &a.Longitude,
	)
}

func (a *AddressRow) ValuesForScan() []interface{} {
	return []interface{}{
		&a.Uid, &a.CityUid, &a.Street, &a.HouseNumber, &a.Latitude, &a.Longitude,
	}
}

func (a *AddressRow) ColumnsForUpdate() []string {
	return []string{
		"street", "house_number", "latitude", "longitude",
	}
}

func (a *AddressRow) ValuesForUpdate() []interface{} {
	return []interface{}{a.Street, a.HouseNumber, a.Latitude, a.Longitude}
}

func NewAddressesRows() *Rows[*AddressRow, entity.Address] {
	return &Rows[*AddressRow, entity.Address]{}
}
