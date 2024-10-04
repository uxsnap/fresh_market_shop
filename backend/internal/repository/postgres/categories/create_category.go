package repositoryCategories

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) CreateCategory(ctx context.Context, category entity.Category) error {
	log.Printf("categoriesRepository.CreateCategory (name: %s)", category.Name)

	row := pgEntity.NewCategoryRow().FromEntity(category)
	if err := r.Create(ctx, row); err != nil {
		log.Printf("failed to create category (name: %s): %v", category.Name, err)
		return errors.WithStack(err)
	}

	return nil
}
