package repositoryProducts

import (
	"context"
	"fmt"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *ProductsRepository) DeleteProductPhotos(ctx context.Context, productUid uuid.UUID, photosUids ...uuid.UUID) error {
	log.Printf("productsRepository.DeleteProductPhotos: product uid %s, product photos uids %v", productUid, photosUids)

	productPhotoRow := pgEntity.NewProductPhotoRow().FromEntity(entity.ProductPhoto{ProductUid: productUid})
	stmt := sq.Delete(
		productPhotoRow.Table(),
	).PlaceholderFormat(
		sq.Dollar,
	).Where(
		productPhotoRow.ConditionProductUidEqual(),
	)

	if len(photosUids) != 0 {
		photosUidsArgs := make([]pgtype.UUID, len(photosUids))
		for i := 0; i < len(photosUids); i++ {
			photosUidsArgs[i] = pgtype.UUID{Bytes: photosUids[i], Status: pgtype.Present}
		}

		stmtIn := strings.Builder{}
		stmtIn.WriteString(" photo_uid IN ($1")

		for i := 2; i <= len(photosUidsArgs); i++ {
			stmtIn.WriteString(fmt.Sprintf(",$%d", i))
		}
		stmtIn.WriteString(")")

		stmt = stmt.Where(stmtIn.String())
	}

	sql, args, err := stmt.ToSql()
	if err != nil {
		log.Printf("failed to build query for delete product %s photos: %v", productUid, err)
		return errors.WithStack(err)
	}

	_, err = r.DB().Exec(ctx, sql, args...)
	if err != nil {
		log.Printf("failed to delete product %s photos: %v", productUid, err)
		return errors.WithStack(err)
	}
	return nil
}
