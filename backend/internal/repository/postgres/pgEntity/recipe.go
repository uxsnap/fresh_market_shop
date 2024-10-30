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
	Products    []ProductRow
	Steps       []RecipeStepRow
	ImgPath     string
}

func NewRecipeRow() *RecipeRow {
	return &RecipeRow{}
}

func (rr *RecipeRow) FromEntity(recipe entity.Recipe) (*RecipeRow, error) {
	rr.Uid = pgUidFromUUID(recipe.Uid)
	rr.Name = recipe.Name
	rr.CookingTime = pgtype.Interval{
		Status:       pgtype.Present,
		Microseconds: int64(recipe.CookingTime) / 1000,
	}
	rr.ImgPath = recipe.ImgPath

	rr.Products = make([]ProductRow, len(recipe.Products))
	rr.Steps = make([]RecipeStepRow, len(recipe.Steps))

	for productInd, product := range recipe.Products {
		rr.Products[productInd] = *NewProductRow().FromEntity(product)
	}

	for stepInd, step := range recipe.Steps {
		rr.Steps[stepInd] = RecipeStepRow{
			RecipeUid: pgtype.UUID{
				Status: pgtype.Present,
				Bytes:  step.RecipeUid,
			},
			Step:        step.Step,
			Description: step.Description,
		}
	}

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

func (rr *RecipeRow) ToEntity() (entity.Recipe, error) {
	r := entity.Recipe{
		Uid:         rr.Uid.Bytes,
		Name:        rr.Name,
		CookingTime: rr.CookingTime.Microseconds,
		CreatedAt:   rr.CreatedAt.Time,
		UpdatedAt:   rr.UpdatedAt.Time,
		ImgPath:     rr.ImgPath,
		Products:    make([]entity.Product, len(rr.Products)),
		Steps:       make([]entity.RecipeStep, len(rr.Steps)),
	}

	for productInd, product := range rr.Products {
		r.Products[productInd] = product.ToEntity()
	}

	for stepInd, step := range rr.Steps {
		r.Steps[stepInd] = entity.RecipeStep{
			RecipeUid:   step.RecipeUid.Bytes,
			Step:        step.Step,
			Description: step.Description,
		}
	}

	return r, nil
}

var recipesTableColumns = []string{"uid", "name", "created_at", "updated_at", "img_path", "cooking_time"}

func (rr *RecipeRow) Values() []interface{} {
	return []interface{}{
		rr.Uid, rr.Name, rr.CookingTime, rr.CreatedAt, rr.UpdatedAt, rr.ImgPath, rr.CookingTime, rr.Products, rr.Steps,
	}
}

func (rr *RecipeRow) ValuesToScan() []interface{} {
	return []interface{}{&rr.Uid, &rr.Name, &rr.CreatedAt, &rr.UpdatedAt, &rr.ImgPath, &rr.CookingTime}
}

func (rr *RecipeRow) Columns() []string {
	return recipesTableColumns
}

func (rr *RecipeRow) Table() string {
	return recipesTableName
}

func (rr *RecipeRow) Scan(row pgx.Row) error {
	return row.Scan(&rr.Uid, &rr.Name, &rr.CreatedAt, &rr.UpdatedAt, &rr.ImgPath, &rr.CookingTime)
}

func (rr *RecipeRow) ColumnsForUpdate() []string {
	return []string{"name", "description", "cooking_time", "updated_at", "img_path"}
}

func (rr *RecipeRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		rr.Name, rr.UpdatedAt, rr.ImgPath, rr.CookingTime,
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

type RecipesRows struct {
	Rows []*RecipeRow
}

func NewRecipesRows() *RecipesRows {
	return &RecipesRows{}
}

func (rr *RecipesRows) ScanAll(rows pgx.Rows) error {
	rr.Rows = []*RecipeRow{}

	for rows.Next() {
		newRow := &RecipeRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		rr.Rows = append(rr.Rows, newRow)
	}

	return nil
}

func (rr *RecipesRows) ToEntity() ([]entity.Recipe, error) {
	if len(rr.Rows) == 0 {
		return nil, nil
	}

	res := make([]entity.Recipe, len(rr.Rows))
	for i := 0; i < len(rr.Rows); i++ {
		val, err := rr.Rows[i].ToEntity()
		if err != nil {
			return nil, err
		}
		res[i] = val
	}
	return res, nil
}
