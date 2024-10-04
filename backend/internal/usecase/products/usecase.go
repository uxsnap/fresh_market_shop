package useCaseProducts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type UseCaseProducts struct {
	productsRepository   ProductsRepository
	categoriesRepository CategoriesRepository
}

func New(
	productsRepository ProductsRepository,
	categoriesRepository CategoriesRepository,
) *UseCaseProducts {
	return &UseCaseProducts{
		productsRepository:   productsRepository,
		categoriesRepository: categoriesRepository,
	}
}

func (s *UseCaseProducts) UpdateProduct(ctx context.Context, product entity.Product) error {
	return nil
}

func (s *UseCaseProducts) GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, error) {
	return entity.Product{Uid: uid}, nil
}

func (s *UseCaseProducts) GetProductsWithPagination(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	return nil, nil
}

func (s *UseCaseProducts) DeleteProduct(ctx context.Context, uid uuid.UUID) error {
	return nil
}
