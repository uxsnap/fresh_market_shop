package repositoryOrders

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("productsRepository.CreateOrder (uid: %s)", order.Uid)

	p := pgEntity.NewOrderRow().FromEntity(order)

	stmt, args, err := squirrel.
		Insert(p.Table()).
		PlaceholderFormat(squirrel.Dollar).
		Columns("uid", "user_uid", "sum", "status", "created_at", "updated_at").
		Values(p.Uid, p.UserUid, p.Sum, p.Status, p.CreatedAt.Time, p.UpdatedAt.Time).ToSql()

	if err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось создать заказ")
	}

	_, err = r.DB().Exec(ctx, stmt, args...)

	if err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderCreateError, "не удалось создать заказ")
	}

	return nil
}

func (r *OrdersRepository) UpdateOrder(ctx context.Context, order entity.Order) error {
	log.Printf("productsRepository.UpdateOrder (uid: %s)", order.Uid)

	p := pgEntity.NewOrderRow().FromEntity(order)
	if err := r.Update(ctx, p, p.ConditionUidEqual()); err != nil {
		log.Printf("failed to create order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderUpdateError, "не удалось обновить заказ")
	}
	return nil
}

func (r *OrdersRepository) GetOrderByUid(ctx context.Context, uid uuid.UUID) (entity.Order, bool, error) {
	log.Printf("ordersRepository.GetOrderByUid: %s", uid)

	orderRow := pgEntity.NewOrderRow().FromEntity(entity.Order{Uid: uid})
	if err := r.GetOne(ctx, orderRow, orderRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Order{}, false, nil
		}
		log.Printf("failed to get order by uid %s: %v", uid, err)
		return entity.Order{}, false, errors.WithStack(err)
	}
	return orderRow.ToEntity(), true, nil
}

func (r *OrdersRepository) GetOrderWithProducts(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error) {
	log.Printf("ordersRepository.GetOrderWithProducts: %s", userUid)

	orderRow := pgEntity.NewOrderRow()
	orderProductsRow := pgEntity.NewOrderProductsRow()
	orderProductsRows := pgEntity.NewOrderProductsRows()

	sqlSelectPart := r.WithPrefix("o", orderRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as o", orderRow.Table())

	sqlSelectPart = append(
		sqlSelectPart,
		fmt.Sprintf(" (select jsonb_agg(jsonb_build_object('orderUid', op.order_uid, 'productUid', op.product_uid, 'count', op.count)) from %v op where op.order_uid = o.uid)", orderProductsRow.Table()),
	)

	sql := squirrel.
		Select(sqlSelectPart...).From(sqlFromPart).
		PlaceholderFormat(squirrel.Dollar)

	if userUid != uuid.Nil {
		sql = sql.Where(squirrel.Eq{"o.user_uid": userUid})
	}

	if qFilters.Limit != 0 {
		sql = sql.Limit(qFilters.Limit)
	}
	if qFilters.Offset != 0 {
		sql = sql.Offset(qFilters.Offset)
	}

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return nil, err
	}

	fmt.Println(stmt)

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get order with products: %v", err)
		return nil, err
	}

	res := []entity.OrderWithProducts{}
	jsonPhotosBuf := []byte{}

	valsForScan := append(orderRow.ValuesForScan(), &jsonPhotosBuf)

	for rows.Next() {
		if err := rows.Scan(valsForScan...); err != nil {
			return nil, err
		}

		owp := entity.OrderWithProducts{
			Order: orderRow.ToEntity(),
		}

		if err := orderProductsRows.FromJson(jsonPhotosBuf); err != nil {
			owp.Products = nil
		} else {
			owp.Products = orderProductsRows.ToEntity()
		}

		res = append(res, owp)
	}

	return res, nil
}
