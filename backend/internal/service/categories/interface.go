package categoriesService

import (
	"context"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type CategoriesRepository interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
}
