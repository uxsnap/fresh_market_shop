package repositoryCategories

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) UpdateCategory(ctx context.Context, category entity.Category) error {
	log.Printf("categoriesRepository.UpdateCategory (uid: %s)", category.Uid)

	row := pgEntity.NewCategoryRow().FromEntity(category)
	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update category (uid: %s): %v", category.Uid, err)
		return errors.WithStack(err)
	}

	return nil
}
