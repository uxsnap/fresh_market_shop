package repositoryOrders

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type OrderProductsRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *OrderProductsRepository {
	return &OrderProductsRepository{
		repositoryPostgres.New(client),
	}
}
