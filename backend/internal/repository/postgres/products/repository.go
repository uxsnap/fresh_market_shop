package repositoryProducts

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type ProductsRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *ProductsRepository {
	return &ProductsRepository{
		repositoryPostgres.New(client),
	}
}
