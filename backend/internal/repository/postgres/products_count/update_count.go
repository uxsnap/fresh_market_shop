package repositoryProductsCount

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (r *ProductsCountRepository) UpdateCount(ctx context.Context, productsCounts entity.ProductsCounts) error {
	log.Printf("productsCountRepository.UpdateCount")

	updateStrs := []string{}

	for _, pc := range productsCounts.Products {
		updateStrs = append(updateStrs, fmt.Sprintf(`
			('%v'::uuid, %v)
		`, pc.ProductUid, pc.Count))
	}

	_, err := r.DB().Exec(ctx, fmt.Sprintf(`
		update products_count as pc set
			stock_quantity = pc.stock_quantity - pc2.stock_quantity
		from (values %v) as pc2(product_uid, stock_quantity)
		where pc2.product_uid = pc.product_uid;
	`, strings.Join(updateStrs, ", ")))

	if err != nil {
		log.Printf("failed to updateCount %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось обновить количество продукта")
	}

	return nil
}
