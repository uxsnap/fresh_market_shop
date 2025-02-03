package pgEntity

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type RecipeStepRow struct {
	RecipeUid   pgtype.UUID
	Step        int64
	Description string
}

const recipeStepTableName = "recipes_steps"

var recipeStepTableColumns = []string{"recipe_uid", "step", "description"}

func NewRecipeStepRow() *RecipeStepRow {
	return &RecipeStepRow{}
}

func (rs *RecipeStepRow) New() *RecipeStepRow {
	return &RecipeStepRow{}
}

func (rs *RecipeStepRow) ValuesToScan() []interface{} {
	return []interface{}{&rs.RecipeUid, &rs.Step, &rs.Description}
}

func (rs *RecipeStepRow) Columns() []string {
	return recipeStepTableColumns
}

func (rs *RecipeStepRow) Table() string {
	return recipeStepTableName
}

func (rs *RecipeStepRow) Scan(row pgx.Row) error {
	return row.Scan(rs.ValuesToScan()...)
}

func (rs *RecipeStepRow) ToEntity() entity.RecipeStep {
	return entity.RecipeStep{
		RecipeUid:   rs.RecipeUid.Bytes,
		Step:        rs.Step,
		Description: rs.Description,
	}
}

func (rs *RecipeStepRow) FromEntity(rse entity.RecipeStep) *RecipeStepRow {
	rs.RecipeUid = pgUidFromUUID(rse.RecipeUid)
	rs.Description = rse.Description
	rs.Step = rse.Step

	return rs
}

func NewRecipeStepRows() *Rows[*RecipeStepRow, entity.RecipeStep] {
	return &Rows[*RecipeStepRow, entity.RecipeStep]{}
}
