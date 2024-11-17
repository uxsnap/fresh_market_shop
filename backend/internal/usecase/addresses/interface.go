package useCaseAddresses

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type AddressesRepository interface {
	GetCities(ctx context.Context) ([]entity.City, error)
}
