package repositoryAddresses

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type AddressesRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *AddressesRepository {
	return &AddressesRepository{
		repositoryPostgres.New(client),
	}
}
