package repositoryOrders

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
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
		log.Printf("failed to update order %s: %v", order.Uid, err)
		return errorWrapper.NewError(errorWrapper.OrderUpdateError, "не удалось обновить заказ")
	}
	return nil
}

func (r *OrdersRepository) GetOrder(ctx context.Context, qFilters entity.QueryFilters) (entity.Order, bool, error) {
	log.Printf("ordersRepository.GetOrderByUid: %s", qFilters.OrderUid)

	if qFilters.OrderUid == uuid.Nil {
		return entity.Order{}, false, errors.New("failed to get order")
	}

	orderRow := pgEntity.NewOrderRow()

	sql := squirrel.Select(orderRow.Columns()...).From(orderRow.Table()).PlaceholderFormat(squirrel.Dollar).Where(
		squirrel.Eq{"uid": qFilters.OrderUid},
	)

	if !uuid.Equal(qFilters.UserUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{"user_uid": qFilters.UserUid})
	}

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to build sql query: %v", err)
		return entity.Order{}, false, errors.WithStack(err)
	}

	err = orderRow.Scan(r.DB().QueryRow(ctx, stmt, args...))
	if err != nil {
		log.Printf("failed to get order: %v", err)
		return entity.Order{}, false, errors.WithStack(err)
	}

	return orderRow.ToEntity(), true, nil
}

func (r *OrdersRepository) GetOrderWithProducts(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error) {
	log.Printf("ordersRepository.GetOrderWithProducts: %s", userUid)

	orderRow := pgEntity.NewOrderRow()
	orderProductsRow := pgEntity.NewOrderProductsRow()
	orderProductsRows := pgEntity.NewOrderProductsRows()

	productPhotoRow := pgEntity.NewProductPhotoRow()

	sqlSelectPart := r.WithPrefix("o", orderRow.Columns())
	sqlFromPart := fmt.Sprintf("%s as o", orderRow.Table())

	sqlSelectPart = append(
		sqlSelectPart,
		fmt.Sprintf(` 
			(select jsonb_agg(jsonb_build_object('orderUid', op.order_uid, 'productUid', op.product_uid, 'count', op.count, 
			'name', (SELECT p.name FROM products p WHERE p.uid = op.product_uid),
			'photos', (select jsonb_agg(jsonb_build_object('id',pp.id,'product_uid',pp.product_uid,'img_path',pp.img_path)) from %v pp where pp.product_uid = op.product_uid)))                  
				from %v op where o.uid = op.order_uid)`,
			productPhotoRow.Table(), orderProductsRow.Table(),
		),
	)

	sql := squirrel.
		Select(sqlSelectPart...).From(sqlFromPart).
		PlaceholderFormat(squirrel.Dollar)

	if userUid != uuid.Nil {
		sql = sql.Where(squirrel.Eq{"o.user_uid": userUid})
	}

	if !uuid.Equal(qFilters.OrderUid, uuid.UUID{}) {
		sql = sql.Where(squirrel.Eq{"o.uid": qFilters.OrderUid})
	}

	if len(qFilters.OrderStatusNotIn) != 0 {
		sql = sql.Where(squirrel.NotEq{"o.status": qFilters.OrderStatusNotIn})
	}

	if qFilters.UpdatedBefore.Unix() >= 0 {
		sql = sql.Where(squirrel.LtOrEq{
			"o.updated_at": pgtype.Timestamp{
				Time:   qFilters.UpdatedBefore,
				Status: pgtype.Present,
			},
		})
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

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get order with products: %v", err)
		return nil, err
	}

	res := []entity.OrderWithProducts{}
	jsonProductsBuf := []byte{}

	valsForScan := append(orderRow.ValuesForScan(), &jsonProductsBuf)

	for rows.Next() {
		if err := rows.Scan(valsForScan...); err != nil {
			return nil, err
		}

		owp := entity.OrderWithProducts{
			Order: orderRow.ToEntity(),
		}

		if err := orderProductsRows.FromJson(jsonProductsBuf); err != nil {
			owp.Products = nil
		} else {
			owp.Products = orderProductsRows.ToEntity()
		}

		res = append(res, owp)
	}

	return res, nil
}

func (r *OrdersRepository) DeleteOrder(ctx context.Context, orderUid uuid.UUID) error {
	log.Printf("ordersRepository.DeleteOrder: %s", orderUid)

	orderRow := pgEntity.NewOrderRow().FromEntity(entity.Order{Uid: orderUid})

	if err := r.Delete(ctx, orderRow, orderRow.ConditionUidEqual()); err != nil {
		log.Printf("ordersRepository.DeleteOrder: failed to delete order %s", orderUid)
		return errors.WithStack(err)
	}
	return nil
}
