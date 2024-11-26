package repositorySupport

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type SupportRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *SupportRepository {
	return &SupportRepository{
		repositoryPostgres.New(client),
	}
}
