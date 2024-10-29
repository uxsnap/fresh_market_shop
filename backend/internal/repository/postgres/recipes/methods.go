package repositoryRecipes

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

func (r *RecipesRepository) CreateRecipe(ctx context.Context, recipe entity.Recipe) error {
	log.Printf("recipesRepository.CreateRecipe: uid %s name %s", recipe.Uid, recipe.Name)

	row, err := pgEntity.NewRecipeRow().FromEntity(recipe)
	if err != nil {
		log.Printf("failed to convert recipe: %v", err)
		return errors.WithStack(err)
	}

	if err := r.Create(ctx, row); err != nil {
		log.Printf("failed to create recipe: %v", err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *RecipesRepository) GetRecipeByUid(ctx context.Context, uid uuid.UUID) (entity.Recipe, bool, error) {
	log.Printf("recipesRepository.GetRecipeByUid: uid %s", uid)

	row, _ := pgEntity.NewRecipeRow().FromEntity(entity.Recipe{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to get recipe %s: %v", uid, err)
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Recipe{}, false, nil
		}
		return entity.Recipe{}, false, errors.WithStack(err)
	}

	res, err := row.ToEntity()
	if err != nil {
		log.Printf("failed to convert recipe %s to entity: %v", uid, err)
		return entity.Recipe{}, true, errors.WithStack(err)
	}

	return res, true, nil
}

func (r *RecipesRepository) GetRecipesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Recipe, error) {
	log.Printf("recipesRepository.GetRecipesByNameLike: name %s", name)

	row, _ := pgEntity.NewRecipeRow().FromEntity(entity.Recipe{Name: name})
	rows := pgEntity.NewRecipesRows()

	limit, offset := qFilters.Limit, qFilters.Offset

	if limit != 0 {
		if err := r.GetWithLimit(ctx, row, rows, row.ConditionNameLike(), limit, offset); err != nil {
			log.Printf("failed to get recipes by name like %s: %v", name, err)
			return nil, errors.WithStack(err)
		}
	} else {
		if err := r.GetSome(ctx, row, rows, row.ConditionNameLike()); err != nil {
			log.Printf("failed to get recipes by name like %s: %v", name, err)
			return nil, errors.WithStack(err)
		}
	}

	res, err := rows.ToEntity()
	if err != nil {
		log.Printf("failed to convert recipes to entity: %v", err)
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (r *RecipesRepository) GetRecipes(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Recipe, error) {
	cookingTime := qFilters.CookingTime
	createdAfter := qFilters.CreatedAfter
	limit := qFilters.Limit
	offset := qFilters.Offset

	log.Printf("recipesRepository.GetRecipes: cookingTime %d createdAfter %s", cookingTime, createdAfter.String())

	row := pgEntity.NewRecipeRow()
	sql := sq.Select(row.Columns()...).
		From(row.Table()).
		PlaceholderFormat(sq.Dollar)

	if cookingTime != 0 {
		sql = sql.Where(sq.LtOrEq{
			"cooking_time": cookingTime,
		})
	}
	if createdAfter.Unix() != 0 {
		sql = sql.Where(sq.GtOrEq{
			"created_at": createdAfter,
		})
	}

	if limit != 0 {
		sql = sql.Limit(limit)
	}
	if offset != 0 {
		sql = sql.Offset(offset)
	}

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to get recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	recipeRows := pgEntity.NewRecipesRows()
	if err := recipeRows.ScanAll(rows); err != nil {
		log.Printf("failed to get recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	return recipeRows.ToEntity()
}

func (r *RecipesRepository) UpdateRecipe(ctx context.Context, recipe entity.Recipe) error {
	log.Printf("recipesRepository.UpdateRecipe: uid %s name %s", recipe.Uid, recipe.Name)

	row, err := pgEntity.NewRecipeRow().FromEntity(recipe)
	if err != nil {
		log.Printf("failed to convert recipe %s: %v", recipe.Uid, err)
		return errors.WithStack(err)
	}

	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update recipe %s: %v", recipe.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *RecipesRepository) DeleteRecipe(ctx context.Context, uid uuid.UUID) error {
	log.Printf("recipesRepository.DeleteRecipe: uid %s", uid)

	row, _ := pgEntity.NewRecipeRow().FromEntity(entity.Recipe{Uid: uid})

	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete recipe %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}
