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

type RecipeStepRows struct {
	rows []*RecipeStepRow
}

const recipeStepTableName = "recipes_steps"

var recipeStepTableColumns = []string{"recipe_uid", "step", "description"}

func NewRecipeStepRow() *RecipeStepRow {
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

func NewRecipeStepRows() *RecipeStepRows {
	return &RecipeStepRows{}
}

func (rs *RecipeStepRow) Scan(row pgx.Row) error {
	return row.Scan(rs.ValuesToScan()...)
}

func (rs *RecipeStepRow) ToEntity() (entity.RecipeStep, error) {
	r := entity.RecipeStep{
		RecipeUid:   rs.RecipeUid.Bytes,
		Step:        rs.Step,
		Description: rs.Description,
	}

	return r, nil
}

func (rs *RecipeStepRows) ScanAll(rows pgx.Rows) error {
	rs.rows = []*RecipeStepRow{}

	for rows.Next() {
		newRow := &RecipeStepRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		rs.rows = append(rs.rows, newRow)
	}

	return nil
}

func (rr *RecipeStepRows) ToEntity() ([]entity.RecipeStep, error) {
	if len(rr.rows) == 0 {
		return nil, nil
	}

	res := make([]entity.RecipeStep, len(rr.rows))
	for i := 0; i < len(rr.rows); i++ {
		val, err := rr.rows[i].ToEntity()
		if err != nil {
			return nil, err
		}
		res[i] = val
	}
	return res, nil
}
