package useCaseProducts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type UseCaseProducts struct {
	productsRepository ProductsRepository
}

func New(
	productsRepository ProductsRepository,
) *UseCaseProducts {
	return &UseCaseProducts{
		productsRepository: productsRepository,
	}
}

// эти методы переедут в файлы отдельные
func (s *UseCaseProducts) CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error) {
	product.Uid = uuid.NewV4()
	_ = s.productsRepository.CreateProduct(ctx, product)
	return product.Uid, nil
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