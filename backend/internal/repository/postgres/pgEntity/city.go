package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const cityTable = "cities"

type CityRow struct {
	NewMaker[CityRow]

	Uid  pgtype.UUID
	Name string
}

func NewCityRow() *CityRow {
	return &CityRow{}
}

func (c *CityRow) FromEntity(city entity.City) *CityRow {
	c.Uid = pgUidFromUUID(city.Uid)
	c.Name = city.Name

	return c
}

func (c *CityRow) ToEntity() entity.City {
	return entity.City{
		Uid:  c.Uid.Bytes,
		Name: c.Name,
	}
}

var cityTableColumns = []string{
	"uid", "name",
}

func (c *CityRow) Values() []interface{} {
	return []interface{}{
		c.Uid, c.Name,
	}
}

func (c *CityRow) Columns() []string {
	return cityTableColumns
}

func (c *CityRow) Table() string {
	return cityTable
}

func (dr *CityRow) Scan(row pgx.Row) error {
	return row.Scan(
		&dr.Uid, &dr.Name,
	)
}

func (c *CityRow) ColumnsForUpdate() []string {
	return []string{
		"name",
	}
}

func (c *CityRow) ValuesForUpdate() []interface{} {
	return []interface{}{c.Name}
}

func NewCitiesRows() *Rows[*CityRow, entity.City] {
	return &Rows[*CityRow, entity.City]{}
}
