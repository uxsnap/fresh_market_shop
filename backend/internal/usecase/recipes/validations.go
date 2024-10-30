package useCaseRecipes

import (
	"errors"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateRecipe(recipe entity.Recipe) error {
	if len(recipe.Name) == 0 {
		return errors.New("empty name")
	}

	return nil
}
