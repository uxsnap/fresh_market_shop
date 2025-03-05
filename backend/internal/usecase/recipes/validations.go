package useCaseRecipes

import (
	"errors"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateRecipe(recipe entity.Recipe) error {
	if recipe.Name == "" {
		return errors.New("длина должна быть больше 0")
	}

	if recipe.Ccal < 1 {
		return errors.New("каллорийность должна быть больше 0")
	}

	if recipe.CookingTime < 1 {
		return errors.New("время приготовления должно быть больше 0")
	}

	return nil
}

func validateRecipeSteps(rSteps []entity.RecipeStep) error {
	if len(rSteps) == 0 {
		return errors.New("длина шагов должна быть больше нуля")
	}

	for _, rS := range rSteps {
		if len(rS.Description) < 1 {
			return errors.New("длина описания должна быть больше нуля")
		}

		if rS.Step < 1 {
			return errors.New("номер шага должен быть больше 0")
		}
	}

	return nil
}
