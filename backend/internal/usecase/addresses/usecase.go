package useCaseAddresses

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseAddresses struct {
	addressesRepository AddressesRepository
	txManager           *transaction.Manager
}

func New(
	addressesRepository AddressesRepository,
	txManager *transaction.Manager,
) *UseCaseAddresses {
	return &UseCaseAddresses{
		addressesRepository: addressesRepository,
		txManager:           txManager,
	}
}
