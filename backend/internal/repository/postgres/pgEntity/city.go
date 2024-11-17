package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const cityTable = "cities"

type CityRow struct {
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

type CitiesRows struct {
	rows []*CityRow
}

func NewCitiesRows() *CitiesRows {
	return &CitiesRows{}
}

func (cr *CitiesRows) ScanAll(rows pgx.Rows) error {
	for rows.Next() {
		newRow := &CityRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		cr.rows = append(cr.rows, newRow)
	}

	return nil
}

func (cr *CitiesRows) ToEntity() []entity.City {
	if len(cr.rows) == 0 {
		return nil
	}

	res := make([]entity.City, len(cr.rows))
	for i := 0; i < len(cr.rows); i++ {
		res[i] = cr.rows[i].ToEntity()
	}
	return res
}

func (c *CityRow) ColumnsForUpdate() []string {
	return []string{
		"name",
	}
}

func (c *CityRow) ValuesForUpdate() []interface{} {
	return []interface{}{c.Name}
}
