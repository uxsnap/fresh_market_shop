package pgEntity

import "github.com/jackc/pgtype"

type RecipeStepRow struct {
	RecipeUid   pgtype.UUID
	Step        int64
	Description string
	ImgPath     string
}

const recipeStepTableName = "recipes_steps"

var recipeStepTableColumns = []string{"recipe_uid", "step", "description", "img_path"}

func NewRecipeStepRow() *RecipeStepRow {
	return &RecipeStepRow{}
}

func (rs *RecipeStepRow) ValuesToScan() []interface{} {
	return []interface{}{&rs.RecipeUid, &rs.Step, &rs.Description, &rs.ImgPath}
}

func (rr *RecipeStepRow) Columns() []string {
	return recipeStepTableColumns
}

func (rr *RecipeStepRow) Table() string {
	return recipeStepTableName
}
