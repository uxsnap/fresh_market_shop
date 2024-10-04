package useCaseProducts

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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
	if uuid.Equal(product.CategoryUid, uuid.UUID{}) {
		return errors.New("category uid is empty")
	}
	return nil
}

func validateCategory(category entity.Category) error {
	if len(category.Name) == 0 {
		return errors.New("empty category name")
	}
	if len(category.Description) == 0 {
		return errors.New("empty category description")
	}
	return nil
}
