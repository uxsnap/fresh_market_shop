package repositoryRecipes

import (
	"context"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *RecipesRepository) CreateRecipe(ctx context.Context, recipe entity.Recipe) error {
	log.Printf("recipesRepository.CreateRecipe: uid %s name %s", recipe.Uid, recipe.Name)

	row, err := pgEntity.NewRecipeRow().FromEntity(recipe)

	if err != nil {
		log.Printf("recipesRepository.CreateRecipe: failed to convert recipe: %v", err)
		return errors.WithStack(err)
	}

	if err := r.Create(ctx, row); err != nil {
		log.Printf("recipesRepository.CreateRecipe: failed to create recipe: %v", err)
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

	return row.ToEntity(), true, nil
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

	return rows.ToEntity(), nil
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

	if qFilters.Name != "" {
		sql = sql.Where(sq.Like{"lower(name)": "%" + strings.ToLower(qFilters.Name) + "%"})
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
		log.Printf("failed to scan recipes: %v", err)
		return nil, errors.WithStack(err)
	}

	return recipeRows.ToEntity(), nil
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

func (r *RecipesRepository) GetRecipeSteps(ctx context.Context, uid uuid.UUID) ([]entity.RecipeStep, error) {

	log.Printf("recipesRepository.GetRecipeSteps: recipe_uid %s", uid)

	row := pgEntity.NewRecipeStepRow()
	sql := sq.Select(row.Columns()...).
		From(row.Table()).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"recipe_uid": uid})

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to get recipe steps: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get recipe steps: %v", err)
		return nil, errors.WithStack(err)
	}

	recipeRows := pgEntity.NewRecipeStepRows()
	if err := recipeRows.ScanAll(rows); err != nil {
		log.Printf("failed to get recipe steps: %v", err)
		return nil, errors.WithStack(err)
	}

	return recipeRows.ToEntity(), nil
}

func (r *RecipesRepository) DeleteRecipePhotos(ctx context.Context, uid uuid.UUID, photosUids ...uuid.UUID) error {
	log.Printf("recipesRepository.DeleteRecipePhotos: recipe uid %s, recipe photos uids %v", uid, photosUids)

	photosUidsArgs := make([]pgtype.UUID, len(photosUids))
	for i := 0; i < len(photosUids); i++ {
		photosUidsArgs[i] = pgtype.UUID{Bytes: photosUids[i], Status: pgtype.Present}
	}

	// Only need to delete files but I don't want to
	return nil
}

func (r *RecipesRepository) GetRecipesTotal(ctx context.Context) (int64, error) {
	log.Printf("recipesRepository.GetRecipesTotal")

	stmt, args, err := sq.Select("count(*)").
		From(pgEntity.NewRecipeRow().Table()).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		log.Printf("failed to get recipes total: %v", err)
		return 0, errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось получить количество рецептов")
	}

	var total int64

	rows, err := r.DB().Query(ctx, stmt, args...)

	for rows.Next() {
		rows.Scan(&total)
	}

	if err != nil {
		log.Printf("failed to get recipesw total: %v", err)
		return 0, errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось получить количество рецептов")
	}

	return total, nil
}

func (r *RecipesRepository) AddRecipeSteps(ctx context.Context, uid uuid.UUID, rSteps []entity.RecipeStep) error {
	log.Printf("recipesRepository.AddRecipeSteps")

	row := pgEntity.NewRecipeStepRow()

	dSql := sq.Delete(row.Table()).Where(sq.Eq{"recipe_uid": uid})
	stmt, args, err := dSql.ToSql()
	if err != nil {
		log.Printf("failed to add recipe steps: %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось добавить шаги рецепта")
	}
	_, err = r.DB().Exec(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to add recipe steps: %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось добавить шаги рецепта")
	}

	sql := sq.Insert(row.Table()).PlaceholderFormat(sq.Dollar)

	for _, step := range rSteps {
		sql = sql.Values(row.FromEntity(step))
	}

	stmt, args, err = sql.ToSql()

	if err != nil {
		log.Printf("failed to add recipe steps: %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось добавить шаги рецепта")
	}

	_, err = r.DB().Exec(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to add recipe steps: %v", err)
		return errorWrapper.NewError(errorWrapper.ProductCountError, "не удалось добавить шаги рецепта")
	}

	return nil
}

func (r *RecipesRepository) DeleteRecipeStep(ctx context.Context, uid uuid.UUID, step int) error {
	log.Printf("recipesRepository.DeleteRecipeStep %v", step)

	row := pgEntity.NewRecipeStepRow().FromEntity(entity.RecipeStep{RecipeUid: uid, Step: int64(step)})

	if err := r.Delete(ctx, row, sq.Eq{"recipe_uid": uid, "step": step}); err != nil {
		log.Printf("failed to delete recipe %s: %v", uid, err)
		return errors.WithStack(err)
	}

	return nil
}
