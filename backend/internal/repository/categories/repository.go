package repositoryCategories

import (
	"context"

	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
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

func (r *CategoriesRepository) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return nil, nil
}
