package pgEntity

import (
	"encoding/json"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const recipesTableName = "recipes"

type RecipeRow struct {
	Uid         pgtype.UUID
	Name        string
	Description string
	CookingTime int64
	Products    pgtype.JSON
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type recipeProduct struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Measure  string  `json:"measure"`
}

func NewRecipeRow() *RecipeRow {
	return &RecipeRow{}
}

func (rr *RecipeRow) FromEntity(recipe entity.Recipe) (*RecipeRow, error) {
	rr.Uid = pgUidFromUUID(recipe.Uid)
	rr.Name = recipe.Name
	rr.Description = recipe.Description
	rr.CookingTime = recipe.CookingTime

	if len(recipe.Products) == 0 {
		rr.Products = pgtype.JSON{
			Status: pgtype.Null,
		}
	} else {
		productsBts, err := json.Marshal(recipe.Products)
		if err != nil {
			log.Printf("failed to marshal recipe products")
			return nil, err
		}

		rr.Products = pgtype.JSON{
			Bytes:  productsBts,
			Status: pgtype.Present,
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
		Description: rr.Description,
		CookingTime: rr.CookingTime,
		CreatedAt:   rr.CreatedAt.Time,
		UpdatedAt:   rr.UpdatedAt.Time,
	}

	if rr.Products.Status == pgtype.Present {
		var products []recipeProduct
		if err := json.Unmarshal(rr.Products.Bytes, &products); err != nil {
			log.Printf("failed to unmarshal recipe products")
			return entity.Recipe{}, err
		}

		r.Products = make([]entity.RecipeProduct, len(products))
		for i := 0; i < len(products); i++ {
			r.Products[i] = entity.RecipeProduct{
				Name:     products[i].Name,
				Quantity: products[i].Quantity,
				Measure:  products[i].Measure,
			}
		}
	}

	return r, nil
}

var recipesTableColumns = []string{"uid", "name", "description", "cooking_time", "products", "created_at", "updated_at"}

func (rr *RecipeRow) Values() []interface{} {
	return []interface{}{
		rr.Uid, rr.Name, rr.Description, rr.CookingTime, rr.Products, rr.CreatedAt, rr.UpdatedAt,
	}
}

func (rr *RecipeRow) Columns() []string {
	return recipesTableColumns
}

func (rr *RecipeRow) Table() string {
	return recipesTableName
}

func (rr *RecipeRow) Scan(row pgx.Row) error {
	return row.Scan(&rr.Uid, &rr.Name, &rr.Description, &rr.CookingTime, &rr.Products, &rr.CreatedAt, &rr.UpdatedAt)
}

func (rr *RecipeRow) ColumnsForUpdate() []string {
	return []string{"name", "description", "cooking_time", "products", "updated_at"}
}

func (rr *RecipeRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		rr.Name, rr.Description, rr.CookingTime, rr.Products, rr.UpdatedAt,
	}
}

func (rr *RecipeRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": rr.Uid,
	}
}

