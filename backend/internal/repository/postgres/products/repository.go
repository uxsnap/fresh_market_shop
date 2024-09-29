package repositoryProducts

import (
	"context"

	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
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

func (r *ProductsRepository) CreateProduct(ctx context.Context, product entity.Product) error {
	return nil
}
