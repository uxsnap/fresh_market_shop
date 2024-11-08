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
	rows []*ProductPhotoRow
}

func NewProductPhotoRows() *ProductPhotoRows {
	return &ProductPhotoRows{}
}

func (ppr *ProductPhotoRows) ScanAll(rows pgx.Rows) error {
	ppr.rows = []*ProductPhotoRow{}
	for rows.Next() {
		newRow := &ProductPhotoRow{}

		if err := newRow.Scan(rows); err != nil {
			return err
		}
		ppr.rows = append(ppr.rows, newRow)
	}

	return nil
}

func (ppr *ProductPhotoRows) FromJson(bts []byte) error {
	ppr.rows = nil

	if err := json.Unmarshal(bts, &ppr.rows); err != nil {
		// log.Printf("failed to unmarshal product photos: %v", err)
		return err
	}
	return nil
}

func (ppr *ProductPhotoRows) ToEntity() []entity.ProductPhoto {
	if len(ppr.rows) == 0 {
		return nil
	}

	res := make([]entity.ProductPhoto, len(ppr.rows))
	for i := 0; i < len(ppr.rows); i++ {
		res[i] = ppr.rows[i].ToEntity()
	}
	return res
}
