package repositoryCategories

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	log.Printf("categoriesRepository.GetAllCategories")

	categoryRow := pgEntity.NewCategoryRow()
	rows := pgEntity.NewCategoriesRows()

	sql, args, err := sq.Select(categoryRow.Columns()...).From(categoryRow.Table()).ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return nil, err
	}

	rs, err := r.DB().Query(ctx, sql, args...)
	if err != nil {
		log.Printf("failed to get all categories: %v", err)
		return nil, err
	}

	if err := rows.ScanAll(rs); err != nil {
		log.Printf("failed to get all categories: %v", err)
		return nil, err
	}

	return rows.ToEntity(), nil
}

func (r *CategoriesRepository) GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, bool, error) {
	log.Printf("categoriesRepository.GetCategoryByUid: uid: %s", uid)

	row := pgEntity.NewCategoryRow().FromEntity(entity.Category{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Category{}, false, nil
		}
		log.Printf("failed to get category by uid %s: %v", uid, err)
		return entity.Category{}, false, errors.WithStack(err)
	}

	return row.ToEntity(), true, nil
}

func (r *CategoriesRepository) GetCategoriesByNameLike(ctx context.Context, name string, limit uint64, offset uint64) ([]entity.Category, error) {
	log.Printf("categoriesRepository.GetCategoriesByNameLike: name: %s", name)

	row := pgEntity.NewCategoryRow().FromEntity(entity.Category{Name: name})
	rows := pgEntity.NewCategoriesRows()

	if limit != 0 {
		if err := r.GetWithLimit(ctx, row, rows, row.ConditionNameLike(), limit, offset); err != nil {
			log.Printf("failed to get GetCategoriesByNameLike: %v", err)
			return nil, errors.WithStack(err)
		}
	} else {
		if err := r.GetSome(ctx, row, rows, row.ConditionNameLike()); err != nil {
			log.Printf("failed to get GetCategoriesByNameLike: %v", err)
			return nil, errors.WithStack(err)
		}
	}

	return rows.ToEntity(), nil
}
