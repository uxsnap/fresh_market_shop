package pgEntity

import (
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const productPhotosTableName = "product_photos"

type ProductPhotoRow struct {
	Uid        pgtype.UUID `json:"id"`
	ProductUid pgtype.UUID `json:"product_uid"`
	ImgPath    string      `json:"img_path"`
}

func NewProductPhotoRow() *ProductPhotoRow {
	return &ProductPhotoRow{}
}

func (pp *ProductPhotoRow) New() *ProductPhotoRow {
	return &ProductPhotoRow{}
}

func (pp *ProductPhotoRow) FromEntity(productPhoto entity.ProductPhoto) *ProductPhotoRow {
	pp.Uid = pgUidFromUUID(productPhoto.Uid)
	pp.ProductUid = pgUidFromUUID(productPhoto.ProductUid)
	pp.ImgPath = productPhoto.FilePath
	return pp
}

func (pp *ProductPhotoRow) ToEntity() entity.ProductPhoto {
	return entity.ProductPhoto{
		Uid:        pp.Uid.Bytes,
		ProductUid: pp.ProductUid.Bytes,
		FilePath:   pp.ImgPath,
	}
}

var productPhotosTableColumns = []string{"id", "product_uid", "img_path"}

func (pp *ProductPhotoRow) Values() []interface{} {
	return []interface{}{pp.Uid, pp.ProductUid, pp.ImgPath}
}

func (pp *ProductPhotoRow) Columns() []string {
	return productPhotosTableColumns
}

func (pp *ProductPhotoRow) Table() string {
	return productPhotosTableName
}

func (pp *ProductPhotoRow) Scan(row pgx.Row) error {
	return row.Scan(&pp.Uid, &pp.ProductUid, &pp.ImgPath)
}

func (pp *ProductPhotoRow) ColumnsForUpdate() []string {
	return []string{"img_path"}
}

func (pp *ProductPhotoRow) ValuesForUpdate() []interface{} {
	return []interface{}{pp.ImgPath}
}

func (pp *ProductPhotoRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"id": pp.Uid,
	}
}

func (pp *ProductPhotoRow) ConditionProductUidEqual() sq.Eq {
	return sq.Eq{
		"product_uid": pp.ProductUid,
	}
}

type ProductPhotoRows struct {
	*Rows[*ProductPhotoRow, entity.ProductPhoto]
}

func NewProductPhotoRows() *ProductPhotoRows {
	return &ProductPhotoRows{
		&Rows[*ProductPhotoRow, entity.ProductPhoto]{},
	}
}

func (ppr *ProductPhotoRows) FromJson(bts []byte) error {
	ppr.rows = nil

	if err := json.Unmarshal(bts, &ppr.rows); err != nil {
		// log.Printf("failed to unmarshal product photos: %v", err)
		return err
	}
	return nil
}
