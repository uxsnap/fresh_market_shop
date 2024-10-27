package repositoryProductsCount

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type ProductsCountRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *ProductsCountRepository {
	return &ProductsCountRepository{
		repositoryPostgres.New(client),
	}
}
