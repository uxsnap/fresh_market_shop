package repositoryProducts

import (
	"context"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) UpdateProductPhotos(ctx context.Context, productUid uuid.UUID, imgPaths []string) error {
	log.Printf("productsRepository.UpdateProductPhotos uid: %s, %v", productUid, imgPaths)

	productPhotoRow := pgEntity.NewProductPhotoRow()

	sql := squirrel.Insert(productPhotoRow.Table()).
		Columns(productPhotoRow.Columns()...).PlaceholderFormat(squirrel.Dollar)

	for _, imgPath := range imgPaths {
		sql = sql.Values(uuid.NewV4(), productUid, imgPath)
	}

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to update product photos %s", productUid)
		return errors.WithStack(err)
	}

	_, err = r.DB().Exec(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to update product photos %s", productUid)
		return errors.WithStack(err)
	}

	return nil
}
