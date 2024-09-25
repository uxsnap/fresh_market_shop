package productsService

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsService struct {
	productsRepository ProductsRepository
}

func New(
	productsRepository ProductsRepository,
) *ProductsService {
	return &ProductsService{
		productsRepository: productsRepository,
	}
}

// эти методы переедут в файлы отдельные
func (s *ProductsService) CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error) {
	product.Uid = uuid.NewV4()
	_ = s.productsRepository.CreateProduct(ctx, product)
	return product.Uid, nil
}

func (s *ProductsService) UpdateProduct(ctx context.Context, product entity.Product) error {
	return nil
}

func (s *ProductsService) GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, error) {
	return entity.Product{Uid: uid}, nil
}

func (s *ProductsService) GetProductsWithPagination(ctx context.Context, limit, offset int) ([]entity.Product, error) {
	return nil, nil
}

func (s *ProductsService) DeleteProduct(ctx context.Context, uid uuid.UUID) error {
	return nil
}
