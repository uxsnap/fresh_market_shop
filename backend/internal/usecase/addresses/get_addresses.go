package useCaseAddresses

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (s *UseCaseAddresses) GetAddresses(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Address, error) {
	log.Printf("ucOrders.GetAddresses; city: %v, name: %v", qFilters.CityUid, qFilters.Name)

	addresses, err := s.addressesRepository.GetAddresses(ctx, qFilters)
	if err != nil {
		log.Printf("failed to get addresses: %v", err)
		return []entity.Address{}, err
	}

	return addresses, nil
}
