package repositoryUsers

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type UsersRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *UsersRepository {
	return &UsersRepository{
		repositoryPostgres.New(client),
	}
}
