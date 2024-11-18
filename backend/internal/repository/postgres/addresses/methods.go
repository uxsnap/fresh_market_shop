package repositoryAddresses

import (
	"context"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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

	addressRow := pgEntity.NewAddressRow().FromEntity(entity.Address{})
	addressRows := pgEntity.NewAddressesRows()

	cond := squirrel.And{
		squirrel.Eq{"city_uid": qFilters.CityUid},
		squirrel.Like{"lower(street)": "%" + strings.ToLower(qFilters.Name) + "%"},
	}

	if qFilters.HouseNumber != "" {
		cond = squirrel.And{cond, squirrel.Like{"lower(house_number)": strings.ToLower(qFilters.HouseNumber) + "%"}}
	} else {
		cond = squirrel.And{cond, squirrel.NotEq{"house_number": "NULL"}}
	}

	if err := r.GetWithLimit(ctx, addressRow, addressRows, cond, qFilters.Limit, 0); err != nil {
		log.Printf("failed to get addresses, %v", err)
		return []entity.Address{}, errors.WithStack(err)
	}

	return addressRows.ToEntity(), nil
}
