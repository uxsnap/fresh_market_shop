package repositoryAddresses

import (
	"context"
	"log"

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
