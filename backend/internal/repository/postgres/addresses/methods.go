package repositoryAddresses

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *AddressesRepository) GetCities(ctx context.Context) ([]entity.City, error) {
	log.Printf("addressesRepository.GetCities")

	cityRow := pgEntity.NewCityRow().FromEntity(entity.City{})
	cityRows := pgEntity.NewCitiesRows()

	if err := r.GetSome(ctx, cityRow, cityRows, nil); err != nil {
		log.Printf("failed to get cities")
		return []entity.City{}, errors.WithStack(err)
	}

	return cityRows.ToEntity(), nil
}

func (r *AddressesRepository) GetAddresses(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Address, error) {
	log.Printf("addressesRepository.GetAddresses; city: %v, name: %v", qFilters.CityUid, qFilters.Name)

	if uuid.Equal(qFilters.CityUid, uuid.UUID{}) {
		return nil, errors.New("empty city uid")
	}
	if len(qFilters.Name) == 0 {
		return nil, errors.New("empty street name")
	}

	addressRow := pgEntity.NewAddressRow().FromEntity(entity.Address{})

	cond := squirrel.And{
		squirrel.Eq{"city_uid": qFilters.CityUid},
	}

	if len(qFilters.HouseNumber) != 0 {
		cond = squirrel.And{cond, squirrel.Like{"lower(house_number)": strings.ToLower(qFilters.HouseNumber) + "%"}}
	} else {
		cond = squirrel.And{cond, squirrel.NotEq{"house_number": "NULL"}}
	}

	if qFilters.Limit != 0 {
		cond = squirrel.And{cond, squirrel.Eq{"limit": qFilters.Limit}}
	}
	if qFilters.Offset != 0 {
		cond = squirrel.And{cond, squirrel.Eq{"offset": qFilters.Offset}}
	}

	sql, args, err := squirrel.Select(
		withPrefix("a", addressRow.Columns())...,
	).From(
		addressRow.Table() + " a",
	).LeftJoin(
		"addresses_streets_vectors av on a.uid=av.address_uid",
	).Where(
		fmt.Sprintf("av.street_vector @@ plainto_tsquery('russian','%s')", qFilters.Name),
	).Where(cond).ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, sql, args...)
	if err != nil {
		log.Printf("failed to get addresses: %v", err)
		return nil, errors.WithStack(err)
	}

	addressRows := pgEntity.NewAddressesRows()
	if err := addressRows.ScanAll(rows); err != nil {
		log.Printf("failed to scan addresses rows: %v", err)
		return nil, errors.WithStack(err)
	}

	return addressRows.ToEntity(), nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
