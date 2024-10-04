package useCaseProducts

import (
	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateProduct(product entity.Product) error {
	if len(product.Name) == 0 {
		return errors.New("invalid product name")
	}
	if len(product.Description) == 0 {
		return errors.New("description is empty")
	}
	if product.Price == 0 {
		return errors.New("invalid product price")
	}
	if product.Ccal == 0 {
		return errors.New("ccal is empty")
	}
	return nil
}
