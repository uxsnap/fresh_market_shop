package useCaseRecipes

import (
	"errors"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateRecipe(recipe entity.Recipe) error {
	if len(recipe.Name) == 0 {
		return errors.New("empty name")
	}
	if len(recipe.Description) == 0 {
		return errors.New("empty description")
	}
	if len(recipe.Products) == 0 {
		return errors.New("empty products")
	}

	return nil
}
