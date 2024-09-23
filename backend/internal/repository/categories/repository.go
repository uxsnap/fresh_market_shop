package repositoryCategories

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository"
)

type CategoriesRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *CategoriesRepository {
	return &CategoriesRepository{
		repositoryPostgres.New(client),
	}
}
