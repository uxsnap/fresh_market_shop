package useCaseSupport

import (
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type UseCaseSupport struct {
	repository RepositorySupport
	txManager  transaction.Manager
}

func New(
	repository RepositorySupport,
	txManager transaction.Manager,
) *UseCaseSupport {
	return &UseCaseSupport{
		repository: repository,
		txManager:  txManager,
	}
}
