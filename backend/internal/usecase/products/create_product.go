package useCaseProducts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

// эти методы переедут в файлы отдельные
func (s *UseCaseProducts) CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error) {
	product.Uid = uuid.NewV4()
	_ = s.productsRepository.CreateProduct(ctx, product)
	return product.Uid, nil
}
