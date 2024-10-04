package repositoryCategories

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) DeleteCategory(ctx context.Context, uid uuid.UUID) error {
	log.Printf("categoriesRepository.DeleteCategory: uid: %s", uid)

	row := pgEntity.NewCategoryRow().FromEntity(entity.Category{Uid: uid})
	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete category %s: %v", uid, err)
		return errors.WithStack(err)
	}

	return nil
}
