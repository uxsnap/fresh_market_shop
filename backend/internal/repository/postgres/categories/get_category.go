package repositoryCategories

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	log.Printf("categoriesRepository.GetAllCategories")

	categoryRow := pgEntity.NewCategoryRow()
	rows := pgEntity.NewCategoriesRows()

	if err := r.GetSome(ctx, categoryRow, rows, categoryRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to get all categories: %v", err)
		return nil, errors.WithStack(err)
	}

	return rows.ToEntity(), nil
}

func (r *CategoriesRepository) GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, error) {
	log.Printf("categoriesRepository.GetCategoryByUid: uid: %s", uid)

	row := pgEntity.NewCategoryRow().FromEntity(entity.Category{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to get category by uid %s: %v", uid, err)
		return entity.Category{}, errors.WithStack(err)
	}

	return row.ToEntity(), nil
}
