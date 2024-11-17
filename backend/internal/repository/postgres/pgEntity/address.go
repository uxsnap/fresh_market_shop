package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const addressTable = "addresses"

type AddressRow struct {
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

type AddressesRows struct {
	rows []*AddressRow
}

func NewAddressesRows() *AddressesRows {
	return &AddressesRows{}
}

func (cr *AddressesRows) ScanAll(rows pgx.Rows) error {
	for rows.Next() {
		newRow := &AddressRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		cr.rows = append(cr.rows, newRow)
	}

	return nil
}

func (cr *AddressesRows) ToEntity() []entity.Address {
	if len(cr.rows) == 0 {
		return []entity.Address{}
	}

	res := make([]entity.Address, len(cr.rows))
	for i := 0; i < len(cr.rows); i++ {
		res[i] = cr.rows[i].ToEntity()
	}
	return res
}

func (a *AddressRow) ColumnsForUpdate() []string {
	return []string{
		"street", "house_number", "latitude", "longitude",
	}
}

func (a *AddressRow) ValuesForUpdate() []interface{} {
	return []interface{}{a.Street, a.HouseNumber, a.Latitude, a.Longitude}
}
