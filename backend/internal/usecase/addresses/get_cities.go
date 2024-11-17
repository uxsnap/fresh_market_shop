package useCaseAddresses

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseAddresses) GetCities(ctx context.Context) ([]entity.City, error) {
	log.Printf("ucOrders.CreateOrder")

	cities, err := s.addressesRepository.GetCities(ctx)

	if err != nil {
		return []entity.City{}, err
	}

	return cities, nil
}
