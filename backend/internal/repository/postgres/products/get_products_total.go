package repositoryProducts

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) GetProductsTotal(ctx context.Context) (int64, error) {
	log.Printf("productsRepository.GetProductsTotal")

	stmt, args, err := squirrel.Select("count(*)").
		From(pgEntity.NewProductRow().Table() + " p").
		Join(fmt.Sprintf("%v pc on p.uid = pc.product_uid", pgEntity.NewProductCountRow(uuid.UUID{}, 0).Table())).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Printf("failed to get products total: %v", err)
		return 0, errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось получить количество продуктов")
	}

	var total int64

	rows, err := r.DB().Query(ctx, stmt, args...)

	for rows.Next() {
		rows.Scan(&total)
	}

	if err != nil {
		log.Printf("failed to get products total: %v", err)
		return 0, errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось получить количество продуктов")
	}

	return total, nil
}
