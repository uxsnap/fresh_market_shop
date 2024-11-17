package useCaseAddresses

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type AddressesRepository interface {
	GetCities(ctx context.Context) ([]entity.City, error)
	GetAddresses(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Address, error)
}
