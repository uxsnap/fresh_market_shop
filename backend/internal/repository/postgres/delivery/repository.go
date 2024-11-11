package repositoryDelivery

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type DeliveryRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *DeliveryRepository {
	return &DeliveryRepository{
		repositoryPostgres.New(client),
	}
}
