package repositoryProducts

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository"
)

type ProductsRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *ProductsRepository {
	return &ProductsRepository{
		repositoryPostgres.New(client),
	}
}
