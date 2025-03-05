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

func (r *ProductsRepository) GetProductPhotos(ctx context.Context, productUid uuid.UUID, photosUids ...uuid.UUID) ([]entity.ProductPhoto, error) {
	log.Printf("productsRepository.GetProductPhotos: product uid %s, product photos uids %v", productUid, photosUids)

	productPhotoRow := pgEntity.NewProductPhotoRow().FromEntity(entity.ProductPhoto{ProductUid: productUid})
	stmt := sq.Select(
		productPhotoRow.Columns()...,
	).From(
		productPhotoRow.Table(),
	).PlaceholderFormat(sq.Dollar).Where(
		productPhotoRow.ConditionProductUidEqual(),
	)

	if len(photosUids) != 0 {
		photosUidsArgs := make([]pgtype.UUID, len(photosUids))
		for i := 0; i < len(photosUids); i++ {
			photosUidsArgs[i] = pgtype.UUID{Bytes: photosUids[i], Status: pgtype.Present}
		}

		stmtIn := strings.Builder{}
		stmtIn.WriteString(" id IN ($1")

		for i := 2; i <= len(photosUidsArgs); i++ {
			stmtIn.WriteString(fmt.Sprintf(",$%d", i))
		}
		stmtIn.WriteString(")")

		stmt = stmt.Where(stmtIn.String())
	}

	sql, args, err := stmt.ToSql()
	if err != nil {
		log.Printf("failed to build query for get product %s photos: %v", productUid, err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, sql, args...)
	if err != nil {
		log.Printf("failed to get product %s photos: %v", productUid, err)
		return nil, errors.WithStack(err)
	}

	photosRows := pgEntity.NewProductPhotoRows()
	if err := photosRows.ScanAll(rows); err != nil {
		log.Printf("failed to scan product %s photos: %v", productUid, err)
		return nil, errors.WithStack(err)
	}

	return photosRows.ToEntity(), nil
}
