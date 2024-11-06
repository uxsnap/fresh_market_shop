package repositoryCategories

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *CategoriesRepository) GetCategoriesByUserOrders(ctx context.Context, userUid uuid.UUID) ([]uuid.UUID, error) {
	log.Printf("categoriesRepository.GetCategoriesByUserOrders")

	categoryRow := pgEntity.NewCategoryRow()
	productRow := pgEntity.NewProductRow()
	orderProductsRow := pgEntity.NewOrderProductsRow()

	sql, args, err := sq.Select(categoryRow.Columns()...).
		Distinct().
		From(categoryRow.Table() + " c").
		Join(productRow.Table() + " p on p.category_uid = c.uid").
		Join(orderProductsRow.Table() + " op on op.product_uid = p.uid").
		ToSql()

	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return []uuid.UUID{}, err
	}

	rs, err := r.DB().Query(ctx, sql, args...)
	if err != nil {
		log.Printf("failed to get all categories: %v", err)
		return []uuid.UUID{}, err
	}

	categoryUids := []uuid.UUID{}

	for rs.Next() {
		uid := uuid.UUID{}

		err := rs.Scan(&uid)

		if err != nil {
			return []uuid.UUID{}, err
		}

		categoryUids = append(categoryUids, uid)
	}

	return categoryUids, nil
}
