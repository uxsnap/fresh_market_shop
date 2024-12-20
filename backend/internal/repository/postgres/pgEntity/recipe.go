package pgEntity

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const recipesTableName = "recipes"

type RecipeRow struct {
	Uid         pgtype.UUID
	Name        string
	CookingTime pgtype.Interval
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Ccal        int64
}

func NewRecipeRow() *RecipeRow {
	return &RecipeRow{}
}

func (rr *RecipeRow) New() *RecipeRow {
	return &RecipeRow{}
}

func (rr *RecipeRow) FromEntity(recipe entity.Recipe) (*RecipeRow, error) {
	rr.Uid = pgUidFromUUID(recipe.Uid)
	rr.Name = recipe.Name
	rr.CookingTime = pgtype.Interval{
		Status:       pgtype.Present,
		Microseconds: int64(recipe.CookingTime) / 1000,
	}

	rr.Ccal = recipe.Ccal

	rr.CreatedAt = pgtype.Timestamp{
		Time:   recipe.CreatedAt,
		Status: pgStatusFromTime(recipe.CreatedAt),
	}

	rr.UpdatedAt = pgtype.Timestamp{
		Time:   recipe.UpdatedAt,
		Status: pgStatusFromTime(recipe.UpdatedAt),
	}
	return rr, nil
}

func (rr *RecipeRow) ToEntity() entity.Recipe {
	return entity.Recipe{
		Uid:         rr.Uid.Bytes,
		Name:        rr.Name,
		CookingTime: rr.CookingTime.Microseconds,
		CreatedAt:   rr.CreatedAt.Time,
		UpdatedAt:   rr.UpdatedAt.Time,
		Ccal:        rr.Ccal,
	}
}

var recipesTableColumns = []string{"uid", "name", "ccal", "created_at", "updated_at", "cooking_time"}

func (rr *RecipeRow) Values() []interface{} {
	return []interface{}{
		rr.Uid, rr.Name, rr.Ccal, rr.CookingTime, rr.CreatedAt, rr.UpdatedAt, rr.CookingTime,
	}
}

func (rr *RecipeRow) ValuesToScan() []interface{} {
	return []interface{}{&rr.Uid, &rr.Name, &rr.Ccal, &rr.CreatedAt, &rr.UpdatedAt, &rr.CookingTime}
}

func (rr *RecipeRow) Columns() []string {
	return recipesTableColumns
}

func (rr *RecipeRow) Table() string {
	return recipesTableName
}

func (rr *RecipeRow) Scan(row pgx.Row) error {
	return row.Scan(&rr.Uid, &rr.Name, &rr.Ccal, &rr.CreatedAt, &rr.UpdatedAt, &rr.CookingTime)
}

func (rr *RecipeRow) ColumnsForUpdate() []string {
	return []string{"name", "description", "ccal", "cooking_time", "updated_at", "img_path"}
}

func (rr *RecipeRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		rr.Name, rr.UpdatedAt, rr.CookingTime,
	}
}

func (rr *RecipeRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": rr.Uid,
	}
}

func (rr *RecipeRow) ConditionNameLike() sq.Like {
	return sq.Like{
		"LOWER(name)": "%" + strings.ToLower(rr.Name) + "%",
	}
}

func NewRecipesRows() *Rows[*RecipeRow, entity.Recipe] {
	return &Rows[*RecipeRow, entity.Recipe]{}
}
