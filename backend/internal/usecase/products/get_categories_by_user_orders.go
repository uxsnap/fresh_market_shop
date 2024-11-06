package useCaseProducts

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (uc *UseCaseProducts) GetCategoriesByUserOrders(ctx context.Context, userUid uuid.UUID) ([]uuid.UUID, error) {
	log.Printf("ucProducts.GetCategoriesByUserOrders: userUid %s", userUid)

	categoriesUids, err := uc.categoriesRepository.GetCategoriesByUserOrders(ctx, userUid)

	if err != nil {
		log.Printf("failed to get categories by orders: %v", err)
		return nil, errorWrapper.NewError(errorWrapper.CategoriesError, "не удалось получить категории по заказу")
	}

	return categoriesUids, nil
}
