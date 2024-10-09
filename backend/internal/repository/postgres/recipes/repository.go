package repositoryRecipes

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	repositoryPostgres "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres"
)

type RecipesRepository struct {
	*repositoryPostgres.BasePgRepository
}

func New(client DBclient.ClientDB) *RecipesRepository {
	return &RecipesRepository{
		repositoryPostgres.New(client),
	}
}
